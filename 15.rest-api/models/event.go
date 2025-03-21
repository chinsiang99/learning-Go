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

// Custom validation error messages for Event using JSON field names
func (e *Event) ValidationMessages() map[string]string {
	return map[string]string{
		"Title":       "Event title cannot be empty.",
		"Description": "Please provide a valid event description.",
		"Location":    "Location must be specified.",
		"DateTime":    "Provide a valid date and time for the event.",
	}
}

func (e *Event) FieldStructTagMapping() map[string]string {
	return map[string]string{
		"Title":       "title",
		"Description": "description",
		"Location":    "location",
		"DateTime":    "dateTime",
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
