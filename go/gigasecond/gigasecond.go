// Package gigasecond calculates the moment someone will have lived one
// gigasecond
package gigasecond

// import path for the time package from the standard library
import (
	"time"
)

// AddGigasecond takes the specified time and returns it after adding a
// gigasecond.
func AddGigasecond(t time.Time) time.Time {
	start := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), int(0), time.UTC)
	afterTenSeconds := start.Add(time.Second * 1000000000)
	return afterTenSeconds
}
