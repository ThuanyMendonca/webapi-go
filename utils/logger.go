package utils

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var Logger = log.Logger.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
