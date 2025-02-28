package main

import "time"

// Schedule("7/25/2019 13:45:00")
// => 2019-07-25 13:45:00 +0000 UTC
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return t
}

// HasPassed("July 25, 2019 13:45:00")
// => true
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return time.Now().After(t)
}

// IsAfternoonAppointment("Thursday, July 25, 2019 13:45:00")
// => true
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	hour := t.Hour()
	return 12 <= hour && hour < 18
}

// Description("7/25/2019 13:45:00")
// => "You have an appointment on Thursday, July 25, 2019, at 13:45."
func Description(date string) string {
	time := Schedule(date)
	return time.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate()
// => 2020-09-15 00:00:00 +0000 UTC
func AnniversaryDate() time.Time {
	currentYear := time.Now().Year()
	return time.Date(currentYear, time.September, 15, 0, 0, 0, 0, time.UTC)
}
