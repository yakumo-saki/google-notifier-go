package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/yakumo-saki/go-envconfig"
)

var LOG_LEVEL string

const CONFIG_FILE = "google-notifier-go.env"

var Config ConfigStruct

// 環境変数からconfigをセット
func Initialize() {
	Config = ConfigStruct{}

	envconfig.AddPath(configOnExecDir())
	envconfig.AddPath(configOnHomeDir())

	envconfig.EnableLogWithDefaultLogger()
	err := envconfig.LoadConfig(&Config)
	if err != nil {
		panic("config load error" + err.Error())
	}
}

func configOnExecDir() string {
	execPath, err := os.Executable()
	if err != nil {
		panic("err" + err.Error())
	}
	dir := filepath.Dir(execPath)
	ret := fmt.Sprintf("%s/%s", dir, CONFIG_FILE)
	return ret
}

func configOnHomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic("err" + err.Error())
	}
	homedir := fmt.Sprintf("%s/%s", home, CONFIG_FILE)
	return homedir
}

func Load() {

}
