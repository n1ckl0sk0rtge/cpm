package cmd

import (
	"fmt"
	"github.com/n1ckl0sk0rtge/cpm/config"
	"github.com/n1ckl0sk0rtge/cpm/helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestView(t *testing.T) {
	config.InitGlobalConfig()
	defer config.RemoveGlobalConfig()

	file := config.GetConfigFilePath(config.GetConfigProperties())
	output := helper.CatchStdOut(t, func() {
		view(file)
	})
	assert.Equal(t, "path: /usr/local/bin/\nruntime: docker\nsocket: /var/run/docker.sock\n", output)
}

func TestSet(t *testing.T) {
	config.InitGlobalConfig()
	defer config.RemoveGlobalConfig()

	file := config.GetConfigFilePath(config.GetConfigProperties())

	key := config.Runtime
	value := fmt.Sprintf("testRuntime")

	set(nil, []string{key, value})

	output := helper.CatchStdOut(t, func() {
		view(file)
	})
	assert.Equal(t, "path: /usr/local/bin/\nruntime: testRuntime\nsocket: /var/run/docker.sock\n", output)
}
