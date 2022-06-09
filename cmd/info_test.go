package cmd

import (
	"fmt"
	"github.com/n1ckl0sk0rtge/cpm/config"
	"github.com/n1ckl0sk0rtge/cpm/cruntime"
	"github.com/n1ckl0sk0rtge/cpm/helper"
	"github.com/stretchr/testify/assert"
	"os"
	"os/exec"
	"testing"
)

func TestInfoCommandNotExists(t *testing.T) {
	config.InitGlobalConfig()
	defer config.RemoveGlobalConfig()

	command := "busybox"
	output := helper.CatchStdOut(t, func() {
		info(nil, []string{command})
	})
	assert.Equal(t, "could not find command busybox\n", output)
}

func TestInfo(t *testing.T) {
	config.InitGlobalConfig()
	defer config.RemoveGlobalConfig()

	confStructure := *config.GetConfigStructure()
	execPath := confStructure[config.ExecPath]

	name := "busybox"
	image := name + "@sha256:205a121ea8a7a142e5f1fdb9ad72c70ffc8e4a56efec5b70b78a93ebfdddae87"
	// download image
	pullImage := exec.Command("sh", "-c", config.Instance.GetString(config.Runtime)+" pull "+image)
	_, err := pullImage.CombinedOutput()

	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// create command
	create(nil, []string{name, "busybox:latest"})
	// remove command add the end
	defer func() {
		err := os.Remove(config.GetConfigProperties().Dir + name)
		err = os.Remove(execPath + name)
		if err != nil {
			fmt.Println(err)
			t.Fail()
		}
	}()

	// get the infos from created command
	command := "busybox"
	output := helper.CatchStdOut(t, func() {
		info(nil, []string{command})
	})

	if err := cruntime.Available(); err != nil {
		assert.Equal(t, "could not inspect image, check if image is available, exit status 1\n", output)
	} else {
		assert.Equal(t, "busybox\nimage:\t\tdocker.io/library/busybox:latest\ndigest:\t\tdocker.io/library/busybox@sha256:205a121ea8a7a142e5f1fdb9ad72c70ffc8e4a56efec5b70b78a93ebfdddae87\nsize:\t\t1464006 byte\nOS/Arch:\tlinux/amd64\n", output)
	}

}
