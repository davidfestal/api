package v1alpha1

type Keyed interface {
	Key() string
}

type TopLevelLists map[string][]string

type TopLevelListContainer interface {
	GetToplevelLists() TopLevelLists
}
