package persons

import "google.golang.org/genproto/googleapis/type/date"

// Person - Models a human Person
type Person struct {
	Name        string
	DOB         date.Date
	Gender      Gender
	Nationality string
}

// Gender - gender
type Gender int

const (
	// Male - male
	Male Gender = iota
	// Female - female
	Female
	// Transwoman - Transwoman
	Transwoman
	// Transman - Transman
	Transman
	// NonBinary - NonBinary
	NonBinary
)
