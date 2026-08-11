package goical

import (
	"bytes"
	"testing"
	"time"
)

// TestEmptyUID tests that events with empty UID are ignored.
// https://github.com/vodolaz095/goical/wiki/Control-Flow
func TestEmptyUID(t *testing.T) {
	calendar := New(time.Local)
	emptyUids := []string{}

	for _, uid := range emptyUids {
		calendar.AddEvent(Event{
			UID:     uid,
			Summary: "Test Meeting",
			Start:   time.Unix(1704067200, 0),
			End:     time.Unix(1704070800, 0),
		})
	}

	// Verify event was not added
	if len(calendar.events) != 0 {
		t.Errorf("Expected 0 events, got %d", len(calendar.events))
	}
}

// TestZeroStart tests that events with zero Start time are ignored.
func TestZeroStart(t *testing.T) {
	calendar := New(time.Local)
	var zeroStartTime time.Time

	calendar.AddEvent(Event{
		UID:     "test-uuid",
		Start:   zeroStartTime,
		End:     time.Unix(1704070800, 0),
		Summary: "Zero Start",
	})

	if len(calendar.events) != 0 {
		t.Errorf("Expected 0 events, got %d", len(calendar.events))
	}
}

// TestZeroEnd tests that events with zero End time are ignored.
func TestZeroEnd(t *testing.T) {
	calendar := New(time.Local)
	var zeroEndTime time.Time

	calendar.AddEvent(Event{
		UID:     "test-uuid",
		Start:   time.Unix(1704067200, 0),
		End:     zeroEndTime,
		Summary: "Zero End",
	})

	if len(calendar.events) != 0 {
		t.Errorf("Expected 0 events, got %d", len(calendar.events))
	}
}

// TestStartAfterEnd tests that events with Start after End are ignored.
// https://github.com/vodolaz095/goical/wiki/Control-Flow
func TestStartAfterEnd(t *testing.T) {
	calendar := New(time.Local)

	calendar.AddEvent(Event{
		UID:     "test-uuid",
		Summary: "Start After End",
		Start:   time.Unix(1704070800, 0), // End time
		End:     time.Unix(1704067200, 0), // Start time
	})

	if len(calendar.events) != 0 {
		t.Errorf("Expected 0 events, got %d", len(calendar.events))
	}
}

// TestValidStartAfterEnd confirms valid events pass the check.
func TestValidStartAfterEnd(t *testing.T) {
	calendar := New(time.Local)

	calendar.AddEvent(Event{
		UID:     "valid-123",
		Summary: "Valid Event",
		Start:   time.Unix(1704067200, 0),
		End:     time.Unix(1704070800, 0),
	})

	if len(calendar.events) != 1 {
		t.Errorf("Expected 1 event, got %d", len(calendar.events))
	}
}

// TestInvalidUIDs tests invalid UID values.
func TestInvalidUIDs(t *testing.T) {
	calendar := New(time.Local)

	// Empty UID is rejected per AddEvent() validation in calendar.go:25-27
	calendar.AddEvent(Event{
		UID:     "",
		Summary: "Empty UID",
		Start:   time.Unix(1704067200, 0),
		End:     time.Unix(1704070800, 0),
	})
	// Expect 0 events after empty UID
	if len(calendar.events) != 0 {
		t.Errorf("Empty UID should be rejected, got %d events", len(calendar.events))
	}

	// Valid non-empty UID is accepted
	secondCalendar := New(time.Local)
	secondCalendar.AddEvent(Event{
		UID:     "valid-uid-123",
		Summary: "Valid UID",
		Start:   time.Unix(1704067200, 0),
		End:     time.Unix(1704070800, 0),
	})
	if len(secondCalendar.events) != 1 {
		t.Errorf("Valid UID should be accepted, got %d events", len(secondCalendar.events))
	}
}

// TestZeroTimestamp tests default timestamp handling.
// https://github.com/vodolaz095/goical/wiki/Control-Flow
func TestZeroTimestamp(t *testing.T) {
	calendar := New(time.Local)
	fixedTimestamp := time.Unix(1704074400, 0)

	calendar.AddEvent(Event{
		UID:       "test-uuid",
		Summary:   "No Timestamp",
		Start:     time.Unix(1704067200, 0),
		End:       time.Unix(1704070800, 0),
		Timestamp: fixedTimestamp, // Set to fixed value
	})

	if len(calendar.events) != 1 {
		t.Errorf("Expected 1 event, got %d", len(calendar.events))
	}
	if calendar.events[0].Timestamp != fixedTimestamp {
		t.Errorf("Expected fixed timestamp, got %v", calendar.events[0].Timestamp)
	}
}

// TestValidEvent tests that valid events are added correctly.
func TestValidEvent(t *testing.T) {
	calendar := New(time.Local)

	calendar.AddEvent(Event{
		UID:         "meet-2024-001",
		Summary:     "Weekly Team Meeting",
		Description: "Discuss project updates and next sprint goals",
		Start:       time.Unix(1704067200, 0),
		End:         time.Unix(1704070800, 0),
		Organizer: Person{
			CommonName: "Alice Smith",
			Email:      "alice@example.com",
		},
	})

	if len(calendar.events) != 1 {
		t.Errorf("Expected 1 event, got %d", len(calendar.events))
	}
	if calendar.events[0].UID != "meet-2024-001" {
		t.Errorf("Expected UID 'meet-2024-001', got '%s'", calendar.events[0].UID)
	}
}

// TestRenderEmptyCalendar tests rendering an empty calendar.
func TestRenderEmptyCalendar(t *testing.T) {
	calendar := New(time.Local)

	var buf bytes.Buffer
	err := calendar.Render(&buf)
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}

	// Should have VCALENDAR header (15 chars + CR/LF ending)
	if len(buf.Bytes()) < 16 {
		t.Fatalf("Expected VCALENDAR header in output, got len=%d", len(buf.Bytes()))
	}
	// Check first 15 characters match BEGIN:VCALENDAR
	expectedPrefix := [15]byte{'B', 'E', 'G', 'I', 'N', ':', 'V', 'C', 'A', 'L', 'E', 'N', 'D', 'A', 'R'}
	for i, c := range expectedPrefix {
		if buf.Bytes()[i] != byte(c) {
			t.Fatalf("Expected %q at position %d, got %q", c, i, buf.Bytes()[i])
		}
	}
	// The next character should be either \r or \n (LF line ending)
	if buf.Bytes()[15] != '\r' && buf.Bytes()[15] != '\n' {
		t.Fatalf("Expected line ending after VCALENDAR header, got %q", buf.Bytes()[15])
	}
}

// TestRenderValidEvent tests rendering a valid event.
func TestRenderValidEvent(t *testing.T) {
	calendar := New(time.Local)
	calendar.AddEvent(Event{
		UID:     "meeting-001",
		Summary: "Team Standup",
		Start:   time.Unix(1704067200, 0),
		End:     time.Unix(1704070800, 0),
	})

	var buf bytes.Buffer
	err := calendar.Render(&buf)
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}

	expected := "BEGIN:VEVENT\r\nUID:meeting-001\r\nSUMMARY:Team Standup\r\n"
	if !bytes.Contains([]byte(buf.String()), []byte(expected)) {
		t.Errorf("Expected event in output, got '%s'", buf.String())
	}
}

// TestRenderMultipleEvents tests rendering multiple events in order.
func TestRenderMultipleEvents(t *testing.T) {
	calendar := New(time.Local)

	calendar.AddEvent(Event{
		UID:     "meeting-002",
		Summary: "Sprint Planning",
		Start:   time.Unix(1704060000, 0), // Earlier
		End:     time.Unix(1704063600, 0),
	})

	calendar.AddEvent(Event{
		UID:     "meeting-001",
		Summary: "Team Standup",
		Start:   time.Unix(1704067200, 0), // Later
		End:     time.Unix(1704070800, 0),
	})

	var buf bytes.Buffer
	err := calendar.Render(&buf)
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}

	output := buf.String()
	// First event should appear first
	if stringsIndex := bytes.Index([]byte(output), []byte("UID:meeting-002")); stringsIndex < 0 || stringsIndex%len("\r\n") != 0 {
		t.Fatal("Expected meeting-002 (earlier) before meeting-001 (later) in output, got: " + string(output))
	}
}

// TestRenderTZID tests TZID format in output.
// https://github.com/vodolaz095/goical/wiki/Architecture-Notes
func TestRenderTZID(t *testing.T) {
	calendar := New(time.Local)
	calendar.AddEvent(Event{
		UID:     "test-001",
		Summary: "Timezone Test",
		Start:   time.Unix(1704067200, 0),
		End:     time.Unix(1704070800, 0),
	})

	var buf bytes.Buffer
	err := calendar.Render(&buf)
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}

	output := buf.String()
	// TZID should appear before DTSTART and DTEND
	tzidIdx := bytes.Index([]byte(output), []byte("TZID="))
	dtstartIdx := bytes.Index([]byte(output), []byte("DTSTART;TZID="))
	dtendIdx := bytes.Index([]byte(output), []byte("DTEND;TZID="))

	if tzidIdx < 0 || dtstartIdx < 0 || dtendIdx < 0 {
		t.Errorf("Expected TZID in output, got '%s'", output)
	}

	if tzidIdx >= dtstartIdx || tzidIdx >= dtendIdx {
		t.Errorf("Expected TZID before DTSTART/DTEND, got TZID=%d, DTSTART=%d, DTEND=%d",
			tzidIdx, dtstartIdx, dtendIdx)
	}
}
