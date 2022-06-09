package config

const (
	Runtime  string = "runtime"
	Socket   string = "socket"
	ExecPath string = "path"
)

const KeyDelimiter = "::"

func GetConfigStructure() *map[string]string {
	config := make(map[string]string)
	config[Runtime] = "docker"
	config[ExecPath] = "/usr/local/bin/"
	config[Socket] = "/var/run/docker.sock"
	return &config
}
