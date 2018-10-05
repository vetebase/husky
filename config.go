package husky

import (
	"log"

	"github.com/joho/godotenv"
)

// Load reads and returns a .env file as a map
func Load() map[string]string {
	config, err := godotenv.Read()
	if err != nil {
		log.Fatal("ERROR: Could not load .env file")
	}
	return config
}
