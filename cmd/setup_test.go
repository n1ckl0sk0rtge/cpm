package cmd

import (
	"github.com/n1ckl0sk0rtge/cpm/config"
	"github.com/n1ckl0sk0rtge/cpm/helper"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("init test configuration")
	config.InitTestConfig()

	code := m.Run()

	log.Println("remove test config")
	helper.RemoveTestConfig()

	os.Exit(code)
}
