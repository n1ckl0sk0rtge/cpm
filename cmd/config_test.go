package cmd

import (
	"fmt"
	"github.com/n1ckl0sk0rtge/cpm/config"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestView(t *testing.T) {
	var conf = config.GetTestConfigProperties("testConfig")
	file := config.GetFilePath(conf)
	output := view(file)
	assert.Equal(t, output, "container:\n  golang@1.17:\n    command: go\n    image: golang\n    parameter: -v \"$PWD\":/usr/src/myapp -w /usr/src/myapp\n    path: /usr/local/bin/golang@1.17\n    tag: 1.17-stretch\n  redis-cli:\n    command: redis-cli\n    image: redis\n    parameter: -i -t --rm\n    path: /usr/local/bin/redis-cli\n    tag: latest\npath: /usr/local/bin/\nruntime: podman\n")
}

func TestSet(t *testing.T) {
	key := config.Runtime
	value := fmt.Sprintf("testRuntime")

	// init config
	var conf = config.GetTestConfigProperties("testConfig")
	testConfig := viper.NewWithOptions(viper.KeyDelimiter(config.KeyDelimiter))
	testConfig.SetConfigName(conf.Name)
	testConfig.SetConfigType(conf.Type)
	testConfig.AddConfigPath(conf.Dir)

	if err := testConfig.ReadInConfig(); err != nil {
		t.Error(err)
	}

	set(key, value, testConfig)

	// set back to default
	defer set(key, "podman", testConfig)

	file := config.GetFilePath(conf)
	output := view(file)
	assert.Equal(t, output, "container:\n  golang@1.17:\n    command: go\n    image: golang\n    parameter: -v \"$PWD\":/usr/src/myapp -w /usr/src/myapp\n    path: /usr/local/bin/golang@1.17\n    tag: 1.17-stretch\n  redis-cli:\n    command: redis-cli\n    image: redis\n    parameter: -i -t --rm\n    path: /usr/local/bin/redis-cli\n    tag: latest\npath: /usr/local/bin/\nruntime: testRuntime\n")
}
