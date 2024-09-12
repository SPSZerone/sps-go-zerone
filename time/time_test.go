package time

import (
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	now := time.Now()
	formats := []string{
		FormatYYYYMMDDHHMMSSNsZSNumHHMMSS, FormatYYYYMMDDHHMMSSNsZZNumHHMMSS, FormatYYYYMMDDHHMMSSNsZName,
		FormatYYYYMMDDHHMMSSNsZNameZSNumHHMMSS, FormatYYYYMMDDHHMMSSNsZNameZZNumHHMMSS,
	}
	for _, format := range formats {
		t.Logf("TestFormat format: %+v | time: %+v", format, now.Format(format))
	}
}
