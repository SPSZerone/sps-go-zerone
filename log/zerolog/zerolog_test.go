package zerolog

import "testing"

func TestLogger(t *testing.T) {
	logger := NewLogger()
	logger.Print("Print")
	logger.Debug().
		Msg("DebugTest")
	logger.Info().
		Caller(0).
		Msg("InfoTest")
	logger.Warn().
		Caller(0).
		Str("key1", "value1").
		Str("key2", "value2").
		Msg("WarnTest")
	logger.Error().
		Caller(0).
		Str("key1", "value1").
		Str("key2", "value2").
		Str("error", "error").
		Msg("ErrorTest")
}
