package utils

import "time"

var (
	STANDARD_DATE_FORMAT = "2006-01-02"
)

func ParsingTime(timeString string, format string) *time.Time {
	time, err := time.Parse(format, timeString)
	if err != nil {
		return nil
	}

	return &time
}

func FormatTime(time *time.Time, format string) *string {
	if time == nil {
		return nil
	}

	formattedTime := time.Format(format)

	return &formattedTime
}

func DaysBetweenTwoDates(t1 time.Time, t2 time.Time) float64 {
	return t2.Sub(t1).Hours() / 24
}

func DaysBetweenTwoDatesPlusOne(t1 time.Time, t2 time.Time) float64 {
	return (t2.Sub(t1).Hours() / 24) + 1
}

func GetDurationFromTwoDates(t1 time.Time, t2 time.Time) int {
	days := 0
	for {
		if t1.Equal(t2) {
			return days
		}

		if t1.Weekday() != 6 && t1.Weekday() != 7 {
			days++
		}

		t1.Add(time.Hour * 24)
	}
}
