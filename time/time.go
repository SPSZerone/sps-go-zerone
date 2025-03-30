package time

import (
	"time"
)

const (
	DayHours = 24
	WeekDays = 7

	DurationDay  = time.Hour * DayHours
	DurationWeek = DurationDay * WeekDays
)
