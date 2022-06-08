package config

const (
	Commands string = "commands"
	Runtime  string = "cruntime"
	Socket   string = "socket"
	ExecPath string = "path"
)

const KeyDelimiter = "::"

func GetConfigStructure() *map[string]string {
	config := make(map[string]string)
	config[Runtime] = "docker"
	config[ExecPath] = "/usr/local/bin/"
	config[Socket] = "/var/run/docker.sock"
	config[Commands] = "{}"
	return &config
}

func GetTestConfigStructure(runtime string, testExecPath string) *map[string]string {
	config := make(map[string]string)
	config[Runtime] = runtime
	config[ExecPath] = testExecPath
	config[Socket] = "/var/run/docker.sock"
	config[Commands] = "{}"
	return &config
}
