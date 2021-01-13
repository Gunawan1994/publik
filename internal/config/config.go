package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v3"
)

var (
	config *Config
)

// option defines configuration option
type option struct {
	configFile string
}

// Init initializes `config` from the default config file.
// use `WithConfigFile` to specify the location of the config file
func Init(opts ...Option) error {
	opt := &option{
		configFile: "http-config.yml",
	}
	for _, optFunc := range opts {
		optFunc(opt)
	}

	out, err := ioutil.ReadFile(opt.configFile)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(out, &config)
}

// Option define an option for config package
type Option func(*option)

// WithConfigFile set `config` to use the given config file
func WithConfigFile(file string) Option {
	return func(opt *option) {
		opt.configFile = file
	}
}

// Get config
func Get() *Config {
	if config == nil {
		config = &Config{}
	}
	return config
}
