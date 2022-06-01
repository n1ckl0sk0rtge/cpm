package cmd

import (
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/n1ckl0sk0rtge/cpm/config"
	"github.com/n1ckl0sk0rtge/cpm/helper"
	"github.com/spf13/cobra"
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

	image := config.Instance.GetString(config.ContainerImage(command))
	tag := config.Instance.GetString(config.ContainerTag(command))
	fullImage := image + ":" + tag

	// get more infos about the image
	getImageInfosCommand := fmt.Sprintf("%s image inspect %s", config.Instance.Get(config.Runtime), fullImage)
	imageInfo := exec.Command("sh", "-c", getImageInfosCommand)
	metaData, err := imageInfo.Output()

	if err != nil {
		e := fmt.Errorf("could not insepect image, check if image is availabe, %s", err)
		fmt.Println(e)
		return
	}

	var imageReference string
	imageReference, err = jsonparser.GetString(metaData, "[0]", "NamesHistory", "[0]")
	if err != nil {
		fullImage = image.(string)
	}

	var digest string
	digest, err = jsonparser.GetString(metaData, "[0]", "NamesHistory", "[1]")
	if err != nil {
		digest = ""
	}

	var size int64
	size, err = jsonparser.GetInt(metaData, "[0]", "Size")
	if err != nil {
		size = 0
	}

	var arch string
	arch, err = jsonparser.GetString(metaData, "[0]", "Architecture")
	if err != nil {
		arch = ""
	}

	var operatingSystem string
	operatingSystem, err = jsonparser.GetString(metaData, "[0]", "Os")
	if err != nil {
		operatingSystem = ""
	}

	fmt.Println(command)
	fmt.Printf("image:\t\t%s\n", imageReference)
	fmt.Printf("digest:\t\t%s\n", digest)
	fmt.Printf("size:\t\t%d byte\n", size)
	fmt.Printf("OS/Arch:\t%s/%s\n", operatingSystem, arch)
}

func getCommandConfig(command string) bool {

	containers := config.Instance.Get(config.Container)

	if containers == "{}" {
		return false
	}

	for key := range containers.(map[string]interface{}) {
		if key == command {
			return true
		}
	}

	return false
}
