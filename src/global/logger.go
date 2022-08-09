package global

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/yakumo-saki/google-notifier-go/cmd"
	"github.com/yakumo-saki/google-notifier-go/src/config"
	"github.com/yakumo-saki/google-notifier-go/src/mdnsclient"
)

func InitializeLogger() {
	switch config.LOG_LEVEL {
	case "DEBUG":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "INFO":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "WARN":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "ERROR":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	}

	zerolog.TimeFieldFormat = time.RFC3339Nano
	output := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "2006-01-02 15:04:05.000"}

	log.Logger = zerolog.New(output).With().Timestamp().Caller().Logger()

	callInitializeLog()
}

func callInitializeLog() {
	cmd.InitializeLog()
	mdnsclient.InitializeLog()
}
