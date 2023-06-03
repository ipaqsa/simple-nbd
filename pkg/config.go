package pkg

type ConfigT struct {
	Port               string     `yaml:"port"`
	ReadOnly           bool       `yaml:"readOnly"`
	MinimumBlockSize   int64      `yaml:"minimumBlockSize"`
	PreferredBlockSize int64      `yaml:"preferredBlockSize"`
	MaximumBlockSize   int64      `yaml:"maximumBlockSize"`
	Dir                string     `yaml:"dir"`
	Instances          []Instance `yaml:"instances"`
}

type Instance struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Path        string `yaml:"path"`
}

var Config = ConfigT{}
