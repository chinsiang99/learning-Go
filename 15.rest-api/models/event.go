package models

import "time"

type Event struct {
	ID          string    `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"dateTime" binding:"required"`
	UserID      int       `json:"userId"`
}

// Custom validation error messages for Event
func (e *Event) ValidationMessages() map[string]string {
	return map[string]string{
		"title":       "Event title cannot be empty.",
		"description": "Please provide a valid event description.",
		"location":    "Location must be specified.",
		"dateTime":    "Provide a valid date and time for the event.",
	}
}

var events = []Event{}

func (e *Event) Save() {
	// later: add it to database
	events = append(events, *e)
}

func GetAllEvents() []Event {
	return events
}

// 🧑‍💻 General Rules of Thumb
// If a method needs to modify the struct → Use a pointer receiver.
// If a method only reads data → Use a value receiver.
// If the struct has more than 3-4 fields → Prefer pointer receivers for efficiency.
// If the struct is immutable or for simple data like coordinates → Use a value receiver.
