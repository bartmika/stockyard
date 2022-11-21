package time

import (
	"time"
)

// Provider provides an interface for abstracting time.
type Provider interface {
	Now() time.Time
	Date(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) time.Time
}

type timeProvider struct{}

// NewTimeProvider Provider contructor that returns the default time provider.
func NewTimeProvider() Provider {
	return timeProvider{}
}

// Now returns the current time.
func (t timeProvider) Now() time.Time {
	return time.Now()
}

func (t timeProvider) Date(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) time.Time {
	return time.Date(year, month, day, hour, min, sec, nsec, loc)
}
