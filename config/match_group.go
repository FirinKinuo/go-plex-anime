package config

type MatchGroups map[string]Match

type Match struct {
	LibraryPath string   `yaml:"library-path"`
	Patterns    []string `yaml:"patterns"`
}

func (m Match) HasLibraryPath() bool {
	return m.LibraryPath != ""
}
