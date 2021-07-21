package types

type ConfigYaml struct {
	Antonyms [][]string `yaml:"antonyms"`
	Homonym  [][]string `yaml:"homonym"`
	Synonyms [][]string `yaml:"synonyms"`
	Confer   [][]string `yaml:"confer"`
	Four     [][]string `yaml:"four"`
}
