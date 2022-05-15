//go:build dev
// +build dev

package main

import (
	"log"
	"os"
)

const DevMode = "CPM_DEV_MODE"

func init() {
	log.Println("dev mode is enabled!")
	os.Setenv(DevMode, "true")
}
