package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	// env var to hold our environment to use in all app
	globalEnv *Environment
)

// Environment struct for environment to our app
type Environment struct {
	Port   string
	Sqlite struct {
		Database struct {
			Name string
		}
	}
}

// InitEnv function to init our environment
func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading file .env")
	}

	if globalEnv == nil {
		globalEnv = new(Environment)

		// set app port from env
		if err := checkOsEnv("PORT"); err != nil {
			log.Fatal(err.Error())
		}
		globalEnv.Port = getOsEnv("PORT")

		// set sqlite db name from env
		if err := checkOsEnv("DB_SQLITE_NAME"); err != nil {
			log.Fatal(err.Error())
		}
		globalEnv.Sqlite.Database.Name = getOsEnv("DB_SQLITE_NAME")
	}
}

// setEnv function to check environment
func checkOsEnv(envName string) error {
	var ok bool

	// check for empty environment
	if _, ok = os.LookupEnv(envName); !ok {
		return fmt.Errorf("Please specify %s in environment", envName)
	}

	return nil
}

// setEnv function to check environment
func getOsEnv(envName string) string {
	return os.Getenv(envName)
}

// GetEnv function to load our environment
func GetEnv() *Environment {
	return globalEnv
}
