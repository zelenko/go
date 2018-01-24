package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

// Configuration Represents database server and credentials
type Configuration struct {
	Server   string
	Database string
}

// Read and parse the configuration file
func (c *Configuration) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
