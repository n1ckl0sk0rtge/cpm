package config

func GetConfigStructure() *map[string]string {
	config := make(map[string]string)
	config["runtime"] = "docker"
	config["path"] = "/usr/local/bin/"
	config["container"] = "{}"
	return &config
}
