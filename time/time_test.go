package time

import (
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	now := time.Now()
	formats := []string{
		FormatYYYYMMDDHHMMSSNNsZSNumHHMMSS, FormatYYYYMMDDHHMMSSNNsZZNumHHMMSS, FormatYYYYMMDDHHMMSSNNsZName,
		FormatYYYYMMDDHHMMSSNNsZNameZSNumHHMMSS, FormatYYYYMMDDHHMMSSNNsZNameZZNumHHMMSS,
	}
	for _, format := range formats {
		t.Logf("TestFormat format: %+v | time: %+v", format, now.Format(format))
	}
}
