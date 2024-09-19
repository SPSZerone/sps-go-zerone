package zerolog

import (
	"fmt"
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

func InitDefault() {
	zerolog.LevelFieldMarshalFunc = LevelFieldMarshalFunc
	zerolog.TimestampFunc = TimestampFunc
	zerolog.TimeFieldFormat = DefaultTimeFormat
	zerolog.CallerMarshalFunc = CallerMarshalFunc
	//zerolog.ErrorHandler = ErrorHandler
}

func LevelFieldMarshalFunc(l zerolog.Level) string {
	return zerolog.FormattedLevels[l]
}

func TimestampFunc() time.Time {
	return time.Now().UTC()
}

func CallerMarshalFunc(pc uintptr, file string, line int) string {
	return Caller()
}

func Caller() string {
	caller := "UNKNOWN CALLER"

	pc := make([]uintptr, 8)
	bingo := false
	skip := 5

	for {
		n := runtime.Callers(skip, pc)
		frames := runtime.CallersFrames(pc[:n])
		more := n > 0

		var frame runtime.Frame
		for more {
			skip++
			frame, more = frames.Next()
			if strings.HasPrefix(frame.Function, "github.com/rs/zerolog") {
				continue
			}

			frameFile := frame.File
			frameFunc := frame.Function[strings.LastIndex(frame.Function, ".")+1:]

			caller = fmt.Sprintf("%v:%v:%v", frameFile, frame.Line, frameFunc)
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
