package goical

import (
	"fmt"
	"io"
	"sort"
	"time"
)

const TimeFormat = "20060102T150405"

type Calendar struct {
	loc    *time.Location
	events []Event
}

func (c *Calendar) AddEvent(input Event) *Calendar {
	if input.UID == "" {
		return c
	}
	if input.Start.IsZero() {
		return c
	}
	if input.End.IsZero() {
		return c
	}
	if input.Timestamp.IsZero() {
		input.Timestamp = time.Now()
	}
	if input.Start.After(input.End) {
		return c
	}
	c.events = append(c.events, input)
	return c
}

func (c *Calendar) Render(writer io.Writer) (err error) {
	var org string
	if c.loc == nil {
		c.loc = time.Local
	}
	_, err = fmt.Fprint(writer, "BEGIN:VCALENDAR\r\nVERSION:2.0\r\nCALSCALE:GREGORIAN\r\nPRODID:https://github.ru/vodolaz095/goical\r\n")
	if err != nil {
		return err
	}
	sort.Slice(c.events, func(i, j int) bool {
		return c.events[i].Start.Before(c.events[j].Start)
	})
	for i := range c.events {
		if c.events[i].UID == "" {
			continue
		}
		if c.events[i].Start.IsZero() {
			continue
		}
		if c.events[i].End.IsZero() {
			continue
		}
		if c.events[i].Timestamp.IsZero() {
			continue
		}
		_, err = fmt.Fprint(writer, "BEGIN:VEVENT\r\n")
		if err != nil {
			return err
		}
		_, err = fmt.Fprintf(writer, "UID:%s\r\n", c.events[i].UID)
		if err != nil {
			return err
		}
		if c.events[i].Summary != "" {
			_, err = fmt.Fprintf(writer, "SUMMARY:%s\r\n", c.events[i].Summary)
			if err != nil {
				return err
			}
		}
		if c.events[i].Description != "" {
			_, err = fmt.Fprintf(writer, "DESCRIPTION:%s\r\n", c.events[i].Description)
			if err != nil {
				return err
			}
		}
		org = c.events[i].Organizer.String()
		if org != "" {
			_, err = fmt.Fprintf(writer, "%s\r\n", org)
			if err != nil {
				return err
			}
		}
		if c.events[i].Location != "" {
			_, err = fmt.Fprintf(writer, "LOCATION:%s\r\n", c.events[i].Location)
			if err != nil {
				return err
			}
		}
		if c.events[i].URL != nil {
			_, err = fmt.Fprintf(writer, "URL:%s\r\n", c.events[i].URL.String())
			if err != nil {
				return err
			}
		}
		_, err = fmt.Fprintf(writer, "DTSTAMP;TZID=%s:%s\r\n",
			c.loc.String(), c.events[i].Timestamp.Format(TimeFormat),
		)
		if err != nil {
			return err
		}
		_, err = fmt.Fprintf(writer, "DTSTART;TZID=%s:%s\r\n",
			c.loc.String(), c.events[i].Start.Format(TimeFormat),
		)
		if err != nil {
			return err
		}
		_, err = fmt.Fprintf(writer, "DTEND;TZID=%s:%s\r\n",
			c.loc.String(), c.events[i].End.Format(TimeFormat),
		)
		if err != nil {
			return err
		}
		_, err = fmt.Fprint(writer, "END:VEVENT\r\n")
		if err != nil {
			return err
		}
	}
	_, err = fmt.Fprint(writer, "END:VCALENDAR\r\n")
	return err
}
