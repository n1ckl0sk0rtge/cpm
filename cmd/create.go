package cmd

import (
	"fmt"
	"github.com/n1ckl0sk0rtge/cpm/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/exec"
	"strings"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create NAME IMAGE",
	Args:  cobra.ExactArgs(2),
	Short: "A brief description of your command",
	Long: `
Create a new command.

 NAME is the name of the new command, which you will use to execute the command. 

 IMAGE is the OCI image o which the command is based on. To use the latest version you can 
just select the image name, for example 'redis' or you can explicitly set the latest-tag 'redis:latest'.
In case you want to select a specific version you can either specify the version by selecting providing 
the image tag together with image 'redis:6.2' or you can use the flag '-t 6.2' to set a version.
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

type flags struct {
	Tag       string
	Parameter string
	Command   string
	Runtime   string
}

var entity flags

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	createCmd.Flags().StringVarP(&entity.Tag, "tag", "t", "",
		"set a version (image-tag) for the nwe command")

	createCmd.Flags().StringVarP(&entity.Parameter, "parameter", "p", "-i -t --rm",
		"set parameter for running the container. Default is '-t -i --rm'")

	createCmd.Flags().StringVarP(&entity.Command, "command", "c", "",
		"set the command that should be executed in the container")

	createCmd.Flags().StringVarP(&entity.Runtime, "runtime", "r", "",
		"provide container runtime, otherwise the value from the config will be used")
}

func create(_ *cobra.Command, args []string) {

	name, image := args[0], args[1]

	version := "latest"
	if strings.Contains(image, ":") {
		parts := strings.Split(image, ":")
		if len(parts) != 2 {
			err := fmt.Errorf("provided image is not valid. Please chack format")
			fmt.Println(err)
		}
		image, version = parts[0], parts[1]
	} else if len(entity.Tag) > 0 {
		version = entity.Tag
	}

	containerRuntime := viper.Get(config.Runtime)
	if len(entity.Runtime) > 0 {
		containerRuntime = entity.Runtime
	}

	// create executable
	filePath := fmt.Sprintf("%s%s", viper.Get(config.ExecPath), name)
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

	execCommand := fmt.Sprintln(
		containerRuntime, "run", entity.Parameter, "--name", name, image, entity.Command, "\"$@\"")

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
	viper.Set(config.ContainerImage(name), image)
	viper.Set(config.ContainerTag(name), version)
	viper.Set(config.ContainerParameter(name), entity.Parameter)
	viper.Set(config.ContainerCommand(name), entity.Command)
	viper.Set(config.ContainerPath(name), filePath)

	err = viper.WriteConfig()

	if err != nil {
		fmt.Println(err)
	}

}
