package model

import "fmt"

// This is a generated file
// Manual changes will be overwritten

// Validate validates a Presentation based on the swagger spec
func (s *Presentation) Validate() (errors []string) {

	if s.Date == nil {
		errors = append(errors, "date is required")
	} else {
	}

	if s.Duration == nil {
		errors = append(errors, "duration is required")
	} else {

		if *s.Duration > 60 {
			errors = append(errors, "duration should be at most 60")
		}

		if *s.Duration < 0 {
			errors = append(errors, "duration should be at least 0")
		}
	}

	if s.Site == nil {
		errors = append(errors, "site is required")
	} else {

		switch *s.Site {
		case "Helsinki", "Tampere", "London", "Berlin", "Munich", "Stockholm", "External": // ok
		default:
			errors = append(errors, fmt.Sprintf("%s is not an allowed value for site", *s.Site))
		}
	}

	if s.Speaker == nil {
		errors = append(errors, "speaker is required")
	} else {

		if len(*s.Speaker) < 1 {
			errors = append(errors, "speaker should be no shorter than 1 characters")
		}
	}

	if s.Title == nil {
		errors = append(errors, "title is required")
	} else {

		if len(*s.Title) < 1 {
			errors = append(errors, "title should be no shorter than 1 characters")
		}
	}
	return
}

// Validate validates a ReadOnlyPresentation based on the swagger spec
func (s *ReadOnlyPresentation) Validate() (errors []string) {
	if e := s.Presentation.Validate(); len(e) > 0 {
		errors = append(errors, e...)
	}

	if s.ID == nil {
		errors = append(errors, "id is required")
	} else {
	}
	return
}

// Validate validates a Presentations based on the swagger spec
func (s *Presentations) Validate() (errors []string) {

	for _, elt := range *s {
		if e := elt.Validate(); len(e) > 0 {
			errors = append(errors, e...)
		}
	}
	return
}

// Validate validates a ReadOnlyPresentations based on the swagger spec
func (s *ReadOnlyPresentations) Validate() (errors []string) {

	for _, elt := range *s {
		if e := elt.Validate(); len(e) > 0 {
			errors = append(errors, e...)
		}
	}
	return
}
