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
	digest := "sha256:205a121ea8a7a142e5f1fdb9ad72c70ffc8e4a56efec5b70b78a93ebfdddae87"

	// remove side effects
	// rm digest image
	_, _ = exec.Command(config.Instance.GetString(config.Runtime), "image", "rm", name+"@"+digest).Output()
	// rm latest image
	_, _ = exec.Command(config.Instance.GetString(config.Runtime), "image", "rm", name+":latest").Output()

	// download image
	pullImage := exec.Command(config.Instance.GetString(config.Runtime), "pull", name+"@"+digest)
	_, err := pullImage.CombinedOutput()

	if err != nil {
		fmt.Println("could not fetch image,", err)
		t.Fail()
	}

	// create command
	create(nil, []string{name, name + "@" + digest})
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
		assert.Equal(t, "busybox\nimage:\t\tbusybox\ndigest:\t\tsha256:3fb5cabb64693474716b56fde1566e26c28b3ad4d0651abb08183ede272e11eb\nsize:\t\t1239756 byte\nOS/Arch:\tlinux/amd64\n", output)
	}

}
