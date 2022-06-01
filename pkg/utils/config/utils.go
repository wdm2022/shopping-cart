package config

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type configType = interface{}

func PrintConfigHelp[C configType](config C) {
	text, err := cleanenv.GetDescription(&config, nil)
	if err != nil {
		log.Fatalf("Error when reading configuration schema: %v", err)
	}
	fmt.Println(text)
	os.Exit(2)
}

func ReadConfig[C configType](configFilename string, config *C) {
	_, err := os.Stat(configFilename)

	if errors.Is(err, os.ErrNotExist) {
		// Config file does not exist - load configuration from environment variables
		readConfigFromEnv(config)
	} else if err == nil {
		// Config file exists - load configuration from it and from environment variables
		readConfigFromFile(configFilename, config)
	} else {
		log.Fatalf("Error when reading configuration: %v", err)
	}
}

func readConfigFromFile[C configType](configFilename string, config *C) {
	err := cleanenv.ReadConfig(configFilename, config)
	if err != nil {
		log.Fatalf("Error when reading configuration from file %s: %v", configFilename, err)
	}
}

func readConfigFromEnv[C configType](config *C) {
	err := cleanenv.ReadEnv(config)
	if err != nil {
		log.Fatalf("Error when reading configuration from environment variables: %v", err)
	}
}
