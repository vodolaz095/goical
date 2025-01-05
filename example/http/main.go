package main

import (
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/vodolaz095/goical"
)

func main() {
	tz, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		log.Fatalf("error loading location: %s", err)
	}

	http.HandleFunc("/holidays", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/calendar")
		err := goical.RussianHolidays(tz, writer)
		if err != nil {
			log.Fatalf("error rendering: %s", err)
		}
	})

	http.HandleFunc("/today", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/calendar")
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
			Description: "Some important meeting summary",
			Location:    "nowhere",
			URL:         u,
			Organizer: goical.Person{
				CommonName: "John Doe",
				Email:      "john.doe@example.org",
			},
		})

		errH := calendar.Render(writer)
		if errH != nil {
			log.Fatalf("error rendering: %s", errH)
		}
	})

	log.Fatal(http.ListenAndServe(":3000", http.DefaultServeMux))
}
