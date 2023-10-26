package config

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type APIConfig struct {
	Port int `yaml:"port"`
}

type AppConfig struct {
	Database DatabaseConfig `yaml:"database"`
	API      APIConfig      `yaml:"api"`
}
