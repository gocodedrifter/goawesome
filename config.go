package main

import (
	"sync"

	"github.com/jinzhu/configor"
)

// Config : configuration for 2pay biller payment
type Config struct {
	Gsp struct {
		TerminalID       string
		PartnerCentralID string
	}

	Iso struct {
		Server struct {
			Listener struct {
				IP   string
				Port string
			}
			Dial struct {
				IP   string
				Port string
			}
		}
	}
}

var configuration = &Config{}

var once sync.Once

// GetConfig : get configuration
func GetConfig() *Config {
	once.Do(func() {
		configor.Load(&configuration, "config.dev.yml")
	})
	return configuration
}
