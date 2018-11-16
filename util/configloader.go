package util

import (
	"encoding/json"
	"log"
	"os"
)

// Environment can be production or development
type Environment string

// use Production to a release environment, Development otherwise
const (
	Production  Environment = "prod"
	Development Environment = "dev"
)

// Conf holds the representation of what is in a configuration file
var Conf Configuration

func init() {
	Conf = loadConfig(Development)
}

// loadConfig loads a configuration file based on the environment type
// returns the parsed configuration file
func loadConfig(env Environment) Configuration {
	fileName := string("config." + env + ".json")
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err = decoder.Decode(&configuration)

	if err != nil {
		log.Fatal(err)
	}

	return configuration
}
