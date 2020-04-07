package configs

import (
	"github.com/jinzhu/gorm"
)

var (
	// config private var to hold our config
	config *Config
)

// Config struct for config to our app
type Config struct {
	Sqlite struct {
		Database *gorm.DB
	}
}

// InitConfig function to init our config
func InitConfig() {
	if config == nil {
		config = new(Config)

		// set db
		config.Sqlite.Database = SqliteHandler()
	}
}

// GetConfig function to load our config
func GetConfig() *Config {
	return config
}
