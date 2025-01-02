package goical

import "time"

func New(tz *time.Location) *Calendar {
	return &Calendar{
		loc: tz,
	}
}
