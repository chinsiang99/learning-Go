package models

import "time"

type Event struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	DateTime    time.Time `json:"dateTime"`
	UserID      int       `json:"userId"`
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
