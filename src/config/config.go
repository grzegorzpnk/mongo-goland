// Based on: https://gitlab.com/project-emco/core/emco-base/-/blob/main/src/orchestrator/pkg/infra/config/config.go

package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Configuration loads up all the values that are used to configure
// backend implementations
type Configuration struct {
	MongoEndpoint string `json:"mongo-endpoint"`
	MongoPort     string `json:"mongo-port"`
	ClientPort    string `json:"client-port"`
}

// Config is the structure that stores the configuration
var gConfig *Configuration

// readConfigFile reads the specified smsConfig file to setup some env variables
func readConfigFile(file string) (*Configuration, error) {
	f, err := os.Open(file)
	if err != nil {
		return defaultConfiguration(), err
	}
	defer f.Close()

	// Setup some defaults here
	// If the json file has values in it, the defaults will be overwritten
	conf := defaultConfiguration()

	// Read the configuration from json file
	decoder := json.NewDecoder(f)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(conf)
	if err != nil {
		return conf, err
	}

	return conf, nil
}

func defaultConfiguration() *Configuration {

	return &Configuration{
		MongoEndpoint: "mongodb://127.0.0.1",
		MongoPort:     "27017",
		ClientPort:    ":12345",
	}
}

// GetConfiguration returns the configuration for the app.
// It will try to load it if it is not already loaded.
func GetConfiguration() *Configuration {
	if gConfig == nil {
		conf, err := readConfigFile("config.json")
		if err != nil {
			fmt.Printf("Error loading config file: \n", err)
			fmt.Printf("Using defaults...\n")
		}
		gConfig = conf
	}

	return gConfig
}
