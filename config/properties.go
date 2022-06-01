package config

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
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

	return &configValues{
		Dir:  home + "/.cpm/",
		Name: "cpm-conf",
		Type: "yaml",
	}
}

func GetTestConfigProperties(testName string) *configValues {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	projectDir := filepath.Dir(d)

	return &configValues{
		Dir:  projectDir + "/tests/",
		Name: testName,
		Type: "yaml",
	}
}
