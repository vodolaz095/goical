# AGENTS.md

## Project Overview

`goical` is a minimal Go library for generating vCalendar (iCalendar) files. It provides a simple API for creating calendar events and rendering them as RFC 5545 compliant calendar data suitable for consumption by calendar clients like Mozilla Thunderbird.

**Repository**: https://github.com/vodolaz095/goical  
**Module**: `github.com/vodolaz095/goical`  
**Version**: Go 1.26

---

## Essential Commands

### Build and Run
```bash
# Run examples from the repository (no build step needed)
go run example/cli/main.go
go run example/holidays/main.go
go run example/http/main.go
```

### Lint
```bash
# Ensure all linting tools are installed
make tools

# Run all linters (gofmt, golint, go vet, staticcheck)
make lint
```

### Package Usage
Should be installed by `go get -u github.com/vodolaz095/ical` and imported as dependency using

```go

import "github.com/vodolaz095/ical"

```

---

## Code Organization

### Package Structure (`/home/vodolaz095/projects/goical/)
- **calendar.go**: Core `Calendar` struct, `AddEvent()` method, and `Render()` vCalendar output
- **event.go**: `Event` struct definition and `Person` struct with `String()` formatter
- **Russian_holidays.go**: Presigned holiday event generator for Russian holidays

### Example Applications (`/home/vodolaz095/projects/goical/example/`)
- **cli/**: CLI tool demonstrating manual calendar creation
- **holidays/**: Example generating Russian holidays calendar
- **http/**: HTTP server returning calendar via REST endpoint

---

## API Reference

### Core Types
```go
// Calendar - manages and renders events to vCalendar format
type Calendar struct {
    loc *time.Location
    events []Event
}

// Event - single calendar event (RFC 5545 compliant)
type Event struct {
    UID         string        // mandatory - unique identifier
    Timestamp   time.Time     // when event was created
    Summary     string        // short human-readable name
    Description string        // long description
    Location    string        // physical location
    URL         *url.URL      // event URL
    Organizer   Person
    Start       time.Time     // mandatory start time
    End         time.Time     // mandatory end time
}

// Person - organizer contact info
type Person struct {
    CommonName string
    Email      string
}
```

### Methods

| Method | Purpose |
|--------|---------|
| `New(loc *time.Location)` | Creates new `Calendar` instance with timezone |
| `AddEvent(e Event) *Calendar` | Adds event (returns self for chaining; validates UID, Start/End, Timestamp) |
| `Render(w io.Writer) error` | Renders vCalendar to writer; outputs sorted by start time |

### Key Constants & Patterns
- `TimeFormat = "20060102T150405"` - UTC timestamp format used for DTSTART/DTEND
- Events filtered in `Render()`: UID must not be empty, Start/End/Timestamp must not be zero
- Timezone awareness via `TZID` parameter in iCalendar output
- Events sorted by `Start` time before rendering

---

## Usage Patterns

### 1. Manual Calendar Creation
```go
calendar := goical.New(tz)
calendar.AddEvent(Event{
    UID:      "my-meeting-001",
    Summary:   "Team Standup",
    Start:     now,
    End:       now.Add(time.Hour),
    // ... other fields
})
err := calendar.Render(os.Stdout)
```

### 2. HTTP Endpoint for Calendar
```go
http.HandleFunc("/calendar", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/calendar")
    err := goical.RussianHolidays(tz, w)
    if err != nil {
        log.Fatal(err)
    }
})
```

---

## Architecture Notes

### Control Flow
1. User creates `Calendar` with optional timezone
2. User calls `AddEvent()` to queue events (validation happens here)
3. `Render()` sorts events and streams vCalendar to writer

### Data Flow in `Render()`
1. Writes iCalendar header: `BEGIN:VCALENDAR VERSION:2.0 CALSCALE:GREGORIAN PRODID:...`
2. Sorts events by `Start` time
3. Filters out zero-valued Start/End/Timestamp events
4. For each event: writes header fields (UID, Summary, Description, etc.)
5. Adds `TZID` prefix to DTSTART/DTEND/DTSTAMP
6. Writes event footer: `END:VEVENT`
7. Writes calendar footer: `END:VCALENDAR`

### Timezones
- Calendar stores timezone in `*time.Location`
- Output uses `TZID=...` parameter before timestamp fields
- If no `Location` is set, `time.Local` is used; otherwise `TZID=time.Location.String()`
- This is critical for clients that display times in local timezone

---

## Testing Strategy

This repository currently **does not have automated tests**. Tests would need to be added manually using Go's standard `go test` framework.

---

## Code Style & Conventions

### Go Version
- **Go 1.26** as required in `go.mod`

### Import Patterns
- All imports are from standard library: `fmt`, `io`, `net/url`, `sort`, `time`
- No external dependencies

### File Conventions
- Core types and methods in same file when logical (e.g., `Event` type with `Person` type in `event.go`)
- Single-purpose functions per file when they don't define types

### Comments
- RFC 5545 URLs documented in `event.go` before `Person` type
- Inline comments explain field semantics

---

## Common Gotchas

### 1. Timezone Display
**Non-obvious**: The `TZID` parameter appears before `DTSTART` and `DTEND`, not after timestamps. This is iCalendar RFC 5545 specification.

```icalendar
DTSTART;TZID=Europe/Moscow:20240101T000000
DTEND;TZID=Europe/Moscow:20240101T010000
```

### 2. Event Validation in `AddEvent()`
- `UID` must not be empty, otherwise event is ignored
- `Start` and `End` must be non-zero (valid times)
- `Timestamp` defaults to `time.Now()` if zero, otherwise ignored
- Events with `Start.After(End)` are silently ignored (validation fails in `AddEvent()`)

### 3. Output Format
- Events are **sorted by start time** before rendering
- Output is **ASCII** (C8 locale), not UTF-8 encoded strings
- Empty `Summary` and `Description` fields are **not output** (only non-blank strings appear in iCalendar)

### 4. URL Handling
- `URL` is a pointer to `url.URL`
- If `URL` is `nil`, no URL field appears in vCalendar output
- URLs for Russian holidays are pre-parsed with zero-error handling

### 5. vCalendar Schema
- Schema is **2.0**, not 1.0 (minor breaking change)
- Uses **Gregorian calendar**, not ISO-8601
- `PRODID` contains repository URL: `https://github.ru/vodolaz095/goical`

---

## Deployment

### As Binary
```bash
go build -o goical .
```

### As Package
```bash
go get github.com/vodolaz095/goical
```

---

## Contributing Guidelines

### Adding New Holiday Support
1. Add events to `Russian_holidays.go` following pattern:
   ```go
   holidays.AddEvent(Event{
       UID:    fmt.Sprintf("holiday_name_%v", now.Year()),
       Summary: "Holiday name",
       // ... other fields
   })
   ```
2. Format events chronologically (sorted in `Render()`)
3. Update if needed: `TimeFormat`, iCalendar header fields, or sorting logic

### Adding New Event Types
1. Ensure events have unique `UID` strings
2. Validate input (non-zero Start/End/Timestamp, non-empty UID)
3. Follow existing pattern in `AddEvent()` validation

---

## License

MIT License (MIT)

Copyright (c) 2025 Ostroumov Anatolij <ostroumov095 at gmail dot com>

---

## Related Project Info

- **Funding**: Listed in `.github/FUNDING.yml` - check for sponsor info
- **CI/CD**: No automated CI configured
- **Documentation**: Minimal, example usage in README.md
