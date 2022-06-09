package cmd

import (
	"fmt"
	"github.com/n1ckl0sk0rtge/cpm/config"
	"github.com/n1ckl0sk0rtge/cpm/helper"
	"github.com/stretchr/testify/assert"
	"os"
	"os/exec"
	"testing"
)

func TestUpdateCommandDoesNotExists(t *testing.T) {
	config.InitGlobalConfig()
	defer config.RemoveGlobalConfig()

	output := helper.CatchStdOut(t, func() {
		updateCommand("busybox")
	})
	assert.Equal(t, "command does not exists\n", output)
}

func TestUpdate(t *testing.T) {
	config.InitGlobalConfig()
	defer config.RemoveGlobalConfig()

	confStructure := *config.GetConfigStructure()
	execPath := confStructure[config.ExecPath]

	image := "docker.io/library/busybox"
	digest := "sha256:5f0395a8920379b7a83cebdc98341f717699ce6b2ab8139fb677c0af4d9a92cb"

	// remove side effects
	// rm digest image
	command := fmt.Sprintf("%s image rm %s",
		config.Instance.GetString(config.Runtime),
		image+"@"+digest,
	)
	_, _ = exec.Command("sh", "-c", command).Output()
	// rm latest image
	command = fmt.Sprintf("%s image rm %s",
		config.Instance.GetString(config.Runtime),
		image+":latest",
	)
	_, _ = exec.Command("sh", "-c", command).Output()

	// download old image
	imageRef := fmt.Sprintf("%s@%s", image, digest)
	command = fmt.Sprintf("%s pull %s",
		config.Instance.GetString(config.Runtime),
		imageRef,
	)
	_, err := exec.Command("sh", "-c", command).Output()
	handleError(t, err)

	// tag it with latest
	command = fmt.Sprintf("%s tag %s %s",
		config.Instance.GetString(config.Runtime),
		imageRef,
		image+":"+"latest",
	)
	_, err = exec.Command("sh", "-c", command).Output()
	handleError(t, err)

	name := "busybox"
	// create command
	entity = flags{}
	create(nil, []string{name, "busybox:latest"})
	// remove command at the end
	defer func() {
		err := os.Remove(config.GetConfigProperties().Dir + name)
		err = os.Remove(execPath + name)
		if err != nil {
			fmt.Println(err)
			t.Fail()
		}
	}()

	output := helper.CatchStdOut(t, func() {
		updateCommand(name)
	})
	assert.NotEqual(t, "command does not exists\n", output)
	assert.NotEqual(t, "busybox is up to date\n", output)

}

func handleError(t *testing.T, err error) {
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}
