package config

// Config ...
type Config struct {
	APIPort int `yaml:"api port"`
	Linx    SQL `yaml:"linx"`
}

// ConfigYaml...
var ConfigYaml struct {
	API struct {
		Host string `yaml:"host,omitempty"`
		Port string `yaml:"port,omitempty"`
	} `yaml:"api,omitempty"`
	SQL struct {
		Host     string `yaml:"host,omitempty"`
		Port     string `yaml:"port,omitempty"`
		User     string `yaml:"user,omitempty"`
		Password string `yaml:"password,omitempty"`
	} `yaml:"sql,omitempty"`
}

// Auth ...
type Auth struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// SQL ...
type SQL struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Auth `yaml:",inline"`
	Db   string `yaml:"db"`
}
