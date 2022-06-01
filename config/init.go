package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var Instance *viper.Viper

func InitConfig() {
	// init directory
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}

	const applicationFolder string = "/.cpm"

	location := home + applicationFolder
	initConfig(location, GetConfigProperties(), GetConfigStructure())
}

func initConfig(location string, values *configValues, configStructure *map[string]string) {
	if _, err := os.Stat(location); os.IsNotExist(err) {
		err := os.Mkdir(location, 0755)
		if err != nil {
			err = fmt.Errorf("could not create application folder '%s', %s", location, err)
			fmt.Println(err)
		}
	}
	// init config
	Instance = viper.NewWithOptions(viper.KeyDelimiter(KeyDelimiter))
	Instance.SetConfigName(values.Name)
	Instance.SetConfigType(values.Type)
	Instance.AddConfigPath(values.Dir)

	configFile := GetFilePath(values)
	if err := Instance.ReadInConfig(); err != nil { // Find and read the config file
		if _, err := os.Create(configFile); err != nil { // perm 0666
			err = fmt.Errorf("could not create config file '%s', %s", configFile, err)
			fmt.Println(err)
		}

		// set default
		for key, value := range *configStructure {
			Instance.SetDefault(key, value)
		}

		err = Instance.WriteConfig()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func InitTestConfig(name string) *configValues {
	config := GetTestConfigProperties(name)
	structure := GetTestConfigStructure("podman", config.Dir)
	initConfig(config.Dir, config, structure)
	return config
}

func RemoveTestConfig(name string) {
	values := GetTestConfigProperties(name)
	if err := os.Remove(GetFilePath(values)); err != nil {
		fmt.Println(err)
	}
}
