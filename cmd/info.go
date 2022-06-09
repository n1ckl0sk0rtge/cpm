package cmd

import (
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/n1ckl0sk0rtge/cpm/command"
	"github.com/n1ckl0sk0rtge/cpm/config"
	"github.com/n1ckl0sk0rtge/cpm/cruntime"
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
		if err := cruntime.Available(); err != nil {
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
	name := args[0]

	crun := config.Instance.GetString(config.Runtime)

	if !command.Exists(name, config.GetConfigProperties()) {
		err := fmt.Errorf("could not find command %s", name)
		fmt.Println(err)
		return
	}

	maybeCommandConfig := command.ReadConfig(name, config.GetConfigProperties())
	if maybeCommandConfig == nil {
		err := fmt.Errorf("could not read/found command command config")
		fmt.Println(err)
		return
	}
	commandConfig := *maybeCommandConfig

	fullImage := commandConfig[command.Image] + ":" + commandConfig[command.Tag]
	helper.Dprintln(fullImage)

	// get more infos about the image
	getImageInfosCommand := fmt.Sprintf("%s image inspect %s", crun, fullImage)
	metaData, err := exec.Command("sh", "-c", getImageInfosCommand).Output()

	if err != nil {
		e := fmt.Errorf("could not inspect image, check if image is available, %s", err)
		fmt.Println(e)
		return
	}

	var imageReference string
	imageReference, err = jsonparser.GetString(metaData, "[0]", "RepoTags", "[0]")
	if err != nil {
		fullImage = commandConfig[command.Image]
	}

	// digest
	var digest string

	if crun == cruntime.Podman {
		digest, err = jsonparser.GetString(metaData, "[0]", "Digest")
	} else {
		digest, err = jsonparser.GetString(metaData, "[0]", "Id")
	}

	if err != nil {
		digest = ""
	}

	// size
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

	fmt.Println(name)
	fmt.Printf("image:\t\t%s\n", imageReference)
	fmt.Printf("digest:\t\t%s\n", digest)
	fmt.Printf("size:\t\t%d byte\n", size)
	fmt.Printf("OS/Arch:\t%s/%s\n", operatingSystem, arch)
}
