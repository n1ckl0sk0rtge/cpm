package cmd

import (
	"fmt"
	"github.com/n1ckl0sk0rtge/cpm/config"
	"github.com/spf13/cobra"
	"math/rand"
	"os"
	"os/exec"
	"regexp"
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

	// sanitize input in respect to viper config key delimiter
	if strings.Contains(name, "::") {
		err := fmt.Errorf("name must not contain string '::'")
		fmt.Println(err)
		return
	} else if strings.Contains(image, "::") {
		err := fmt.Errorf("image must not contain string '::'")
		fmt.Println(err)
		return
	}

	version := "latest"
	if strings.Contains(image, ":") {
		parts := strings.Split(image, ":")
		if len(parts) != 2 {
			err := fmt.Errorf("provided image is not valid. Please check format")
			fmt.Println(err)
			return
		}
		image, version = parts[0], parts[1]
	} else if len(entity.Tag) > 0 {
		// if the tag was provided via flag use this instate
		version = entity.Tag
	}

	// check name fulfills regex for container execution
	// otherwise, rename container
	containerName := name
	if matched, _ := regexp.MatchString(`^[a-zA-Z\d][a-zA-Z\d_.-]*$`, name); matched != true {
		// change container name to image name plus random numbers
		// example: golang2357
		random := fmt.Sprintf("%d", rand.Int())
		containerName = image + random
	}

	containerRuntime := config.Instance.Get(config.Runtime)
	if len(entity.Runtime) > 0 {
		containerRuntime = entity.Runtime
	}

	// create executable
	filePath := fmt.Sprintf("%s%s", config.Instance.GetString(config.ExecPath), name)
	executable, err := os.Create(filePath)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer func(executable *os.File) {
		err := executable.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(executable)

	execCommand := fmt.Sprintln(
		containerRuntime, "run", entity.Parameter, "--name", containerName, image+":"+version, entity.Command, "\"$@\"")

	fileContent := fmt.Sprintf("#!/bin/sh\n%s", execCommand)
	_, err = executable.WriteString(fileContent)

	if err != nil {
		fmt.Println(err)
		return
	}

	// make the file executable
	chmodCommand := fmt.Sprintf("chmod +x %s", filePath)
	chmod := exec.Command("sh", "-c", chmodCommand)
	_, err = chmod.Output()

	if err != nil {
		fmt.Println(err)
		return
	}

	// add a link to the command in the configuration
	config.Instance.Set(config.ContainerImage(name), image)
	config.Instance.Set(config.ContainerTag(name), version)
	config.Instance.Set(config.ContainerParameter(name), entity.Parameter)
	config.Instance.Set(config.ContainerCommand(name), entity.Command)
	config.Instance.Set(config.ContainerPath(name), filePath)

	err = config.Instance.WriteConfig()

	if err != nil {
		fmt.Println(err)
		return
	}

}
