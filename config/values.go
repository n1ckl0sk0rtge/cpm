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
