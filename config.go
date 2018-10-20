package husky

import (
	"log"

	"github.com/joho/godotenv"
)

// Config handles a Husky service's configuration
type Config struct{}

// Load reads and returns a .env file as a map
func (config *Config) Load() map[string]string {
	c, err := godotenv.Read()
	if err != nil {
		log.Fatal("ERROR: Could not load .env file")
	}
	return c
}
