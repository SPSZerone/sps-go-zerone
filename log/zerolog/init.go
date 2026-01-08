package zerolog

import "github.com/rs/zerolog"

func init() {
	InitDefault()

	SetGlobalLevelForProduction()
}

func InitDefault() {
	zerolog.LevelFieldMarshalFunc = LevelFieldMarshalFunc
	zerolog.TimestampFunc = TimestampFunc
	zerolog.TimeFieldFormat = DefaultTimeFormat
	zerolog.CallerMarshalFunc = CallerMarshalFunc
	//zerolog.ErrorHandler = ErrorHandler
}

func SetGlobalLevelForProduction() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

func SetGlobalLevelForDevelopment() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
}
