package plugins

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"

	"gopkg.in/src-d/go-billy.v4"

	"github.com/blang/semver"
	log "github.com/sirupsen/logrus"
	"gopkg.in/src-d/go-billy.v4/osfs"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/storage/memory"

	"github.com/allez-chauffe/marcel/pkg/config"
	"github.com/allez-chauffe/marcel/pkg/db/plugins"
	gitConfig "gopkg.in/src-d/go-git.v4/config"
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

func fetchVersionsFromGit(url string) (Versions, error) {
	remote := git.NewRemote(memory.NewStorage(), &gitConfig.RemoteConfig{
		Name: "origin",
		URLs: []string{url},
	})

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

func fetchManifestFromGit(url string, ref plumbing.ReferenceName, fs billy.Filesystem) (*plugins.Plugin, error) {
	log.Debugf("Cloning %s (%s) into %s ...", url, ref.Short(), fs.Root())

	repo, err := git.Clone(memory.NewStorage(), fs, &git.CloneOptions{
		URL:           url,
		SingleBranch:  true,
		NoCheckout:    true,
		Depth:         1,
		Tags:          git.NoTags,
		ReferenceName: ref,
	})
	if err != nil {
		return nil, fmt.Errorf("Error while cloning %s into %s : %s", ref.Short(), fs.Root(), err)
	}

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

	plugin := &plugins.Plugin{}
	if err := json.NewDecoder(manifest).Decode(plugin); err != nil {
		return nil, fmt.Errorf("Error while reading manifest : %s", err)
	}

	return plugin, nil
}

// fetchFromGit returns the plugin found in the git repo pointed by url
// It also returns the fullpath of the temporary directory where the plugin's repo content is stored
// The caller should take care of the temporary directory removal
func FetchFromGit(url string) (plugin *plugins.Plugin, tempDir string, err error) {

	versions, err := fetchVersionsFromGit(url)
	if err != nil {
		return nil, tempDir, fmt.Errorf("Error while retreiving versions: %s", err)
	}

	latest, err := versions.Last()
	if err != nil {
		latest = master
		log.Warnf("No versions were found on %s. Using default reference (%s)", url, latest.Short())
	}

	tempDir, err = ioutil.TempDir(config.Default().API().PluginsDir(), "new_plugin")
	if err != nil {
		return nil, tempDir, fmt.Errorf("Error while trying to create temporary directory: %s", err)
	}

	log.Debug("Checking out manifest...")

	plugin, err = fetchManifestFromGit(url, latest.ReferenceName, osfs.New(tempDir))
	if err != nil {
		return nil, tempDir, fmt.Errorf("Error while fetching manifest: %s", err)
	}

	plugin.URL = url
	for _, version := range versions {
		plugin.Versions = append(plugin.Versions, version.String())
	}

	return plugin, tempDir, nil
}
