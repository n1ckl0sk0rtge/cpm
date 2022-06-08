package config

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
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

func GetTestConfigProperties(testName string) *Values {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	projectDir := filepath.Dir(d)

	return &Values{
		Dir:  projectDir + "/tests/",
		Name: testName,
		Type: "yaml",
	}
}
