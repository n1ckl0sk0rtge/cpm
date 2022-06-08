package cmd

import (
	"fmt"
	"github.com/n1ckl0sk0rtge/cpm/config"
	"github.com/n1ckl0sk0rtge/cpm/helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestView(t *testing.T) {
	test := "testView"
	conf := config.InitTestGlobalConfig(test)
	defer config.RemoveTestGlobalConfig(test)

	file := config.GetConfigFilePath(conf)
	output := helper.CatchStdOut(t, func() {
		view(file)
	})
	assert.Equal(t, "container: '{}'\npath: /Users/nkoertge/_projects/cpm/tests/\nruntime: podman\n", output)
}

func TestSet(t *testing.T) {
	test := "testSet"
	conf := config.InitTestGlobalConfig(test)
	defer config.RemoveTestGlobalConfig(test)

	file := config.GetConfigFilePath(conf)

	key := config.Runtime
	value := fmt.Sprintf("testRuntime")

	set(nil, []string{key, value})

	output := helper.CatchStdOut(t, func() {
		view(file)
	})
	assert.Equal(t, "container: '{}'\npath: /Users/nkoertge/_projects/cpm/tests/\nruntime: testRuntime\n", output)
}
