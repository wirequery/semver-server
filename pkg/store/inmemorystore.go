package store

type InMemoryStore struct {
	Store
	versionStore     map[ProjectKey]VersionValue
	projectCandidate map[string]string
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		versionStore:     make(map[ProjectKey]VersionValue),
		projectCandidate: make(map[string]string),
	}
}

func (i *InMemoryStore) AddVersion(key ProjectKey, value VersionValue) {
	i.versionStore[key] = value
}

func (i *InMemoryStore) GetLatestVersion(key ProjectKey) *VersionValue {
	if val, ok := i.versionStore[key]; ok {
		return &val
	}
	return nil
}
