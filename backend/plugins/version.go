package plugins

import (
	"errors"

	"github.com/blang/semver"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

type Versions []Version
type Version struct {
	plumbing.ReferenceName
	semver.Version
}

func (versions Versions) Len() int {
	return len(versions)
}

func (versions Versions) Swap(i, j int) {
	versions[i], versions[j] = versions[j], versions[i]
}

func (versions Versions) Less(i, j int) bool {
	return versions[i].LE(versions[j].Version)
}

// Last returns the last version from a silce
// The slice should be sorted
func (versions Versions) Last() (Version, error) {
	if len(versions) == 0 {
		return Version{}, errors.New("No versions")
	}
	return versions[len(versions)-1], nil
}

func (version Version) String() string {
	return version.Version.String()
}
