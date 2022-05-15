package cmd

import (
	"fmt"
	"github.com/n1ckl0sk0rtge/cpm/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/exec"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create NAME IMAGE",
	Args:  cobra.ExactArgs(2),
	Short: "A brief description of your command",
	Long: `
Create a new command.

 NAME is the name of the new command, which you will use to execute the command. 

 IMAGE is the OCI image o which the command is based on.
`,

	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			err := fmt.Errorf("not enough arguments provided, need 2 got %d", len(args))
			fmt.Println(err)
			os.Exit(1)
		}
	},

	Run: create,
}

type entity struct {
	Parameter string
	Command   string
	Runtime   string
}

var Entity entity

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

	createCmd.Flags().StringVarP(&Entity.Parameter, "parameter", "p", "-i -t --rm",
		"set parameter for running the container. Default is '-t -i --rm'")

	createCmd.Flags().StringVarP(&Entity.Command, "command", "c", "",
		"set the command that should be executed in the container")

	createCmd.Flags().StringVarP(&Entity.Runtime, "runtime", "r", "",
		"provide container runtime, otherwise the value from the config will be used.")
}

func create(_ *cobra.Command, args []string) {

	containerRuntime := viper.Get(config.Runtime)
	if len(Entity.Runtime) > 0 {
		containerRuntime = Entity.Runtime
	}

	name, image := args[0], args[1]

	execCommand := fmt.Sprintln(
		containerRuntime, "run", Entity.Parameter, "--name", name, image, Entity.Command, "\"$@\"")

	// create executable
	filePath := fmt.Sprintf("%s%s", viper.Get(config.Runtime), name)

	executable, err := os.Create(filePath)

	if err != nil {
		fmt.Println(err)
	}

	defer func(executable *os.File) {
		err := executable.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(executable)

	fileContent := fmt.Sprintf("#!/bin/sh\n%s", execCommand)
	_, err = executable.WriteString(fileContent)

	if err != nil {
		fmt.Println(err)
	}

	// make the file executable
	chmodCommand := fmt.Sprintf("chmod +x %s", filePath)
	chmod := exec.Command("sh", "-c", chmodCommand)
	_, err = chmod.Output()

	if err != nil {
		fmt.Println(err)
	}

	// add a link to the command in the configuration
	commandKey := config.Container + "." + name + "."

	viper.Set(commandKey+"image", image)
	viper.Set(commandKey+"parameter", Entity.Parameter)
	viper.Set(commandKey+"command", Entity.Command)

	err = viper.WriteConfig()

	if err != nil {
		fmt.Println(err)
	}

}
