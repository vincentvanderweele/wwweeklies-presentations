package presentation

import (
	"strconv"

	"github.com/vincentvanderweele/wwweeklies-presentations/generated/model"
	"github.com/vincentvanderweele/wwweeklies-presentations/generated/router"
)

type presentationStorage struct {
	presentations model.ReadOnlyPresentations
	nextID        int
}

// NewStorage creates a new presentation storage
func NewStorage() router.Handler {
	return &presentationStorage{
		presentations: model.ReadOnlyPresentations{},
		nextID:        0,
	}
}

func (s *presentationStorage) GetPresentations() (model.ReadOnlyPresentations, error) {
	return s.presentations, nil
}

func (s *presentationStorage) CreatePresentation(newPresentation model.Presentation) (model.ReadOnlyPresentation, error) {
	presentation := model.ReadOnlyPresentation{
		ID:           s.getNextID(),
		Presentation: newPresentation,
	}
	s.presentations = append(s.presentations, presentation)

	return presentation, nil
}

func (s *presentationStorage) GetPresentation(id string) (presentation model.ReadOnlyPresentation, err error) {
	var p *model.ReadOnlyPresentation
	if p, _, err = s.getPresentation(id); err == nil {
		presentation = *p
	}
	return
}

func (s *presentationStorage) UpdatePresentation(id string, update model.Presentation) (err error) {
	var p *model.ReadOnlyPresentation
	if p, _, err = s.getPresentation(id); err == nil {
		p.Presentation = update
	}
	return
}

func (s *presentationStorage) DeletePresentation(id string) error {
	if _, i, err := s.getPresentation(id); err == nil {
		s.presentations[i] = s.presentations[len(s.presentations)-1]
		s.presentations = s.presentations[:len(s.presentations)-1]
	}
	return nil
}

func (s *presentationStorage) getNextID() *string {
	id := strconv.Itoa(s.nextID)
	s.nextID++
	return &id
}

func (s *presentationStorage) getPresentation(id string) (*model.ReadOnlyPresentation, int, error) {
	for i := range s.presentations {
		if *s.presentations[i].ID == id {
			return &s.presentations[i], i, nil
		}
	}

	return nil, 0, model.NotFoundError("")
}
