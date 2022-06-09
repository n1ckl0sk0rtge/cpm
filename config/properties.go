package config

import (
	"fmt"
	"os"
)

type Values struct {
	Dir  string
	Name string
	Type string
}

func GetConfigFilePath(config *Values) string {
	return config.Dir + config.Name + "." + config.Type
}

func GetConfigProperties() *Values {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Could not get home directory %w", err)
		os.Exit(1)
	}

	return &Values{
		Dir:  home + "/.cpm/",
		Name: ".cpm-conf",
		Type: "yaml",
	}
}
