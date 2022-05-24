package main

import (
	"github.com/n1ckl0sk0rtge/cpm/config"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("init test configuration")
	config.InitTestConfig()

	code := m.Run()

	log.Println("remove test config")
	RemoveTestConfig()

	os.Exit(code)

}

func RemoveTestConfig() {
	values := config.GetTestConfigProperties("init")

	err := os.Remove(config.GetFilePath(values))

	if err != nil {
		log.Fatal(err)
	}
}
