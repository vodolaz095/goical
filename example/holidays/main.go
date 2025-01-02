package main

import (
	"log"
	"os"
	"time"

	"github.com/vodolaz095/goical"
)

func main() {
	tz, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		log.Fatalf("error loading location: %s", err)
	}
	err = goical.RussianHolidays(tz, os.Stdout)
	if err != nil {
		log.Fatalf("error rendering: %s", err)
	}
}
