package helper

import (
	"github.com/n1ckl0sk0rtge/cpm/config"
	"log"
	"os"
)

func RemoveTestConfig() {
	values := config.GetTestConfigProperties("init")
	err := os.Remove(config.GetFilePath(values))

	if err != nil {
		log.Fatal(err)
	}
}
