package router

// This is a generated file
// Manual changes will be overwritten

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"

	"github.com/vincentvanderweele/wwweeklies-presentations/generated/model"
)

// Handler implements the actual functionality of the service
type Handler interface {
	// Other
	GetPresentations() (model.ReadOnlyPresentations, error)
	CreatePresentation(bodyPresentation model.Presentation) (model.ReadOnlyPresentation, error)
	GetPresentation(pathPresentationID string) (model.ReadOnlyPresentation, error)
	UpdatePresentation(pathPresentationID string, bodyUpdate model.Presentation) error
	DeletePresentation(pathPresentationID string) error
}

// ErrorTransformer transforms errors in standard format into the format according to the swagger spec
type ErrorTransformer interface {
	ValidationErrorsToInvalidDataError(errs []string) model.InvalidDataError
	ErrorToString(err error) string
}

type middleware struct {
	handler          Handler
	errorTransformer ErrorTransformer
}

// NewServer creates a http handler with a router for all methods of the service
func NewServer(handler Handler, errorTransformer ErrorTransformer) http.Handler {
	m := &middleware{
		handler,
		errorTransformer,
	}

	router := httprouter.New()

	router.GET("/presentations", m.getPresentations)
	router.POST("/presentations", m.createPresentation)
	router.GET("/presentations/:presentationId", m.getPresentation)
	router.PUT("/presentations/:presentationId", m.updatePresentation)
	router.DELETE("/presentations/:presentationId", m.deletePresentation)

	return router
}

func (m *middleware) getPresentations(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	errorTransformer := func(err error) interface{} { return m.errorTransformer.ErrorToString(err) }

	defer func() {
		if recovered := recover(); recovered != nil {
			err := errors.New("Recovered")
			log.WithField("error", recovered).Error(err)
			respondJSON(w, m.errorTransformer.ErrorToString(err), "string", http.StatusInternalServerError, errorTransformer)
		}
	}()

	var (
		result model.ReadOnlyPresentations
		err    error
		errs   []string
	)

	if result, err = m.handler.GetPresentations(); err != nil {
		switch err.(type) {
		default:
			respondJSON(w, m.errorTransformer.ErrorToString(err), "string", http.StatusInternalServerError, errorTransformer)
		}
		return
	}

	if errs = result.Validate(); len(errs) > 0 {
		err := errors.New("Invalid response data")
		log.WithFields(log.Fields{
			"dataType": "ReadOnlyPresentations",
			"error":    strings.Join(errs, "\n"),
		}).Error(err)
		respondJSON(w, m.errorTransformer.ErrorToString(err), "string", http.StatusInternalServerError, errorTransformer)
		return
	}

	respondJSON(w, result, "ReadOnlyPresentations", http.StatusOK, errorTransformer)
}

func (m *middleware) createPresentation(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	errorTransformer := func(err error) interface{} { return m.errorTransformer.ErrorToString(err) }

	defer func() {
		if recovered := recover(); recovered != nil {
			err := errors.New("Recovered")
			log.WithField("error", recovered).Error(err)
			respondJSON(w, m.errorTransformer.ErrorToString(err), "string", http.StatusInternalServerError, errorTransformer)
		}
	}()

	var (
		result model.ReadOnlyPresentation
		err    error
		errs   []string
	)

	var bodyPresentation model.Presentation
	if err = json.NewDecoder(r.Body).Decode(&bodyPresentation); err != nil {
		errs = append(errs, err.Error())
		log.WithFields(log.Fields{
			"bodyType": "Presentation",
			"error":    err,
		}).Error("Failed to parse body data")
	} else if e := bodyPresentation.Validate(); len(e) > 0 {
		errs = append(errs, e...)
	}
	if len(errs) > 0 {
		log.WithFields(log.Fields{
			"handler": "createPresentation",
			"errs":    strings.Join(errs, "\n"),
		})
		respondJSON(w, m.errorTransformer.ValidationErrorsToInvalidDataError(errs), "InvalidDataError", http.StatusBadRequest, errorTransformer)
		return
	}
	if result, err = m.handler.CreatePresentation(bodyPresentation); err != nil {
		switch err.(type) {
		case model.InvalidDataError:
			respondJSON(w, err, "InvalidDataError", 400, errorTransformer)
		default:
			respondJSON(w, m.errorTransformer.ErrorToString(err), "string", http.StatusInternalServerError, errorTransformer)
		}
		return
	}

	if errs = result.Validate(); len(errs) > 0 {
		err := errors.New("Invalid response data")
		log.WithFields(log.Fields{
			"dataType": "ReadOnlyPresentation",
			"error":    strings.Join(errs, "\n"),
		}).Error(err)
		respondJSON(w, m.errorTransformer.ErrorToString(err), "string", http.StatusInternalServerError, errorTransformer)
		return
	}

	respondJSON(w, result, "ReadOnlyPresentation", http.StatusOK, errorTransformer)
}

func (m *middleware) getPresentation(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	errorTransformer := func(err error) interface{} { return m.errorTransformer.ErrorToString(err) }

	defer func() {
		if recovered := recover(); recovered != nil {
			err := errors.New("Recovered")
			log.WithField("error", recovered).Error(err)
			respondJSON(w, m.errorTransformer.ErrorToString(err), "string", http.StatusInternalServerError, errorTransformer)
		}
	}()

	var (
		result model.ReadOnlyPresentation
		err    error
		errs   []string
	)

	pathPresentationID := params.ByName("presentationId")

	if result, err = m.handler.GetPresentation(pathPresentationID); err != nil {
		switch err.(type) {
		case model.NotFoundError:
			respondJSON(w, err, "NotFoundError", 404, errorTransformer)
		default:
			respondJSON(w, m.errorTransformer.ErrorToString(err), "string", http.StatusInternalServerError, errorTransformer)
		}
		return
	}

	if errs = result.Validate(); len(errs) > 0 {
		err := errors.New("Invalid response data")
		log.WithFields(log.Fields{
			"dataType": "ReadOnlyPresentation",
			"error":    strings.Join(errs, "\n"),
		}).Error(err)
		respondJSON(w, m.errorTransformer.ErrorToString(err), "string", http.StatusInternalServerError, errorTransformer)
		return
	}

	respondJSON(w, result, "ReadOnlyPresentation", http.StatusOK, errorTransformer)
}

func (m *middleware) updatePresentation(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	errorTransformer := func(err error) interface{} { return m.errorTransformer.ErrorToString(err) }

	defer func() {
		if recovered := recover(); recovered != nil {
			err := errors.New("Recovered")
			log.WithField("error", recovered).Error(err)
			respondJSON(w, m.errorTransformer.ErrorToString(err), "string", http.StatusInternalServerError, errorTransformer)
		}
	}()

	var (
		err  error
		errs []string
	)

	pathPresentationID := params.ByName("presentationId")

	var bodyUpdate model.Presentation
	if err = json.NewDecoder(r.Body).Decode(&bodyUpdate); err != nil {
		errs = append(errs, err.Error())
		log.WithFields(log.Fields{
			"bodyType": "Presentation",
			"error":    err,
		}).Error("Failed to parse body data")
	} else if e := bodyUpdate.Validate(); len(e) > 0 {
		errs = append(errs, e...)
	}
	if len(errs) > 0 {
		log.WithFields(log.Fields{
			"handler": "updatePresentation",
			"errs":    strings.Join(errs, "\n"),
		})
		respondJSON(w, m.errorTransformer.ValidationErrorsToInvalidDataError(errs), "InvalidDataError", http.StatusBadRequest, errorTransformer)
		return
	}
	if err = m.handler.UpdatePresentation(pathPresentationID, bodyUpdate); err != nil {
		switch err.(type) {
		case model.InvalidDataError:
			respondJSON(w, err, "InvalidDataError", 400, errorTransformer)
		case model.NotFoundError:
			respondJSON(w, err, "NotFoundError", 404, errorTransformer)
		default:
			respondJSON(w, m.errorTransformer.ErrorToString(err), "string", http.StatusInternalServerError, errorTransformer)
		}
		return
	}

	w.Write([]byte("OK"))
}

func (m *middleware) deletePresentation(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	errorTransformer := func(err error) interface{} { return m.errorTransformer.ErrorToString(err) }

	defer func() {
		if recovered := recover(); recovered != nil {
			err := errors.New("Recovered")
			log.WithField("error", recovered).Error(err)
			respondJSON(w, m.errorTransformer.ErrorToString(err), "string", http.StatusInternalServerError, errorTransformer)
		}
	}()

	var (
		err error
	)

	pathPresentationID := params.ByName("presentationId")

	if err = m.handler.DeletePresentation(pathPresentationID); err != nil {
		switch err.(type) {
		case model.NotFoundError:
			respondJSON(w, err, "NotFoundError", 404, errorTransformer)
		default:
			respondJSON(w, m.errorTransformer.ErrorToString(err), "string", http.StatusInternalServerError, errorTransformer)
		}
		return
	}

	w.Write([]byte("OK"))
}

func respondJSON(w http.ResponseWriter, data interface{}, dataType string, statusCode int, errorTransformer func(error) interface{}) {
	response, err := json.Marshal(data)
	if err != nil {
		log.WithFields(log.Fields{
			"dataType": dataType,
			"error":    err.Error(),
		}).Error("Failed to convert to json")

		// we need to assume here that converting the error does not lead to json marshalling errors
		// it is the responsibility of the implementer to not mess this up
		response, _ = json.Marshal(errorTransformer(err))
		statusCode = http.StatusInternalServerError
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func parseArray(s string) []string {
	// we treat the empty string as an empty array, rather than an array with one empty element
	if len(s) == 0 {
		return []string{}
	}
	return strings.Split(s, ",")
}

func validateString(s, name string, minLength, maxLength *int, enum []string) (errs []string) {
	if minLength != nil {
		if len(s) < *minLength {
			errs = append(errs, fmt.Sprintf("%s should be no shorter than %d characters", name, *minLength))
		}
	}

	if maxLength != nil {
		if len(s) > *maxLength {
			errs = append(errs, fmt.Sprintf("%s should be no longer than %d characters", name, *maxLength))
		}
	}

	if enum != nil {
		found := false
		for i := range enum {
			if s == enum[i] {
				found = true
				break
			}
		}
		if !found {
			errs = append(errs, fmt.Sprintf("%s is not an allowed value for %s", s, name))
		}
	}

	return
}

func validateArray(a []string, name string, minItems, maxItems *int, uniqueItems bool) (errs []string) {
	if minItems != nil {
		if len(a) < *minItems {
			errs = append(errs, fmt.Sprintf("%s should have no less than %d elements", name, *minItems))
		}
	}

	if maxItems != nil {
		if len(a) > *maxItems {
			errs = append(errs, fmt.Sprintf("%s should have no more than %d elements", name, *maxItems))
		}
	}

	if uniqueItems {
		seen := map[string]struct{}{}
		for _, elt := range a {
			if _, duplicate := seen[elt]; duplicate {
				errs = append(errs, fmt.Sprintf("%s occurs multiple times in %s", elt, name))
			}
			seen[elt] = struct{}{}
		}
	}

	return
}
