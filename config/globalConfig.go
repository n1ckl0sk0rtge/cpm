package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var Instance *viper.Viper

func InitGlobalConfig() {
	config := GetConfigProperties()
	createConfFolder(config.Dir)
	initGlobalConfig(config, GetConfigStructure())
}

func initGlobalConfig(values *ConfigValues, configStructure *map[string]string) {
	// init config
	Instance = viper.NewWithOptions(viper.KeyDelimiter(KeyDelimiter))
	Instance.SetConfigName(values.Name)
	Instance.SetConfigType(values.Type)
	Instance.AddConfigPath(values.Dir)

	configFile := GetConfigFilePath(values)
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

func createConfFolder(location string) {
	if _, err := os.Stat(location); os.IsNotExist(err) {
		err := os.Mkdir(location, 0755)
		if err != nil {
			err = fmt.Errorf("could not create application folder '%s', %s", location, err)
			fmt.Println(err)
			return
		}
	}
}

func InitTestGlobalConfig(name string) *ConfigValues {
	config := GetTestConfigProperties(name)
	structure := GetTestConfigStructure("podman", config.Dir)

	createConfFolder(config.Dir)
	initGlobalConfig(config, structure)
	return config
}

func RemoveTestGlobalConfig(name string) {
	values := GetTestConfigProperties(name)
	if err := os.Remove(GetConfigFilePath(values)); err != nil {
		fmt.Println(err)
	}
}
