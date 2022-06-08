package cmd

import (
	"fmt"
	"github.com/n1ckl0sk0rtge/cpm/config"
	"github.com/n1ckl0sk0rtge/cpm/helper"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestListEmpty(t *testing.T) {
	test := "testListEmpty"
	_ = config.InitTestGlobalConfig(test)
	defer config.RemoveTestGlobalConfig(test)

	// empty command
	output := helper.CatchStdOut(t, func() {
		list(nil, nil)
	})
	assert.Equal(t, "  NAME  IMAGE  TAG  PARAMETER  COMMAND  PATH  \n", output)
}

func TestList(t *testing.T) {
	test := "testList"
	conf := config.InitTestGlobalConfig(test)
	defer config.RemoveTestGlobalConfig(test)

	name := "busybox"
	// create command
	entity = flags{}
	create(nil, []string{name, "busybox:latest"})
	// remove command add the end
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			fmt.Println(err)
			t.Fail()
		}
	}(conf.Dir + name)

	output := helper.CatchStdOut(t, func() {
		list(nil, nil)
	})
	assert.Equal(t, "  NAME     IMAGE    TAG     PARAMETER  COMMAND  PATH                                         \n  busybox  busybox  latest                      /Users/nkoertge/_projects/cpm/tests/busybox  \n", output)

}
