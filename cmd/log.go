package cmd

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var l zerolog.Logger

func InitializeLog() {
	l = log.Logger
}
