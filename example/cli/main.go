package main

import (
	"log"
	"net/url"
	"os"
	"time"

	"github.com/vodolaz095/goical"
)

func main() {
	tz, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		log.Fatalf("error loading location: %s", err)
	}
	calendar := goical.New(tz)
	now := time.Now()
	u, _ := url.Parse("http://example.org")
	calendar.AddEvent(goical.Event{
		// mandatory fields
		UID:   "today_with_john_doe",
		Start: now.Add(time.Minute),
		End:   now.Add(time.Hour),
		// optional fields
		Timestamp:   now,
		Summary:     "Some important meeting summary",
		Description: "Some important meeting description",
		Location:    "nowhere",
		URL:         u,
		Organizer: goical.Person{
			CommonName: "John Doe",
			Email:      "john.doe@example.org",
		},
	})

	err = calendar.Render(os.Stdout)
	if err != nil {
		log.Fatalf("error rendering: %s", err)
	}
}
