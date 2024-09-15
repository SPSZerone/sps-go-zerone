package zerolog

import (
	"github.com/rs/zerolog"
)

func init() {
	InitDefault()
}

func NewLogger() zerolog.Logger {
	return NewConsoleLogger()
}

func NewConsoleLogger() zerolog.Logger {
	logger := zerolog.New(NewConsoleWriter()).With().Timestamp().Logger()
	return logger
}
