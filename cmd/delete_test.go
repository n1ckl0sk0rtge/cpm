package cmd

import (
	"github.com/n1ckl0sk0rtge/cpm/config"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDelete(t *testing.T) {
	config.InitGlobalConfig()
	defer config.RemoveGlobalConfig()

	confStructure := *config.GetConfigStructure()
	execPath := confStructure[config.ExecPath]

	name := "busybox"
	entity = flags{}
	create(nil, []string{name, "busybox:latest"})

	_, err := os.Stat(config.GetConfigProperties().Dir + name)
	assert.NoError(t, err)
	_, err = os.Stat(execPath + name)
	assert.NoError(t, err)

	deletion(nil, []string{"busybox"})

	_, err = os.Stat(config.GetConfigProperties().Dir + name)
	assert.Error(t, err)
	_, err = os.Stat(execPath + name)
	assert.Error(t, err)
}
