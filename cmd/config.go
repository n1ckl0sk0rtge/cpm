package cmd

import (
	"fmt"
	"github.com/n1ckl0sk0rtge/cpm/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"path/filepath"
)

var conf = config.GetConfigProperties()

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "View and modify the config file",
	Long:  `View and modify the config file`,
	//Run: func(cmd *cobra.Command, args []string) { },
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(configCmd)

	// subcommands
	configCmd.AddCommand(viewCmd)

	configCmd.AddCommand(setCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "Display config settings",
	Long:  `Display config settings`,

	Run: func(cmd *cobra.Command, args []string) {
		configFile := conf.Dir + conf.Name + "." + conf.Type

		filename, _ := filepath.Abs(configFile)
		yamlFile, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(yamlFile))
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

	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			err := fmt.Errorf("not enough arguments provided, need 2 got %d", len(args))
			fmt.Println(err)
			os.Exit(1)
		}
	},

	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		value := args[1]

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
	},
}

func initConfig() {
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

	// init config
	viper.SetConfigName(conf.Name)
	viper.SetConfigType(conf.Type)
	viper.AddConfigPath(conf.Dir)

	configFile := conf.Dir + conf.Name + "." + conf.Type
	if err := viper.ReadInConfig(); err != nil { // Find and read the config file
		if _, err := os.Create(configFile); err != nil { // perm 0666
			err = fmt.Errorf("could not create config file '%s', %s", configFile, err)
			fmt.Println(err)
		}

		// set default
		values := config.GetConfigStructure()
		for key, value := range *values {
			viper.SetDefault(key, value)
		}

		err = viper.WriteConfig()
		if err != nil {
			fmt.Println(err)
		}
	}
}
