package cmd

import (
	"github.com/n1ckl0sk0rtge/cpm/config"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestDelete(t *testing.T) {
	test := "testDelete"
	conf := config.InitTestConfig(test)
	defer config.RemoveTestConfig(test)

	name := "busybox"
	entity = flags{}
	create(nil, []string{name, "busybox:latest"})

	deletion(nil, []string{"busybox"})

	filename, _ := filepath.Abs(config.GetFilePath(conf))
	yamlFile, err := ioutil.ReadFile(filename)
	assert.NoError(t, err)

	assert.Equal(t, "container: {}\npath: /Users/nkoertge/_projects/cpm/tests/\nruntime: podman\n", string(yamlFile))

	_, err = os.Stat(conf.Dir + name)
	assert.Error(t, err)

}
