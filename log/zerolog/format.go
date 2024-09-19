package zerolog

import (
	"bytes"
	"fmt"
	"time"

	"github.com/rs/zerolog"
)

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
	return sgr.Output(Caller(), true)
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
