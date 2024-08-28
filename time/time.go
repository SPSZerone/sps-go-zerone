package time

import "time"

const FormatYYYYMMDDHHMMSSNsZSNum = "2006-01-02 15:04:05.000000000 -07:00:00"
const FormatYYYYMMDDHHMMSSNsZZNum = "2006-01-02 15:04:05.000000000 Z07:00:00"
const FormatYYYYMMDDHHMMSSNsZName = "2006-01-02 15:04:05.000000000 MST"

const FormatYYYYMMDDHHMMSSNsZNameZSNum = "2006-01-02 15:04:05.000000000 MST -07:00:00"
const FormatYYYYMMDDHHMMSSNsZNameZZNum = "2006-01-02 15:04:05.000000000 MST Z07:00:00"

const FormatDefault = FormatYYYYMMDDHHMMSSNsZNameZSNum

func LocalFormatDefault(t time.Time) string {
	return t.Local().Format(FormatDefault)
}
