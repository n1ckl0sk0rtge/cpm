package cmd

import (
	"fmt"
	"github.com/n1ckl0sk0rtge/cpm/config"
	"github.com/n1ckl0sk0rtge/cpm/helper"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestListEmpty(t *testing.T) {
	config.InitGlobalConfig()
	defer config.RemoveGlobalConfig()

	// empty command
	output := helper.CatchStdOut(t, func() {
		list(nil, nil)
	})
	assert.Equal(t, "  NAME  IMAGE  TAG  PARAMETER  COMMAND  \n", output)
}

func TestList(t *testing.T) {
	config.InitGlobalConfig()
	defer config.RemoveGlobalConfig()

	confStructure := *config.GetConfigStructure()
	execPath := confStructure[config.ExecPath]

	name := "busybox"
	// create command
	entity = flags{}
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

	output := helper.CatchStdOut(t, func() {
		list(nil, nil)
	})
	assert.Equal(t, "  NAME     IMAGE    TAG     PARAMETER  COMMAND  \n  busybox  busybox  latest                      \n", output)

}
