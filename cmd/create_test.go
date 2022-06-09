package cmd

import (
	"github.com/n1ckl0sk0rtge/cpm/config"
	"github.com/n1ckl0sk0rtge/cpm/helper"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestCreate(t *testing.T) {
	config.InitGlobalConfig()
	defer config.RemoveGlobalConfig()

	confStructure := *config.GetConfigStructure()
	execPath := confStructure[config.ExecPath]

	// homeDir
	home, err := os.UserHomeDir()

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

	// Test with command flag
	name := "testCommand"
	entity = flags{Command: "sh"}
	create(nil, []string{name, "busybox"})
	envFile, _ := filepath.Abs(config.GetConfigProperties().Dir + name)
	env, err := ioutil.ReadFile(envFile)
	assert.NoError(t, err)
	assert.Equal(t, "COMMAND=\"sh\"\nIMAGE=\"busybox\"\nNAME=\""+name+"\"\nPARAMETER=\"\"\nTAG=\"latest\"\n", string(env))
	_ = os.Remove(config.GetConfigProperties().Dir + name)

	aliasFile, _ := filepath.Abs(execPath + name)
	alias, err := ioutil.ReadFile(aliasFile)
	assert.NoError(t, err)
	assert.Equal(t, "#!/bin/sh\nset -o allexport; source "+home+"/.cpm/.runtime; source "+home+"/.cpm/"+name+"; set +o allexport\n$(echo ${RUNTIME}) run $(echo ${PARAMETER}) --name $(echo ${NAME}) $(echo ${IMAGE}):$(echo ${TAG}) $(echo ${COMMAND}) \"$@\"", string(alias))
	_ = os.Remove(execPath + name)

	// Test with tag flag
	name = "testTag"
	entity = flags{Tag: "1.0.0"}
	create(nil, []string{name, "busybox"})
	envFile, _ = filepath.Abs(config.GetConfigProperties().Dir + name)
	env, err = ioutil.ReadFile(envFile)
	assert.NoError(t, err)
	assert.Equal(t, "COMMAND=\"\"\nIMAGE=\"busybox\"\nNAME=\""+name+"\"\nPARAMETER=\"\"\nTAG=\"1.0.0\"\n", string(env))
	_ = os.Remove(config.GetConfigProperties().Dir + name)
	_ = os.Remove(execPath + name)

	// Test with parameter flag
	name = "testParameter"
	entity = flags{Parameter: "-i"}
	create(nil, []string{name, "busybox"})
	envFile, _ = filepath.Abs(config.GetConfigProperties().Dir + name)
	env, err = ioutil.ReadFile(envFile)
	assert.NoError(t, err)
	assert.Equal(t, "COMMAND=\"\"\nIMAGE=\"busybox\"\nNAME=\""+name+"\"\nPARAMETER=\"-i\"\nTAG=\"latest\"\n", string(env))
	_ = os.Remove(config.GetConfigProperties().Dir + name)
	_ = os.Remove(execPath + name)

	// Test default
	name = "busybox"
	entity = flags{}
	create(nil, []string{name, "busybox:latest"})
	envFile, _ = filepath.Abs(config.GetConfigProperties().Dir + name)
	env, err = ioutil.ReadFile(envFile)
	assert.NoError(t, err)
	assert.Equal(t, "COMMAND=\"\"\nIMAGE=\"busybox\"\nNAME=\""+name+"\"\nPARAMETER=\"\"\nTAG=\"latest\"\n", string(env))
	_ = os.Remove(config.GetConfigProperties().Dir + name)
	_ = os.Remove(execPath + name)
}
