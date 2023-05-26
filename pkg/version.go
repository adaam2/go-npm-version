package npm

import (
	"github.com/aquasecurity/go-version/pkg/semver"
)

// Version represents a semantic version.
type Version = semver.Version

// NewVersion parses a given version and returns an instance of Version
func NewVersion(s string) (Version, error) {
	c, err := newConstraint(s)

	var version = s
	if err == nil {
		version = c.version.String()
	}

	v, err := semver.Parse(version)
	if err != nil {
		return Version{}, err
	}
	return v, nil
}
