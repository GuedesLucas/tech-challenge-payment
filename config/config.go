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

type WebhookConfig struct {
	BaseURL string `mapstructure:"baseUrl"`
	Path    string `mapstructure:"path"`
}

type AppConfig struct {
	Database DatabaseConfig `yaml:"database"`
	Server   APIConfig      `yaml:"server"`
	Webhook  WebhookConfig  `yaml:"webhook"`
}
