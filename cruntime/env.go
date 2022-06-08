package cruntime

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/n1ckl0sk0rtge/cpm/config"
	"os"
)

const Runtime = "RUNTIME"

func GetEnvPath(globalConfig *config.ConfigValues) string {
	return globalConfig.Dir + Runtime
}

func InitEnvFile() {
	initEnvFile(config.GetConfigProperties())
}

func initEnvFile(globalConfig *config.ConfigValues) {
	file := GetEnvPath(globalConfig)

	if _, err := os.Stat(file); err != nil {
		if _, err := os.Create(file); err != nil {
			fmt.Println(err)
			return
		}
	}

	WriteEnvFile(globalConfig, config.Instance.GetString(config.Runtime))
}

func WriteEnvFile(globalConfig *config.ConfigValues, runtime string) {
	env, err := godotenv.Unmarshal(fmt.Sprintf("Runtime=%s", runtime))

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
