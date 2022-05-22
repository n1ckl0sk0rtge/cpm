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
	if _, err := os.Stat(home + applicationFolder); os.IsNotExist(err) {
		err := os.Mkdir(home+applicationFolder, 0755)
		if err != nil {
			err = fmt.Errorf("could not create application folder '%s', %s", applicationFolder, err)
			fmt.Println(err)
		}
	}

	var conf = GetConfigProperties()

	// init config
	Instance = viper.NewWithOptions(viper.KeyDelimiter(KeyDelimiter))
	Instance.SetConfigName(conf.Name)
	Instance.SetConfigType(conf.Type)
	Instance.AddConfigPath(conf.Dir)

	configFile := conf.Dir + conf.Name + "." + conf.Type
	if err := Instance.ReadInConfig(); err != nil { // Find and read the config file
		if _, err := os.Create(configFile); err != nil { // perm 0666
			err = fmt.Errorf("could not create config file '%s', %s", configFile, err)
			fmt.Println(err)
		}

		// set default
		values := GetConfigStructure()
		for key, value := range *values {
			viper.SetDefault(key, value)
		}

		err = Instance.WriteConfig()
		if err != nil {
			fmt.Println(err)
		}
	}
}
