package zerolog

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

const (
	DefaultTimeFormat = "2006-01-02 15:04:05.000000 Z07:00"
)

var (
	FormattedLevels = map[string]zerolog.Level{
		zerolog.FormattedLevels[zerolog.TraceLevel]: zerolog.TraceLevel,
		zerolog.FormattedLevels[zerolog.DebugLevel]: zerolog.DebugLevel,
		zerolog.FormattedLevels[zerolog.InfoLevel]:  zerolog.InfoLevel,
		zerolog.FormattedLevels[zerolog.WarnLevel]:  zerolog.WarnLevel,
		zerolog.FormattedLevels[zerolog.ErrorLevel]: zerolog.ErrorLevel,
		zerolog.FormattedLevels[zerolog.FatalLevel]: zerolog.FatalLevel,
		zerolog.FormattedLevels[zerolog.PanicLevel]: zerolog.PanicLevel,
	}
)

func init() {
	InitDefault()
}

func InitDefault() {
	zerolog.LevelFieldMarshalFunc = LevelFieldMarshalFunc
	zerolog.TimestampFunc = TimestampFunc
	zerolog.TimeFieldFormat = DefaultTimeFormat
	zerolog.CallerMarshalFunc = CallerMarshalFunc
	//zerolog.ErrorHandler = ErrorHandler
}

func NewLogger() zerolog.Logger {
	output := zerolog.ConsoleWriter{
		Out: os.Stdout,
		PartsOrder: []string{
			zerolog.LevelFieldName,
			zerolog.TimestampFieldName,

			zerolog.CallerFieldName,
			zerolog.MessageFieldName,
		},
	}

	output.FormatTimestamp = FormatTimestamp
	output.FormatLevel = FormatLevel

	output.FormatCaller = FormatCaller
	output.FormatMessage = FormatMessage

	output.FormatFieldName = FormatFieldName
	output.FormatFieldValue = FormatFieldValue

	output.FormatErrFieldName = FormatErrFieldName
	output.FormatErrFieldValue = FormatErrFieldValue

	output.FormatExtra = FormatExtra
	output.FormatPrepare = FormatPrepare

	logger := zerolog.New(output).With().Timestamp().Logger()
	return logger
}

func LevelFieldMarshalFunc(l zerolog.Level) string {
	return zerolog.FormattedLevels[l]
}

func TimestampFunc() time.Time {
	return time.Now().UTC()
}

func CallerMarshalFunc(pc uintptr, file string, line int) string {
	return Caller("", "")
}

func Caller(prefix, suffix string) string {
	caller := "UNKNOWN CALLER"

	pc := make([]uintptr, 8)
	bingo := false
	skip := 2

	for {
		n := runtime.Callers(skip, pc)
		frames := runtime.CallersFrames(pc[:n])
		more := n > 0

		var frame runtime.Frame
		for more {
			skip++
			frame, more = frames.Next()
			if strings.HasSuffix(frame.Function, "log/zerolog.Caller") ||
				strings.HasSuffix(frame.Function, "log/zerolog.CallerMarshalFunc") ||
				strings.HasSuffix(frame.Function, "log/zerolog.FormatCaller") ||
				strings.HasPrefix(frame.Function, "github.com/rs/zerolog") {
				continue
			}

			frameFile := frame.File
			frameFunc := frame.Function[strings.LastIndex(frame.Function, ".")+1:]

			caller = fmt.Sprintf("%s%v:%v:%v%s", prefix, frameFile, frame.Line, frameFunc, suffix)
			bingo = true
			break
		}

		if bingo {
			break
		}
	}

	return caller
}

func ErrorHandler(err error) {

}

func FormatTimestamp(time any) string {
	sgr := SGRTimestamp()
	return sgr.Output(time, true)
}

func FormatLevel(level any) string {
	l := FormattedLevels[level.(string)]
	sgr := SGRLevel(l)
	return sgr.Output(level, true)
}

func FormatCaller(caller any) string {
	sgr := SGRCaller()
	if caller != nil {
		return sgr.Output(caller, true)
	}
	return sgr.Output(Caller("", ""), true)
}

func FormatMessage(msg any) string {
	sgr := SGRMessage()
	return sgr.Output(msg, true)
}

func FormatFieldName(name any) string {
	sgr := SGRFieldName()
	sgrSep := SGRFieldSeparator()
	return fmt.Sprintf("%s%s", sgr.Output(name, true), sgrSep.Output("=", true))
}

func FormatFieldValue(value any) string {
	sgr := SGRFieldValue()
	return sgr.Output(value, true)
}

func FormatErrFieldName(name any) string {
	sgr := SGRErrFieldName()
	sgrSep := SGRErrFieldSeparator()
	return fmt.Sprintf("%s%s", sgr.Output(name, true), sgrSep.Output("=", true))
}

func FormatErrFieldValue(value any) string {
	sgr := SGRErrFieldValue()
	return sgr.Output(value, true)
}

func FormatExtra(extra map[string]interface{}, buff *bytes.Buffer) error {
	timeLocal := FormatTimestamp(time.Now().Local().Format(DefaultTimeFormat))
	level := FormatLevel(extra[zerolog.LevelFieldName])

	buff.WriteString(fmt.Sprintf(" %s %s", timeLocal, level))
	return nil
}

func FormatPrepare(prepare map[string]any) error {
	return nil
}
