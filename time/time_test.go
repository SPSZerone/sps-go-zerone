package time

import (
	"fmt"
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

func TestZeroTime(t *testing.T) {
	testDayZeroTime(t)
	testWeekZeroTime(t, time.Sunday)
	testWeekZeroTime(t, time.Monday)
	testWeekZeroTime(t, time.Tuesday)
	testWeekZeroTime(t, time.Wednesday)
	testWeekZeroTime(t, time.Thursday)
	testWeekZeroTime(t, time.Friday)
	testWeekZeroTime(t, time.Saturday)
	testMonthZeroTime(t)
}

func testDayZeroTime(t *testing.T) {
	timeNow := time.Now()
	t.Logf("testDayZeroTime NowTime %s", timeContent(timeNow))

	dayZeroTime := DayZero(timeNow)
	t.Logf("%s #### %s", timeContent(timeNow), timeContent(dayZeroTime))
}

func testWeekZeroTime(t *testing.T, firstDayOfWeek time.Weekday) {
	timeNow := time.Now()
	t.Logf("testWeekZeroTime firstDayOfWeek:%v NowTime:%s", firstDayOfWeek, timeContent(timeNow))

	timeTest := NowYearFirstMonthZero()
	timeTest = time.Date(timeTest.Year(), timeTest.Month(), timeTest.Day(), timeNow.Hour(), timeNow.Minute(), timeNow.Second(), timeNow.Nanosecond(), timeTest.Location())
	const weekCount = 5
	const totalDay = weekCount * 7

	timeTest = timeTest.Add(-DurationDay * (totalDay >> 1))
	for i := 0; i < totalDay; i++ {
		weekTime := WeekZero(timeTest, firstDayOfWeek)
		t.Logf("%s ---- %s", timeContent(timeTest), timeContent(weekTime))
		timeTest = timeTest.Add(DurationDay)
	}
}

func testMonthZeroTime(t *testing.T) {
	timeNow := time.Now()
	t.Logf("testMonthZeroTime NowTime %s", timeContent(timeNow))

	timeTest := NowYearFirstMonthZero()
	timeTest = time.Date(timeTest.Year(), timeTest.Month(), timeTest.Day(), timeNow.Hour(), timeNow.Minute(), timeNow.Second(), timeNow.Nanosecond(), timeTest.Location())
	timeTest = timeTest.AddDate(0, 0, timeNow.Day()-1)
	const monthCount = 8
	timeTest = timeTest.AddDate(0, -(monthCount >> 1), 0)
	for i := 0; i < monthCount; i++ {
		monthTime := MonthZero(timeTest, 1)
		t.Logf("%s #### %s", timeContent(timeTest), timeContent(monthTime))
		timeTest = timeTest.AddDate(0, 1, 0)
	}
}

func timeContent(time time.Time) string {
	year, week := time.ISOWeek()
	return fmt.Sprintf("time:%v | year:%v month:%v day:%v week:%v weekday:%v", time, year, time.Month(), time.Day(), week, time.Weekday())
}
