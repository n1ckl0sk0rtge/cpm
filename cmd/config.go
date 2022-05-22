package cmd

import (
	"fmt"
	"github.com/n1ckl0sk0rtge/cpm/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"path/filepath"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "View and modify the config file",
	Long:  `View and modify the config file`,
}

func init() {
	// init config
	cobra.OnInitialize(config.InitConfig)
	rootCmd.AddCommand(configCmd)
	// sub commands
	configCmd.AddCommand(viewCmd)
	configCmd.AddCommand(setCmd)
}

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "Display config settings",
	Long:  `Display config settings`,
	Run: func(cmd *cobra.Command, args []string) {
		var conf = config.GetConfigProperties()
		configFilePath := config.GetFilePath(conf)
		fmt.Println(view(configFilePath))
	},
}

var setCmd = &cobra.Command{
	Use:   "set PROPERTY_NAME PROPERTY_VALUE",
	Args:  cobra.ExactArgs(2),
	Short: "Set a value in a config file",
	Long: `
Set a value in a config file

 PROPERTY_NAME is a dot delimited name where each token represents either an attribute name or a map key. Map keys may
not contain dots.

 PROPERTY_VALUE is the new value you want to set. Binary fields.

Specifying an attribute name that already exists will replace teh value of existing values.
`,
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		value := args[1]
		set(key, value, config.Instance)
	},
}

func view(configFilePath string) string {
	filename, _ := filepath.Abs(configFilePath)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	return string(yamlFile)
}

func set(key string, value string, viper *viper.Viper) {

	configStructure := *config.GetConfigStructure()

	if _, ok := configStructure[key]; ok {
		viper.Set(key, value)
	} else {
		err := fmt.Errorf("the provided key is not valid, Please ensure to provid an existing key")
		fmt.Println(err)
	}

	err := viper.WriteConfig()
	if err != nil {
		fmt.Println(err)
	}
}
