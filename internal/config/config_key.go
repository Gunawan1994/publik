package config

type (
	// Config is main struct to contain all configuration definition
	Config struct {
		Server Server   `yaml:"server"`
		DB     Database `yaml:"database"`
		Cache  Redis    `yaml:"redis"`
		Application
	}

	// Server is for server related configuration
	Server struct {
		HTTPPort string `yaml:"http_port"`
	}
	// Database holds connection information to DB
	Database struct {
		Driver   string `yaml:"driver"`
		Master   string `yaml:"master"`
		Follower string `yaml:"follower"`
	}
	// Redis struct holds all configuration for interacting with Redis server
	Redis struct {
		MaxIdle   int    `yaml:"maxidle"`
		MaxActive int    `yaml:"maxactive"`
		TimeOut   int    `yaml:"timeout"`
		Wait      bool   `yaml:"wait"`
		Address   string `yaml:"address"`
	}
	// Application is config value for configure application behavior
	Application struct {
		TelewicaraHost string `yaml:"telewicara_host"`
	}
)
