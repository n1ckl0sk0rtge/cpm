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
	name := "busybox"
	create(nil, []string{name, "busybox:latest"})
	filename, _ := filepath.Abs(testDir + name)
	alias, err := ioutil.ReadFile(filename)
	assert.NoError(t, err)
	assert.Equal(t, string(alias), "#!/bin/sh\npodman run -i -t --rm --name "+name+" busybox:latest  \"$@\"\n")
	_ = os.Remove(testDir + name)

	name = "testRuntime"
	entity = flags{Runtime: "testRuntime"}
	create(nil, []string{name, "busybox:latest"})
	filename, _ = filepath.Abs(testDir + name)
	alias, err = ioutil.ReadFile(filename)
	assert.NoError(t, err)
	assert.Equal(t, string(alias), "#!/bin/sh\ntestRuntime run  --name "+name+" busybox:latest  \"$@\"\n")
	_ = os.Remove(testDir + name)

	name = "testTag"
	entity = flags{Tag: "1.0.0"}
	create(nil, []string{name, "busybox"})
	filename, _ = filepath.Abs(testDir + name)
	alias, err = ioutil.ReadFile(filename)
	assert.NoError(t, err)
	assert.Equal(t, string(alias), "#!/bin/sh\npodman run  --name "+name+" busybox:1.0.0  \"$@\"\n")
	_ = os.Remove(testDir + name)

	name = "testCommand"
	entity = flags{Command: "sh"}
	create(nil, []string{name, "busybox"})
	filename, _ = filepath.Abs(testDir + name)
	alias, err = ioutil.ReadFile(filename)
	assert.NoError(t, err)
	assert.Equal(t, string(alias), "#!/bin/sh\npodman run  --name "+name+" busybox:latest sh \"$@\"\n")
	_ = os.Remove(testDir + name)

	name = "testParameter"
	entity = flags{Parameter: "-i"}
	create(nil, []string{name, "busybox"})
	filename, _ = filepath.Abs(testDir + name)
	alias, err = ioutil.ReadFile(filename)
	assert.NoError(t, err)
	assert.Equal(t, string(alias), "#!/bin/sh\npodman run -i --name "+name+" busybox:latest  \"$@\"\n")
	_ = os.Remove(testDir + name)

}
