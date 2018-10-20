package husky

import (
	"log"

	"github.com/joho/godotenv"
)

// Configuration handles a Husky service's configuration
type Configuration struct{}

// Load reads and returns a .env file as a map
func (configuration *Configuration) Load() map[string]string {
	config, err := godotenv.Read()
	if err != nil {
		log.Fatal("ERROR: Could not load .env file")
	}
	return config
}
