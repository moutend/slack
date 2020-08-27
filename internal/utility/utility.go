package utility

import (
	"fmt"
	"time"
)

// Ago converts time to human readable string.
func Ago(t time.Time) string {
	duration := time.Now().UTC().Sub(t)

	const (
		day  = 24 * time.Hour
		week = 7 * day
	)

	if duration < 2*time.Minute {
		return "a minute ago"
	} else if duration < time.Hour {
		return fmt.Sprintf("%v minutes ago", int(duration/time.Minute))
	} else if duration < 2*time.Hour {
		return "an hour ago"
	} else if duration < day {
		return fmt.Sprintf("%v hours ago", int(duration/time.Hour))
	} else if duration < 2*day {
		return "yesterday"
	} else if duration < week {
		return fmt.Sprintf("%v days ago", int(duration/day))
	} else if duration < 2*week {
		return "a week ago"
	} else if duration < 5*week {
		return fmt.Sprintf("%v weeks ago", int(duration/week))
	} else {
		return fmt.Sprintf("%v %v, %v", t.Day(), t.Month(), t.Year())
	}
}
