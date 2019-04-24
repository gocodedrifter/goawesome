package config

import (
	"sync"

	"github.com/jinzhu/configor"
)

// Config : configuration for 2pay biller payment
type Config struct {
	Gsp struct {
		Terminal string
		Partner  string
		Prepaid  struct {
			Pan string
		}
		Nontaglis struct {
			Pan string
		}
		Postpaid struct {
			Pan string
		}
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
		Messaging struct {
			IP       string
			Port     string
			Handlers string
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

	Db struct {
		URI        string
		Document   string
		Collection string
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
