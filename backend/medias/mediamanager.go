package medias

import (
	"encoding/json"
	"errors"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"log"
	"os"
	"fmt"
	"github.com/Zenika/MARCEL/backend/commons"
	"github.com/Zenika/MARCEL/backend/plugins"
	"strconv"
)

type Manager struct {
	configPath     string
	configFileName string
	configFullpath string
	Config         *Configuration

	pluginManager *plugins.Manager
}

func NewManager(pluginManager *plugins.Manager, configPath, configFilename string) *Manager {
	manager := new(Manager)

	manager.configPath = configPath
	manager.configFileName = configFilename

	manager.pluginManager = pluginManager

	manager.configFullpath = fmt.Sprintf("%s%c%s", configPath, os.PathSeparator, configFilename)
	manager.Config = NewConfiguration()

	return manager
}

// LoadMedias loads medias configuration from DB and stor it in memory
func (m *Manager) Load() {
	log.Printf("Start Loading Medias from DB.")

	m.CreateSaveFileIfNotExist(m.configPath, m.configFileName)

	//Medias configurations are loaded from a JSON file on the FS.
	content, err := ioutil.ReadFile(m.configFullpath)
	commons.Check(err)

	var obj interface{}
	json.Unmarshal([]byte(content), &obj)

	if obj == nil {
		obj = make(map[string]interface{})
	}
	err = mapstructure.Decode(obj.(map[string]interface{}), m.Config)
	if err != nil {
		panic(err)
	}

	log.Print("Medias configurations is loaded...")
}

func (m *Manager) GetConfiguration() (*Configuration) {
	log.Println("Getting global medias config")

	return m.Config
}

func (m *Manager) GetAll() ([]Media) {
	log.Println("Getting all medias")

	return m.Config.Medias
}

// GetMedia Return the media with this id
func (m *Manager) Get(idMedia int) (*Media, error) {

	log.Println("Getting media with id: ", idMedia)
	for _, media := range m.Config.Medias {
		if idMedia == media.ID {
			return &media, nil
		}
	}

	return nil, errors.New("NO_MEDIA_FOUND")
}

// CreateMedia Create a new Media, save it into memory and commit
func (m *Manager) CreateEmpty() (*Media) {

	log.Println("Creating media")

	newMedia := NewMedia()
	newMedia.ID = m.GetNextID()

	//save it into the MediasConfiguration
	m.Save(newMedia)
	m.Commit()

	return newMedia
}

// RemoveMedia Remove media from memory
func (m *Manager) Remove(media *Media) {
	log.Println("Removing media")
	i := m.getPosition(media)

	if i >= 0 {
		m.Config.Medias = append(m.Config.Medias[:i], m.Config.Medias[i+1:]...)
	}
}

// SaveMedia Save media information in memory.
func (m *Manager) Save(media *Media) {
	log.Println("Saving media")
	m.Remove(media)
	m.Config.Medias = append(m.Config.Medias, *media)
}

// Commit Save all medias in DB.
// Here DB is a JSON file
func (m *Manager) Commit() {
	content, _ := json.Marshal(m.Config)

	err := ioutil.WriteFile(m.configFullpath, content, 0644)

	if err != nil {
		log.Println("Cannot save medias configuration:")
		log.Panic(err)
	}
}

// CreateSaveFileIfNotExist check if the save file for medias exists and create it if not.
func (m *Manager) CreateSaveFileIfNotExist(filePath string, fileName string) {

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Println("Data directory did not exist. Create it.")
		os.Mkdir(filePath, 0755)
	}

	var fullPath string = fmt.Sprintf("%s%c%s", filePath, os.PathSeparator, fileName)
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {

		f, err := os.Create(fullPath)
		commons.Check(err)

		log.Println("Medias configuration file created at %v", fullPath)

		f.Close()

		//commit a first time to ensure the configuration has been saved
		m.Commit()
	}
}

func (m *Manager) Activate(media *Media) error {

	//todo : start all backends instances

	sep := string(os.PathSeparator)

	fmt.Printf("Media '%v' has plugin : ", media.Name)
	for _, mp := range media.Plugins {

		fmt.Printf("    * %v (instanceId = %v)\n", mp.EltName, mp.InstanceId)

		// 2.a : duplicate plugin files into "medias/{idMedia}/{plugins_EltName}/{idInstance}"
		err := commons.CopyDir(
			"/home/gwennael.buchet/PROJETS/MARCEL/backend/dist/plugins"+sep+mp.EltName,
			"/home/gwennael.buchet/PROJETS/MARCEL/backend/dist/medias"+sep+strconv.Itoa(media.ID)+sep+mp.EltName+sep+mp.InstanceId)

		if err != nil {
			return err
		}

		// 2.b : if p has backend
		//           => pull backend image or load it (if it has been tar.gzed) into "/medias/{idMedia}/{plugins_EltName}/{idInstance}/back"
		//           => create a dedicated volume
		//           => get a new port to be mapped with the backend, from the pool of ports (getPortNumberForPlugin)
		//           => run docker container
		//           => save the name of the container into a map to be easily stopped later (with Deactivate)

		//create a new instance for the plugin, i.e.:
		//   - duplicate the folder into /medias/{idMedia}/{plugins_EltName}/{idInstance}
		//   - if the plugin has a backend to launch: map the port and run docker container...

	}

	media.IsActive = true

	m.Save(media)

	return nil
}

func (m *Manager) Deactivate(media *Media) error {

	//stop all backends instances and free ports number
	for _, mp := range media.Plugins {
		if mp.BackEnd != nil {

			//todo : call docker manager to stop container for this plugin
			//containerManager.stopImage(mp.BackEnd.ContainerID)

			if mp.BackEnd.Ports != nil {
				ports := mp.BackEnd.Ports
				for _, port := range ports {
					m.FreePortNumberForPlugin(port)
				}
			}
		}
	}

	//remove plugins files
	sep := string(os.PathSeparator)
	os.RemoveAll("medias" + sep + strconv.Itoa(media.ID))

	media.IsActive = false

	m.Save(media)

	return nil
}

// GetMediaPosition Return position of a media in the list
func (m *Manager) getPosition(media *Media) int {
	for p, m := range m.Config.Medias {
		if m.ID == media.ID {
			return p
		}
	}
	return -1
}

func (m *Manager) GetPortNumberForPlugin() int {

	// 1 : try to pop a port number from the pool
	if len(m.Config.PortsPool) > 0 {
		p := m.Config.PortsPool[0]
		//remove the first number
		m.Config.PortsPool = m.Config.PortsPool[1:]
		return p
	}

	// 2 : if pool is empty, just increment the counter
	p := m.Config.NextFreePortNumber
	m.Config.NextFreePortNumber += 1

	return p
}

func (m *Manager) FreePortNumberForPlugin(portNumber int) {
	for i, v := range m.Config.PortsPool {
		if portNumber == v {
			m.Config.PortsPool = append(m.Config.PortsPool[:i], m.Config.PortsPool[i+1:]...)
		}
	}
}

func (m *Manager) GetNextID() int {
	m.Config.LastID = m.Config.LastID + 1
	return m.Config.LastID
}
