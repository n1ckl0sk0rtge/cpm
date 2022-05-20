package helper

import (
	"fmt"
	"github.com/n1ckl0sk0rtge/cpm/config"
	"github.com/spf13/viper"
	"os/exec"
	"strings"
)

func Available() error {
	runtimeString := viper.Get(config.Runtime).(string)

	switch runtimeString {
	case "podman":
		return podmanAvailable()
	default:
		return dockerAvailable()
	}
}

func podmanAvailable() error {
	command := fmt.Sprintf("podman machine list")
	machines := exec.Command("sh", "-c", command)
	machinesOutput, err := machines.Output()

	if err != nil {
		fmt.Println(err)
	}

	output := string(machinesOutput)

	if strings.Contains(output, "Currently running") {
		return nil
	} else {
		return fmt.Errorf("container runtime is not availabe. Please check and provide a valid runtime")
	}
}

func dockerAvailable() error {
	//TODO
	return fmt.Errorf("container runtime is not availabe. Please check and provide a valid runtime")
}
