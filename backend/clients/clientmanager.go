package clients

import (
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/Zenika/MARCEL/backend/commons"
)

//Manager manage clients config.
type Manager struct {
	Config         *Config
	configPath     string
	configFileName string
	configFullPath string
}

type Config struct {
	Clients map[string]*Client `json:"clients"`
}

func newManager(configPath string, configFileName string) *Manager {
	return &Manager{
		Config:         newEmptyConfig(),
		configPath:     configPath,
		configFileName: configFileName,
		configFullPath: filepath.Join(configPath, configFileName),
	}
}

func newEmptyConfig() *Config {
	return &Config{
		Clients: make(map[string]*Client),
	}
}

//LoadFromDB loads medias configuration from DB and stor it in memory
func (m *Manager) LoadFromDB() {
	log.Debugln("Start Loading Clients from DB.")

	commons.LoadFromDB(m)

	log.Debugln("Clients configurations is loaded...")
}

//Commit writes the current configuration to the save file
func (m *Manager) Commit() error {
	return commons.Commit(m)
}

//GetSaveFilePath return the path of the save file for the configuration
//Returns (fullPath, configPath, configFileName)
func (m *Manager) GetSaveFilePath() (string, string, string) {
	return m.configFullPath, m.configPath, m.configFileName
}

//GetConfig return a non-typed pointer to the configuration of the manager
//This is use by commons.LoadFromDB for Decoding configuration from JSON.
func (m *Manager) GetConfig() interface{} {
	return m.Config
}

//Get return the client associated to an ID
func (m *Manager) Get(clientID string) (*Client, bool) {
	client, found := m.Config.Clients[clientID]
	return client, found
}

//GetAll returns the complete map of registered clients (the client ID is the key)
func (m *Manager) GetAll() map[string]*Client {
	return m.Config.Clients
}

func (m *Manager) addNewClient(name string, mediaID int) *Client {
	client := newClient()
	client.Name = name
	client.MediaID = mediaID
	m.Config.Clients[client.ID] = client
	m.Commit()
	return client
}

func (m *Manager) deleteClient(clientID string) {
	delete(m.Config.Clients, clientID)
	m.Commit()
}

func (m *Manager) deleteAllClients() {
	m.Config.Clients = map[string]*Client{}
	m.Commit()
}

func (m *Manager) updateClient(client *Client) *Client {
	m.Config.Clients[client.ID] = client
	m.Commit()
	return client
}
