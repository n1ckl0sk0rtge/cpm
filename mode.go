//go:build dev
// +build dev

package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	log.Println("dev mode is enabled!")
	err := os.Setenv("CPM_DEV_MODE", "true")
	if err != nil {
		fmt.Println(err)
	}
}
