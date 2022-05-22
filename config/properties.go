package config

import (
	"log"
	"os"
)

type configValues struct {
	Dir  string
	Name string
	Type string
}

func GetFilePath(config *configValues) string {
	return config.Dir + config.Name + "." + config.Type
}

func GetConfigProperties() *configValues {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln("Could not get home directory %w", err)
	}

	var dir string = home + "/.cpm/"
	return &configValues{
		Dir:  dir,
		Name: "cpm-conf",
		Type: "yaml",
	}
}

func GetTestConfigProperties(testName string) *configValues {
	return &configValues{
		Dir:  "../tests/config/",
		Name: testName,
		Type: "yaml",
	}
}
