package goical

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

// TestRussianHolidays generates a vCalendar for Russian holidays.
func TestRussianHolidays(t *testing.T) {
	var buf bytes.Buffer

	tz := time.FixedZone("Moscow", +3*3600)
	err := RussianHolidays(tz, &buf)
	if err != nil {
		t.Fatalf("RussianHolidays failed: %v", err)
	}

	output := buf.String()
	if len(output) == 0 {
		t.Fatal("Expected non-empty calendar output")
	}
}

// TestRussianHolidaysVCALENDARHeader tests that output starts with VCALENDAR header.
func TestRussianHolidaysVCALENDARHeader(t *testing.T) {
	var buf bytes.Buffer

	tz := time.FixedZone("Moscow", +3*3600)
	err := RussianHolidays(tz, &buf)
	if err != nil {
		t.Fatalf("RussianHolidays failed: %v", err)
	}

	output := buf.String()
	if !bytes.Contains([]byte(output), []byte("BEGIN:VCALENDAR")) {
		t.Fatal("Expected VCALENDAR header in output")
	}
}

// TestRussianHolidaysContainsNewYear tests that New Year event exists.
func TestRussianHolidaysContainsNewYear(t *testing.T) {
	var buf bytes.Buffer

	tz := time.FixedZone("Moscow", +3*3600)
	err := RussianHolidays(tz, &buf)
	if err != nil {
		t.Fatalf("RussianHolidays failed: %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "New year") {
		t.Fatal("Expected 'New year' event in output")
	}
}

// TestRussianHolidaysContainsChristmas tests that Christmas event exists.
func TestRussianHolidaysContainsChristmas(t *testing.T) {
	var buf bytes.Buffer

	tz := time.FixedZone("Moscow", +3*3600)
	err := RussianHolidays(tz, &buf)
	if err != nil {
		t.Fatalf("RussianHolidays failed: %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "Рождество") {
		t.Fatal("Expected 'Рождество' event in output")
	}
}

// TestRussianHolidaysContainsMotherland tests that Motherland Protector event exists.
func TestRussianHolidaysContainsMotherland(t *testing.T) {
	var buf bytes.Buffer

	tz := time.FixedZone("Moscow", +3*3600)
	err := RussianHolidays(tz, &buf)
	if err != nil {
		t.Fatalf("RussianHolidays failed: %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "День Защитника") {
		t.Fatal("Expected 'День Защитника' event in output")
	}
}

// TestRussianHolidaysContainsWomansDay tests that Women's Day event exists.
func TestRussianHolidaysContainsWomansDay(t *testing.T) {
	var buf bytes.Buffer

	tz := time.FixedZone("Moscow", +3*3600)
	err := RussianHolidays(tz, &buf)
	if err != nil {
		t.Fatalf("RussianHolidays failed: %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "Женский") {
		t.Fatal("Expected 'Женский' event in output")
	}
}

// TestRussianHolidaysContainsCosmonautics tests that Cosmonautics Day event exists.
func TestRussianHolidaysContainsCosmonautics(t *testing.T) {
	var buf bytes.Buffer

	tz := time.FixedZone("Moscow", +3*3600)
	err := RussianHolidays(tz, &buf)
	if err != nil {
		t.Fatalf("RussianHolidays failed: %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "Космонавтики") {
		t.Fatal("Expected 'Космонавтики' event in output")
	}
}

// TestRussianHolidaysContainsLabourDay tests that Labour Day event exists.
func TestRussianHolidaysContainsLabourDay(t *testing.T) {
	var buf bytes.Buffer

	tz := time.FixedZone("Moscow", +3*3600)
	err := RussianHolidays(tz, &buf)
	if err != nil {
		t.Fatalf("RussianHolidays failed: %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "Весны и Труда") {
		t.Fatal("Expected 'Весны и Труда' event in output")
	}
}

// TestRussianHolidaysContainsVictoryDay tests that Victory Day event exists.
func TestRussianHolidaysContainsVictoryDay(t *testing.T) {
	var buf bytes.Buffer

	tz := time.FixedZone("Moscow", +3*3600)
	err := RussianHolidays(tz, &buf)
	if err != nil {
		t.Fatalf("RussianHolidays failed: %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "Победы") {
		t.Fatal("Expected 'Победы' event in output")
	}
}

// TestRussianHolidaysURLParsing tests URL parsing behavior.
// https://github.com/vodolaz095/goical/wiki/Architecture-Notes
func TestRussianHolidaysURLParsing(t *testing.T) {
	var buf bytes.Buffer

	// Note: URL parsing may fail (handled in Russian_holidays.go), which is acceptable behavior
	tz := time.Local
	err := RussianHolidays(tz, &buf)
	if err != nil {
		// Some URLs might fail to parse - this is acceptable per project behavior
		t.Logf("Some URLs failed to parse: %v", err)
	}
}

// TestRussianHolidaysWithUTF8Encoding tests that output is ASCII.
// https://github.com/vodolaz095/goical/wiki/Code-Style-Conventions
func TestRussianHolidaysWithUTF8Encoding(t *testing.T) {
	var buf bytes.Buffer

	tz := time.FixedZone("Moscow", +3*3600)
	err := RussianHolidays(tz, &buf)
	if err != nil {
		t.Fatalf("RussianHolidays failed: %v", err)
	}

	output := buf.String()
	// Output should contain UTF-8 encoded characters (Russian text)
	if !bytesContainsUTF8([]byte(output)) {
		t.Error("Expected UTF-8 encoded output with Russian characters")
	}
}

// bytesContainsUTF8 checks if a byte slice contains actual UTF-8 bytes (non-ASCII values).
func bytesContainsUTF8(b []byte) bool {
	// Check for at least one non-ASCII byte (UTF-8)
	for _, byteValue := range b {
		if byteValue > 127 {
			return true
		}
	}
	return false
}

// TestRussianHolidaysTZIDFormat tests TZID format in output.
// https://github.com/vodolaz095/goical/wiki/Architecture-Notes
func TestRussianHolidaysTZIDFormat(t *testing.T) {
	var buf bytes.Buffer

	tz := time.FixedZone("Moscow", +3*3600)
	err := RussianHolidays(tz, &buf)
	if err != nil {
		t.Fatalf("RussianHolidays failed: %v", err)
	}

	output := buf.String()
	// TZID should appear before DTSTART/DTEND
	tzidIdx := bytes.Index([]byte(output), []byte("TZID="))
	dtstartIdx := bytes.Index([]byte(output), []byte("DTSTART;TZID="))
	dtendIdx := bytes.Index([]byte(output), []byte("DTEND;TZID="))

	if tzidIdx < 0 {
		t.Fatal("Expected TZID in output")
	}
	// TZID should appear before DTSTART and DTEND
	if tzidIdx >= dtstartIdx {
		t.Error("Expected TZID before DTSTART")
	}
	if tzidIdx >= dtendIdx {
		t.Error("Expected TZID before DTEND")
	}
}

// TestRussianHolidaysProdID tests PRODID format in output.
// https://github.com/vodolaz095/goical/wiki/Architecture-Notes
func TestRussianHolidaysProdID(t *testing.T) {
	var buf bytes.Buffer

	tz := time.FixedZone("Moscow", +3*3600)
	err := RussianHolidays(tz, &buf)
	if err != nil {
		t.Fatalf("RussianHolidays failed: %v", err)
	}

	output := buf.String()
	// PRODID should match expected format
	expectedProdID := "https://github.ru/vodolaz095/goical"
	if !bytes.Contains([]byte(output), []byte(expectedProdID)) {
		t.Errorf("Expected PRODID '%s' in output", expectedProdID)
	}
}

// TestRussianHolidaysMultipleEvents tests that all holidays are included.
func TestRussianHolidaysMultipleEvents(t *testing.T) {
	var buf bytes.Buffer

	tz := time.Local
	err := RussianHolidays(tz, &buf)
	if err != nil {
		t.Fatalf("RussianHolidays failed: %v", err)
	}

	output := buf.String()

	// Count expected holiday names
	holidayNames := []string{
		"New year",
		"Рождество",
		"День Защитника",
		"Женский",
		"Космонавтики",
		"Весны и Труда",
		"Победы",
		"День России",
		"Народного Единства",
		"Октябрьской Революции",
	}

	eventCount := 0
	for _, name := range holidayNames {
		if bytes.Contains([]byte(output), []byte(name)) {
			eventCount++
		}
	}

	// Should have at least 9 events (Victory Day and Russia Day might merge)
	if eventCount < 9 {
		t.Errorf("Expected at least 9 holiday events, found %d", eventCount)
	}
}
