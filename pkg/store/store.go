package store

import "github.com/wirequery/semver-server/pkg/versioning"

type Store interface {
	AddVersion(key ProjectKey, value VersionValue)
	GetLatestVersion(key ProjectKey) *VersionValue
}

type VersionValue struct {
	Version versioning.SemanticVersion
}

type ProjectKey struct {
	Group string
	Name  string
}
