package config

const (
	Container string = "container"
	Runtime   string = "runtime"
	ExecPath  string = "filePath"
)

const KeyDelimiter = "::"

func GetConfigStructure() *map[string]string {
	config := make(map[string]string)
	config[Runtime] = "docker"
	config[ExecPath] = "/usr/local/bin/"
	config[Container] = "{}"
	return &config
}

func GetTestConfigStructure(runtime string, testExecPath string) *map[string]string {
	config := make(map[string]string)
	config[Runtime] = runtime
	config[ExecPath] = testExecPath
	config[Container] = "{}"
	return &config
}

const (
	image     string = "image"
	tag       string = "tag"
	parameter string = "parameter"
	command   string = "command"
	filePath  string = "filePath"
)

func CommandExists(name string) bool {
	return Instance.IsSet(Container + KeyDelimiter + name)
}

func ContainerImage(name string) string {
	return Container + KeyDelimiter + name + KeyDelimiter + image
}

func ContainerTag(name string) string {
	return Container + KeyDelimiter + name + KeyDelimiter + tag
}

func ContainerParameter(name string) string {
	return Container + KeyDelimiter + name + KeyDelimiter + parameter
}

func ContainerCommand(name string) string {
	return Container + KeyDelimiter + name + KeyDelimiter + command
}

func ContainerPath(name string) string {
	return Container + KeyDelimiter + name + KeyDelimiter + filePath
}
