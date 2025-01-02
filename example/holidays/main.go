package main

import (
	"log"
	"os"

	"github.com/vodolaz095/goical"
)

func main() {
	err := goical.RussianHolidays(os.Stdout)
	if err != nil {
		log.Fatalf("error rendering: %s", err)
	}
}
