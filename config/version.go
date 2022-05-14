package config

import (
	_ "embed"
	"runtime"
)

//go:embed VERSION
var version string

type versionStruct struct {
	Name   string
	Cpm    string
	Golang string
	Os     string
	Arch   string
}

var Version = versionStruct{
	Name:   "cpm",
	Cpm:    version,
	Golang: runtime.Version(),
	Os:     runtime.GOOS,
	Arch:   runtime.GOARCH,
}
