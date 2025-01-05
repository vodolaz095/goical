package goical

import (
	"fmt"
	"net/url"
	"time"
)

// https://icalendar.org/iCalendar-RFC-5545/3-8-4-2-contact.html
// https://icalendar.org/iCalendar-RFC-5545/3-8-4-3-organizer.html
// https://icalendar.org/iCalendar-RFC-5545/3-8-4-6-uniform-resource-locator.html

type Person struct {
	CommonName string
	Email      string
}

func (p *Person) String() string {
	if p.CommonName != "" && p.Email != "" {
		return fmt.Sprintf("ORGANIZER;CN=%q:mailto:%s", p.CommonName, p.Email)
	}
	if p.Email != "" {
		return fmt.Sprintf("ORGANIZER:mailto:%s", p.Email)
	}
	return ""
}

type Event struct {
	// UID - unique identifier of event, mandatory
	UID string
	// Timestamp - when event was created
	Timestamp time.Time
	// Summary - short, human-readable description of event
	Summary string
	// Description - long, human-readable description of event
	Description string
	// Location - human-readable description where event happens, for example `Main Hall`, `Church`, `10th White st` and so on
	Location string
	// URL - http(s) protocol url with some additional info on event
	URL *url.URL
	// Organizer - contact information (name and email) of person who started event
	Organizer Person
	// Start - date and time when event starts - mandatory field
	Start time.Time
	// End  - date and time when event finishes - mandatory field
	End time.Time
}
