package cmd

import (
	"fmt"
	"github.com/n1ckl0sk0rtge/cpm/config"
	"github.com/n1ckl0sk0rtge/cpm/helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestView(t *testing.T) {
	conf := config.GetTestConfigProperties("init")
	file := config.GetFilePath(conf)

	output := helper.CatchStdOut(t, func() {
		view(file)
	})
	assert.Equal(t, output, "container: '{}'\nfilepath: /Users/nkoertge/_projects/cpm/tests/\nruntime: podman\n")
}

func TestSet(t *testing.T) {
	conf := config.GetTestConfigProperties("init")
	file := config.GetFilePath(conf)

	key := config.Runtime
	value := fmt.Sprintf("testRuntime")

	set(nil, []string{key, value})

	// set back to default
	defer set(nil, []string{key, "podman"})

	output := helper.CatchStdOut(t, func() {
		view(file)
	})
	assert.Equal(t, output, "container: '{}'\nfilepath: /Users/nkoertge/_projects/cpm/tests/\nruntime: testRuntime\n")
}
