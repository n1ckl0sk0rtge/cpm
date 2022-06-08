package runtime

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/n1ckl0sk0rtge/cpm/config"
	"os"
)

func GetEnvPath(globalConfig *config.ConfigValues) string {
	const name = "runtime"
	return globalConfig.Dir + name
}

func initEnvFile(globalConfig *config.ConfigValues) {
	file := GetEnvPath(globalConfig)

	if _, err := os.Stat(file); err != nil {
		if _, err := os.Create(file); err != nil {
			fmt.Println(err)
			return
		}
	}

	env, err := godotenv.Unmarshal(fmt.Sprintf("RUNTIME=%s", config.Instance.GetString(config.Runtime)))

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
