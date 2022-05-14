package config

func GetConfigStructure() *map[string]string {
	config := make(map[string]string)
	config["runtime"] = "docker"
	config["container"] = "{}"
	return &config
}
