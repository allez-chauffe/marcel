package plugins

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"

	"gopkg.in/src-d/go-billy.v4"

	log "github.com/sirupsen/logrus"
	"gopkg.in/src-d/go-billy.v4/osfs"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/storage/memory"

	"github.com/Zenika/MARCEL/backend/commons"
	"github.com/Zenika/MARCEL/backend/config"

	"github.com/blang/semver"
)

const (
	ErrPluginNotFound errPluginNotFound = "NO_PLUGIN_FOUND"
	masterRef                           = "refs/heads/master"
)

var (
	master = Version{ReferenceName: masterRef}
)

type errPluginNotFound string

func (err errPluginNotFound) Error() string {
	return string(err)
}

type Manager struct {
	ConfigPath     string
	ConfigFileName string
	ConfigFullpath string
	Config         *Configuration
}

// NewManager instantiates a new plugin manager
// and initializes its configuration
func NewManager(configPath, configFilename string) *Manager {
	manager := new(Manager)

	manager.ConfigPath = configPath
	manager.ConfigFileName = configFilename

	manager.ConfigFullpath = filepath.Join(configPath, configFilename)
	manager.Config = NewConfiguration()

	return manager
}

// LoadFromDB loads plugins configuration from DB and stores it in memory
func (m *Manager) LoadFromDB() {
	log.Debugln("Start Loading Plugins from DB.")

	commons.LoadFromDB(m)

	log.Debugln("Plugins configurations is loaded...")
}

// GetConfig returns the current configuration
func (m *Manager) GetConfig() interface{} {
	return m.Config
}

// GetAll returns the entire list of registered plugins
func (m *Manager) GetAll() []Plugin {
	log.Debugln("Getting all plugins")

	return m.Config.Plugins
}

// Get Return the plugin
func (m *Manager) Get(eltName string) (*Plugin, error) {

	log.Debugln("Getting plugin with eltName: ", eltName)
	for _, plugin := range m.Config.Plugins {
		if eltName == plugin.EltName {
			return &plugin, nil
		}
	}

	return nil, ErrPluginNotFound
}

// RemoveFromDB Remove plugin
// This is a no-op if the plugin doesn't exists
func (m *Manager) RemoveFromDB(plugin *Plugin) {
	log.Debugln("Removing plugin")
	i := m.GetPosition(plugin)

	if i >= 0 {
		m.Config.Plugins = append(m.Config.Plugins[:i], m.Config.Plugins[i+1:]...)
	} else {
		log.Debugf("Plugin doesn't exsits")
	}

	m.Commit()
}

// Add plugin information.
// Panics if the plugin already exists
func (m *Manager) Add(plugin *Plugin) {
	log.Debugln("Saving plugin")
	if plugin == nil {
		log.Fatal("Can't add nil to plugin list")
	}
	if m.Exists(plugin.EltName) {
		log.Fatalf("Plugin already exists")
	}
	m.Config.Plugins = append(m.Config.Plugins, *plugin)

	m.Commit()
}

// Replace an existing plugin
// This is a no-op when the given plugin is not found
func (m *Manager) Replace(plugin *Plugin) {
	log.Debugf("Replacing plugin")
	i := m.GetPosition(plugin)

	if i >= 0 {
		m.Config.Plugins[i] = *plugin
		m.Commit()
	} else {
		log.Debug("Plugin not found. Replacing ingored")
	}
}

// Commit Save all plugins in DB.
// Here DB is a JSON file
func (m *Manager) Commit() error {
	return commons.Commit(m)
}

// GetPosition returns the position of a plugin in the list
func (m *Manager) GetPosition(plugin *Plugin) int {
	for p, m := range m.Config.Plugins {
		if m.EltName == plugin.EltName {
			return p
		}
	}
	return -1
}

// GetSaveFilePath returns the three components needed to build
// the plugin full directory path
func (m *Manager) GetSaveFilePath() (string, string, string) {
	return m.ConfigFullpath, m.ConfigPath, m.ConfigFileName
}

// Exists checks if a plugin is registered with the given element name
func (m *Manager) Exists(eltName string) bool {
	i := m.GetPosition(&Plugin{EltName: eltName})
	return i != -1
}

// GetDirectory returns the plugin's static files directory path
func (p *Plugin) GetDirectory() string {
	return filepath.Join(config.Config.PluginsPath, p.EltName)
}

// FetchVersionsFromGit returns a sorted list of versions found in the remote tag list
func FetchVersionsFromGit(url string) (Versions, error) {
	repo, err := CloneGitRepository(url, "", nil)
	if err != nil {
		return nil, fmt.Errorf("Error while cloning %s: %s", url, err)
	}

	remote, err := repo.Remote("origin")
	if err != nil {
		return nil, fmt.Errorf("Error retrieving origin remote from %s: %s", url, err)
	}

	log.Debug("Fetching tags...")
	refs, err := remote.List(&git.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("Error fetching tags from %s: %s", url, err)
	}

	var versions Versions
	for _, ref := range refs {
		name := ref.Name()
		if name.IsTag() {
			if version, err := semver.ParseTolerant(name.Short()); err != nil {
				log.Debugf("Ignoring non semver tag: %s", name.Short())
			} else {
				versions = append(versions, Version{name, version})
			}
		}
	}

	sort.Sort(versions)

	return versions, nil
}

// FetchManifestFromGit reads the marcel's manifest file from the given repository
func FetchManifestFromGit(repo *git.Repository, ref plumbing.ReferenceName) (*Plugin, error) {
	wt, err := repo.Worktree()
	if err != nil {
		return nil, fmt.Errorf("Error while getting WorkTree : %s", err)
	}

	if err = wt.Checkout(&git.CheckoutOptions{Branch: ref}); err != nil {
		return nil, fmt.Errorf("Error while checking out manifest: %s", err)
	}

	manifest, err := wt.Filesystem.Open("marcel.json")
	if err != nil {
		return nil, fmt.Errorf("Error while opening manifest : %s", err)
	}
	defer manifest.Close()

	plugin := &Plugin{}
	if err := json.NewDecoder(manifest).Decode(plugin); err != nil {
		return nil, fmt.Errorf("Error while reading manifest : %s", err)
	}

	return plugin, nil
}

// CloneGitRepository returns a repo initialised for url and checked out for ref.
// ref can be omited to fetch default remote's HEAD
// fs can be omitted to avoid checking out repository's content
func CloneGitRepository(url string, ref plumbing.ReferenceName, fs billy.Filesystem) (*git.Repository, error) {
	if fs != nil {
		log.Debugf("Cloning %s (%s) into %s ...", url, ref.Short(), fs.Root())
	} else {
		log.Debugf("Cloning %s (%s)...", url, ref.Short())
	}

	repo, err := git.Clone(memory.NewStorage(), fs, &git.CloneOptions{
		URL:           url,
		SingleBranch:  true,
		NoCheckout:    true,
		Depth:         1,
		Tags:          git.NoTags,
		ReferenceName: ref,
	})
	if err != nil {
		return nil, err
	}

	return repo, nil
}

// FetchFromGit returns the plugin found in the git repo pointed by url
// It also returns the fullpath of the temporary directory where the plugin's repo content is stored
// The caller should take care of the temporary directory removal
func (m *Manager) FetchFromGit(url string) (plugin *Plugin, tempDir string, err error) {

	versions, err := FetchVersionsFromGit(url)
	if err != nil {
		return nil, tempDir, fmt.Errorf("Error while retreiving versions: %s", err)
	}

	latest, err := versions.Last()
	if err != nil {
		latest = master
		log.Warnf("No versions were found on %s. Using default reference (%s)", url, latest.Short())
	}

	tempDir, err = ioutil.TempDir(config.Config.PluginsPath, "new_plugin")
	if err != nil {
		return nil, tempDir, fmt.Errorf("Error while trying to create temporary directory: %s", err)
	}

	repo, err := CloneGitRepository(url, latest.ReferenceName, osfs.New(tempDir))
	if err != nil {
		return nil, tempDir, fmt.Errorf("Error while cloning %s into %s : %s", latest.Short(), tempDir, err)
	}

	log.Debug("Checking out manifest...")

	plugin, err = FetchManifestFromGit(repo, latest.ReferenceName)
	if err != nil {
		return nil, tempDir, fmt.Errorf("Error while fetching manifest: %s", err)
	}

	plugin.URL = url
	for _, version := range versions {
		plugin.Versions = append(plugin.Versions, version.String())
	}

	return plugin, tempDir, nil
}
