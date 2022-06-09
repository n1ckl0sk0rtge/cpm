package cruntime

import (
	"github.com/n1ckl0sk0rtge/cpm/config"
	"github.com/stretchr/testify/assert"
	"os/exec"
	"testing"
)

func TestAvailable(t *testing.T) {
	config.InitGlobalConfig()
	defer config.RemoveGlobalConfig()

	runtimeString := config.Instance.GetString(config.Runtime)

	_, err := exec.Command("sh", "-c", runtimeString+" ps").Output()

	if err != nil {
		assert.Error(t, Available())
	} else {
		assert.NoError(t, Available())
	}

}
