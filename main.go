package main

import (
	"github.com/yakumo-saki/google-notifier-go/cmd"

	"github.com/yakumo-saki/google-notifier-go/src/config"
	"github.com/yakumo-saki/google-notifier-go/src/global"
)

func main() {

	config.Initialize()
	global.InitializeLogger()

	cmd.Execute()
	// os.Exit(ret)
}
