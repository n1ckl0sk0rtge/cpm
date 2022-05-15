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
	tag       string = "tag"
	parameter string = "parameter"
	command   string = "command"
	path      string = "path"
)

func ContainerImage(name string) string {
	return Container + "." + name + "." + image
}

func ContainerTag(name string) string {
	return Container + "." + name + "." + tag
}

func ContainerParameter(name string) string {
	return Container + "." + name + "." + parameter
}

func ContainerCommand(name string) string {
	return Container + "." + name + "." + command
}

func ContainerPath(name string) string {
	return Container + "." + name + "." + path
}
