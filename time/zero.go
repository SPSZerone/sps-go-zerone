package time

import (
	"time"
)

// NowDayZero zero time of now day
func NowDayZero() time.Time {
	return DayZero(time.Now())
}

// DayZero zero time of day
func DayZero(t time.Time) time.Time {
	return time.Date(
		t.Year(), t.Month(), t.Day(),
		0, 0, 0, 0,
		t.Location(),
	)
}

// NowWeekZero zero time of now week
func NowWeekZero(firstDayOfWeek time.Weekday) time.Time {
	return WeekZero(time.Now(), firstDayOfWeek)
}

// WeekZero zero time of week
func WeekZero(t time.Time, firstDayOfWeek time.Weekday) time.Time {
	offsetDays := int(t.Weekday() - firstDayOfWeek)
	if offsetDays < 0 {
		offsetDays += WeekDays
	}
	firstDayTime := t.Add(-time.Duration(offsetDays) * DurationDay)
	return time.Date(
		firstDayTime.Year(), firstDayTime.Month(), firstDayTime.Day(),
		0, 0, 0, 0,
		firstDayTime.Location(),
	)
}

// NowMonthFirstDayZero zero time of now month first day
func NowMonthFirstDayZero() time.Time {
	return MonthZero(time.Now(), 1)
}

// NowMonthZero zero time of now month
func NowMonthZero(day int) time.Time {
	return MonthZero(time.Now(), day)
}

// MonthFirstDayZero zero time of month first day
func MonthFirstDayZero(t time.Time) time.Time {
	return MonthZero(t, 1)
}

// MonthZero zero time of month
func MonthZero(t time.Time, day int) time.Time {
	return time.Date(
		t.Year(), t.Month(), day,
		0, 0, 0, 0,
		t.Location(),
	)
}

// NowYearFirstMonthZero zero time of now year first month
func NowYearFirstMonthZero() time.Time {
	return YearZero(time.Now(), time.January)
}

// NowYearZero zero time of now year
func NowYearZero(month time.Month) time.Time {
	return YearZero(time.Now(), month)
}

// YearFirstMonthZero zero time of year first month
func YearFirstMonthZero(t time.Time) time.Time {
	return YearZero(t, time.January)
}

// YearZero zero time of year
func YearZero(t time.Time, month time.Month) time.Time {
	return time.Date(
		t.Year(), month, 1,
		0, 0, 0, 0,
		t.Location(),
	)
}
