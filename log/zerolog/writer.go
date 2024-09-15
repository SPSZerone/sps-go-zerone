package zerolog

import (
	"os"

	"github.com/rs/zerolog"
)

func NewConsoleWriter() zerolog.ConsoleWriter {
	writer := zerolog.ConsoleWriter{
		Out: os.Stdout,
		PartsOrder: []string{
			zerolog.LevelFieldName,
			zerolog.TimestampFieldName,

			zerolog.CallerFieldName,
			zerolog.MessageFieldName,
		},
	}

	writer.FormatTimestamp = FormatTimestamp
	writer.FormatLevel = FormatLevel

	writer.FormatCaller = FormatCaller
	writer.FormatMessage = FormatMessage

	writer.FormatFieldName = FormatFieldName
	writer.FormatFieldValue = FormatFieldValue

	writer.FormatErrFieldName = FormatErrFieldName
	writer.FormatErrFieldValue = FormatErrFieldValue

	writer.FormatExtra = FormatExtra
	writer.FormatPrepare = FormatPrepare
	return writer
}
