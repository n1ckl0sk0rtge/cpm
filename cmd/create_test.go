package cmd

import (
	"fmt"
	"github.com/n1ckl0sk0rtge/cpm/config"
	"github.com/n1ckl0sk0rtge/cpm/helper"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestCreate(t *testing.T) {

	fmt.Println(config.Instance.AllKeys())

	// Test name contains unwanted char's
	output := helper.CatchStdOut(t, func() {
		create(nil, []string{"busybox::", "busybox"})
	})
	assert.Equal(t, "name must not contain string '::'\n", output)

	// Test image contains unwanted char's
	output = helper.CatchStdOut(t, func() {
		create(nil, []string{"busybox", "busybox::"})
	})
	assert.Equal(t, "image must not contain string '::'\n", output)

	// Test for undefined image name
	output = helper.CatchStdOut(t, func() {
		create(nil, []string{"busybox", "busybox:latest:3.2"})
	})
	assert.Equal(t, "provided image is not valid. Please check format\n", output)

	testDir := config.GetTestConfigProperties("init").Dir

	// Tests

	create(nil, []string{"busybox", "busybox:latest"})
	filename, _ := filepath.Abs(testDir + "busybox")
	alias, err := ioutil.ReadFile(filename)
	assert.NoError(t, err)
	assert.Equal(t, string(alias), "#!/bin/sh\npodman run -i -t --rm --name busybox busybox  \"$@\"\n")
	_ = os.Remove(testDir + "busybox")

	create(nil, []string{"busybox", "busybox:latest"})
	filename, _ = filepath.Abs(testDir + "busybox")
	alias, err = ioutil.ReadFile(filename)
	assert.NoError(t, err)
	assert.Equal(t, string(alias), "#!/bin/sh\npodman run -i -t --rm --name busybox busybox  \"$@\"\n")
	_ = os.Remove(testDir + "busybox")

}
