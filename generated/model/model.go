package model

import "time"

// This is a generated file
// Manual changes will be overwritten

// Presentation A presentation at the WWWeeklies
type Presentation struct {
	Date     *time.Time `json:"date" db:"date"`
	Duration *int64     `json:"duration" db:"duration"`
	Remarks  *string    `json:"remarks" db:"remarks"`
	Site     *string    `json:"site" db:"site"`
	Speaker  *string    `json:"speaker" db:"speaker"`
	Title    *string    `json:"title" db:"title"`
}

// NewPresentation returns a new Presentation
func NewPresentation(date time.Time, duration int64, remarks string, site string, speaker string, title string) Presentation {
	return Presentation{
		&date,
		&duration,
		&remarks,
		&site,
		&speaker,
		&title,
	}
}

// ReadOnlyPresentation A presentation at the WWWeeklies
type ReadOnlyPresentation struct {
	Presentation
	ID *string `json:"id" db:"id"`
}

// NewReadOnlyPresentation returns a new ReadOnlyPresentation
func NewReadOnlyPresentation(date time.Time, duration int64, id string, remarks string, site string, speaker string, title string) ReadOnlyPresentation {
	return ReadOnlyPresentation{
		NewPresentation(date, duration, remarks, site, speaker, title),
		&id,
	}
}

// Presentations A list of presentations
type Presentations []Presentation

// ReadOnlyPresentations A list of presentations
type ReadOnlyPresentations []ReadOnlyPresentation
