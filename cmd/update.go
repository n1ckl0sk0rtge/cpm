package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/n1ckl0sk0rtge/cpm/config"
	"github.com/n1ckl0sk0rtge/cpm/helper"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"strings"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update [command]",
	Args:  cobra.MaximumNArgs(1),
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple`,

	PreRun: func(cmd *cobra.Command, args []string) {
		if err := helper.Available(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			updateAll()
		} else {
			command := args[0]
			updateCommand(command)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

func updateAll() {
	data := GetCommands()
	for _, value := range data {
		name := value[0]
		updateCommand(name)
	}
}

type imageInspect []struct {
	Id          string   `json:"Id"`
	RepoDigests []string `json:"RepoDigests"`
}

type manifest struct {
	Manifests []struct {
		Digest   string `json:"digest"`
		Platform struct {
			Architecture string `json:"architecture"`
			OS           string `json:"os"`
		} `json:"platform"`
	} `json:"manifests"`
}

func updateCommand(c string) {

	if !config.CommandExists(c) {
		err := fmt.Errorf("command does not exists")
		fmt.Println(err)
		return
	}

	output := fmt.Sprintf("Check for updates for %s...", c)
	fmt.Println(output)

	image := config.Instance.GetString(config.ContainerImage(c))
	tag := config.Instance.GetString(config.ContainerTag(c))
	imageRef := image + ":" + tag

	helper.Dprintln(imageRef)

	getCurrentDigestCommand :=
		fmt.Sprintf("%s inspect %s", config.Instance.GetString(config.Runtime), imageRef)
	currentDigestCommand := exec.Command("sh", "-c", getCurrentDigestCommand)
	currentDigest, err := currentDigestCommand.Output()

	if err != nil {
		fmt.Println(err)
		return
	}

	var imageInspect imageInspect

	err = json.Unmarshal(currentDigest, &imageInspect)

	if err != nil {
		fmt.Println(err)
		return
	}

	digests := imageInspect[0].RepoDigests
	helper.DprintlnSlice(digests)

	// fetch remote digest

	getRemoteDigestsCommand :=
		fmt.Sprintf("%s manifest inspect %s", config.Instance.Get(config.Runtime), imageRef)
	remoteDigestsCommand := exec.Command("sh", "-c", getRemoteDigestsCommand)
	remoteDigests, err := remoteDigestsCommand.Output()

	if err != nil {
		fmt.Println(err)
		return
	}

	var manifest manifest

	err = json.Unmarshal(remoteDigests, &manifest)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Todo check for different os/arch
	remoteDigest := manifest.Manifests[0].Digest
	helper.Dprintln(remoteDigest)

	localDigest := ""
	for _, repoDigest := range imageInspect[0].RepoDigests {
		if strings.Contains(repoDigest, remoteDigest) {
			localDigest = repoDigest
		}
	}
	helper.Dprintln(localDigest)

	if len(localDigest) == 0 {
		output = fmt.Sprintf("update availabe for %s!", c)
		fmt.Println(output)

		output = fmt.Sprintf("download new version %s@%s", image, remoteDigest)
		fmt.Println(output)

		pullNewVersionCommand :=
			fmt.Sprintf("%s pull %s", config.Instance.Get(config.Runtime), imageRef)
		helper.Dprintln(pullNewVersionCommand)
		pullCommand := exec.Command("sh", "-c", pullNewVersionCommand)
		_, err := pullCommand.Output()

		if err != nil {
			fmt.Println(err)
			return
		}

		output = fmt.Sprintf("%s updated successfuly", c)
		fmt.Println(output)

	} else {
		output = fmt.Sprintf("%s is up to date", c)
		fmt.Println(output)
		return
	}

}
