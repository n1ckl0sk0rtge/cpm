package cmd

import (
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/n1ckl0sk0rtge/cpm/config"
	"github.com/n1ckl0sk0rtge/cpm/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/exec"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info COMMAND",
	Args:  cobra.ExactArgs(1),
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple`,

	PreRun: func(cmd *cobra.Command, args []string) {
		if err := helper.Available(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},

	Run: info,
}

func init() {
	rootCmd.AddCommand(infoCmd)
}

func info(_ *cobra.Command, args []string) {
	command := args[0]

	if !getCommandConfig(command) {
		err := fmt.Errorf("could not find command %s", command)
		fmt.Println(err)
		return
	}

	image := viper.Get(config.ContainerImage(command))

	// get more infos about the image
	getImageInfosCommand := fmt.Sprintf("%s image inspect %s", viper.Get(config.Runtime), image)
	imageInfo := exec.Command("sh", "-c", getImageInfosCommand)
	metaData, err := imageInfo.Output()

	if err != nil {
		fmt.Println(err)
	}

	fullImage, err := jsonparser.GetString(metaData, "[0]", "RepoTags", "[0]")
	if err != nil {
		fmt.Println(err)
		fullImage = image.(string)
	}

	var digest string
	digest, err = jsonparser.GetString(metaData, "[0]", "Digest")
	if err != nil {
		fmt.Println(err)
		digest = ""
	}

	var size int64
	size, err = jsonparser.GetInt(metaData, "[0]", "Size")
	if err != nil {
		fmt.Println(err)
		size = 0
	}

	var arch string
	arch, err = jsonparser.GetString(metaData, "[0]", "Architecture")
	if err != nil {
		fmt.Println(err)
		arch = ""
	}

	var operatingSystem string
	operatingSystem, err = jsonparser.GetString(metaData, "[0]", "Os")
	if err != nil {
		fmt.Println(err)
		operatingSystem = ""
	}

	fmt.Println(command)
	fmt.Printf("image:\t\t%s\n", fullImage)
	fmt.Printf("digest:\t\t%s\n", digest)
	fmt.Printf("size:\t\t%d byte\n", size)
	fmt.Printf("OS/Arch:\t%s/%s\n", operatingSystem, arch)
}

func getCommandConfig(command string) bool {

	containers := viper.Get(config.Container).(map[string]interface{})

	for key := range containers {
		if key == command {
			return true
		}
	}

	return false
}
