package goical

import "time"

// New creates a new Calendar instance with specified timezone
func New(tz *time.Location) *Calendar {
	return &Calendar{
		loc: tz,
	}
}
