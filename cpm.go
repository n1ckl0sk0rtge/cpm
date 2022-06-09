package main

import (
	"fmt"
	"github.com/n1ckl0sk0rtge/cpm/cmd"
	"os"
)

func init() {
	// set dev mode
	err := os.Setenv("CPM_DEV_MODE", "false")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	cmd.Execute()
}
