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
	UID         string
	Timestamp   time.Time
	Summary     string
	Description string
	Location    string
	URL         *url.URL
	Organizer   Person
	Start       time.Time
	End         time.Time
}

type tplEvent struct {
	UID         string
	Summary     string
	Description string
	Location    string
	URL         string
	Organizer   string
	Start       string
	Timestamp   string
	End         string
}

func (tpe *tplEvent) GetLocation() string {
	if tpe.Location != "" {
		return "\r\nLOCATION:" + tpe.Location
	}
	return ""
}

func (tpe *tplEvent) GetURL() string {
	if tpe.URL != "" {
		return "\r\nURL:" + tpe.URL
	}
	return ""
}

func (tpe *tplEvent) GetOrganizer() string {
	if tpe.Organizer != "" {
		return "\r\n" + tpe.Organizer
	}
	return ""
}

func (c *Calendar) convert(e *Event) (out tplEvent) {
	out = tplEvent{
		UID:         e.UID,
		Summary:     e.Summary,
		Description: e.Description,
		Start:       e.Start.In(c.loc).Format(TimeFormat),
		Location:    e.Location,
	}
	if !e.Timestamp.IsZero() {
		out.Timestamp = e.Timestamp.In(c.loc).Format(TimeFormat)
	}
	if !e.End.IsZero() {
		out.End = e.Start.In(c.loc).Add(DefaultDuration).Format(TimeFormat)
	}
	if e.URL != nil {
		out.URL = e.URL.String()
	}
	if e.Organizer.String() != "" {
		out.Organizer = e.Organizer.String()
	}
	return out
}

var tplMain = `{{ define "calendar" }}BEGIN:VCALENDAR
VERSION:2.0
CALSCALE:GREGORIAN
PRODID: https://github.ru/vodolaz095/goical
{{ range .Events }}BEGIN:VEVENT
UID:{{.UID}}{{.GetOrganizer}}{{.GetLocation}}{{.GetURL}}
SUMMARY:{{.Summary}}
DESCRIPTION:{{.Description}}
DTSTAMP;TZID={{$.TimeZoneID}}:{{.Timestamp}}
DTSTART;TZID={{$.TimeZoneID}}:{{.Start}}
DTEND;TZID={{$.TimeZoneID}}:{{.End}}  
END:VEVENT
{{ end }}END:VCALENDAR{{ end }}
`

type tplData struct {
	TimeZoneID string
	Events     []tplEvent
}
