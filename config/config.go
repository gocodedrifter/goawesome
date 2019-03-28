package config

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

	Mti struct {
		Netman struct {
			Request  string
			Response string
		}
		Inquiry struct {
			Request  string
			Response string
		}
		Payment struct {
			Request  string
			Response string
		}
		Reversal struct {
			Request  string
			Response string
			Repeat   struct {
				Request  string
				Response string
			}
		}
	}
}

var configuration = &Config{}

var once sync.Once

// Get : get configuration
func Get() *Config {
	once.Do(func() {
		configor.Load(&configuration, "config/config.dev.yml")
	})
	return configuration
}
