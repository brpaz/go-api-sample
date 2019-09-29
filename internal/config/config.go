package config

import "github.com/kelseyhightower/envconfig"

const envPrefix = "app"

// AppConfig struct that holds application configuration variables
type AppConfig struct {
	Env   string `default:"prod"`
	Port  int    `default:"1323" envconfig:"PORT"`
	Debug bool   `default:"false"`
}

// Config struct with all the configurations
var config AppConfig

// Load Loads the application config
func Load() error {
	return envconfig.Process(envPrefix, &config)
}

// Get Returns the struct contains the application configuration
func Get() AppConfig {
	return config
}
