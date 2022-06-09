package cruntime

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/n1ckl0sk0rtge/cpm/config"
	"os"
)

const Runtime = "RUNTIME"

func GetEnvPath(globalConfig *config.Values) string {
	const fileName = ".runtime"
	return globalConfig.Dir + fileName
}

func InitEnvFile() {
	initEnvFile(config.GetConfigProperties())
}

func initEnvFile(globalConfig *config.Values) {
	file := GetEnvPath(globalConfig)

	if _, err := os.Stat(file); err != nil {
		if _, err := os.Create(file); err != nil {
			fmt.Println(err)
			return
		}
	}

	WriteEnvFile(globalConfig, config.Instance.GetString(config.Runtime))
}

func WriteEnvFile(globalConfig *config.Values, runtime string) {
	env, err := godotenv.Unmarshal(fmt.Sprintf("%s=%s", Runtime, runtime))

	if err != nil {
		fmt.Println(err)
		return
	}

	err = godotenv.Write(env, GetEnvPath(globalConfig))

	if err != nil {
		fmt.Println(err)
		return
	}
}
