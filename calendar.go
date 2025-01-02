package goical

import (
	"io"
	"text/template"
	"time"
)

const TimeFormat = "20060102T030405"
const DefaultDuration = time.Hour

type Calendar struct {
	loc    *time.Location
	Events []Event
}

func (c *Calendar) AddEvent(input Event) *Calendar {
	c.Events = append(c.Events, input)
	return c
}

func (c *Calendar) Render(writer io.Writer) error {
	if c.loc == nil {
		c.loc = time.Local
	}
	tpl, err := template.New("calendar").Parse(tplMain)
	if err != nil {
		return err
	}
	events := make([]tplEvent, len(c.Events))
	for i := range c.Events {
		events[i] = c.convert(&c.Events[i])
	}
	data := tplData{
		TimeZoneID: c.loc.String(),
		Events:     events,
	}
	return tpl.Execute(writer, data)
}
