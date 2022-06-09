package command

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/n1ckl0sk0rtge/cpm/config"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

const (
	Name      string = "NAME"
	Image     string = "IMAGE"
	Tag       string = "TAG"
	Parameter string = "PARAMETER"
	Commands  string = "COMMAND"
)

func CreateConfig(id string, globalConfig *config.Values) error {
	file := GetConfigPath(id, globalConfig)

	if _, err := os.Stat(file); err == nil {
		// file exists
		err = fmt.Errorf("command already exists")
		return err
	}

	if _, err := os.Create(file); err != nil {
		return err
	}

	return nil
}

func ReadConfig(id string, globalConfig *config.Values) *map[string]string {
	file := GetConfigPath(id, globalConfig)

	err := godotenv.Load(file)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	var commandsEnv map[string]string
	commandsEnv, err = godotenv.Read(file)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &commandsEnv
}

func WriteConfig(commandConfig map[string]string, id string, globalConfig *config.Values) error {
	file := GetConfigPath(id, globalConfig)
	return godotenv.Write(commandConfig, file)
}

func GetConfigPath(id string, config *config.Values) string {
	return config.Dir + id
}

func GenerateRandomCommandIdString() string {
	const length = 16
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ_"

	var seededRand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(b)
}

func Exists(id string, globalConfig *config.Values) bool {
	file := GetConfigPath(id, globalConfig)

	if _, err := os.Stat(file); err == nil {
		// file exists
		return true
	}

	return false
}

func List() ([]string, error) {
	var files []string
	err := filepath.Walk(config.GetConfigProperties().Dir, visit(&files))
	return files, err
}

func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		name := info.Name()
		firstChar := string([]rune(name)[0])
		if firstChar != "." {
			*files = append(*files, name)
		}
		return nil
	}
}
