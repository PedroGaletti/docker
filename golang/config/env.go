package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config : config data struct
type Config struct {
	Port     string
	LogLevel string
}

var (
	// Env export environment variables
	Env = GetConfig()
)

// LoadEnvs : load environment variables
func LoadEnvs() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Runnning the application without a .env file.")
	}
}

// GetConfig : get values of environment variables
func GetConfig() Config {
	var config Config

	LoadEnvs()

	config.Port = os.Getenv("PORT")
	config.LogLevel = os.Getenv("LOG_LEVEL")

	return config
}
