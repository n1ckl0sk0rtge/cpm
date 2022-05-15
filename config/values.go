package config

const (
	Container string = "container"
	Runtime   string = "runtime"
	ExecPath  string = "path"
)

func GetConfigStructure() *map[string]string {
	config := make(map[string]string)
	config[Runtime] = "docker"
	config[ExecPath] = "/usr/local/bin/"
	config[Container] = "{}"
	return &config
}

const (
	image     string = "image"
	parameter string = "parameter"
	command   string = "command"
)

func ContainerImage(name string) string {
	return Container + "." + name + "." + image
}

func ContainerParameter(name string) string {
	return Container + "." + name + "." + parameter
}

func ContainerCommand(name string) string {
	return Container + "." + name + "." + command
}
