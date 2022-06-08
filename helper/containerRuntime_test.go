package helper

import (
	"github.com/n1ckl0sk0rtge/cpm/config"
	"github.com/stretchr/testify/assert"
	"os/exec"
	"testing"
)

func TestAvailable(t *testing.T) {
	test := "testAvailable"
	_ = config.InitTestConfig(test)
	defer config.RemoveTestConfig(test)

	runtimeString := config.Instance.GetString(config.Runtime)

	_, err := exec.Command("sh", "-c", runtimeString+" ps").Output()

	if err != nil {
		assert.Error(t, Available())
	} else {
		assert.NoError(t, Available())
	}

}
