package logger

import (
	"github.com/rs/zerolog"
	"os"
)

var Log zerolog.Logger

func init() {
	output := zerolog.ConsoleWriter{Out: os.Stderr}
	Log = zerolog.New(output).With().Timestamp().Logger()
}
