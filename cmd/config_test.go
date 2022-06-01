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
	conf := config.InitTestConfig(test)
	defer config.RemoveTestConfig(test)

	file := config.GetFilePath(conf)
	output := helper.CatchStdOut(t, func() {
		view(file)
	})
	assert.Equal(t, "container: '{}'\nfilepath: /Users/nkoertge/_projects/cpm/tests/\nruntime: podman\n", output)
}

func TestSet(t *testing.T) {
	test := "testSet"
	conf := config.InitTestConfig(test)
	defer config.RemoveTestConfig(test)

	file := config.GetFilePath(conf)

	key := config.Runtime
	value := fmt.Sprintf("testRuntime")

	set(nil, []string{key, value})

	output := helper.CatchStdOut(t, func() {
		view(file)
	})
	assert.Equal(t, "container: '{}'\nfilepath: /Users/nkoertge/_projects/cpm/tests/\nruntime: testRuntime\n", output)
}
