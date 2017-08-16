package presentation

import (
	"net/http"

	"github.com/vincentvanderweele/wwweeklies-presentations/generated/router"
)

type presentationNotFoundError string

func (e presentationNotFoundError) Error() string {
	return string(e)
}

type errorTransformer struct{}

// NewErrorTransformer returns a new error transformer
func NewErrorTransformer() router.ErrorTransformer {
	return &errorTransformer{}
}

func (h *errorTransformer) Transform(err error) (message string, code int) {
	switch err.(type) {
	case presentationNotFoundError:
		message = "Presentation not found"
		code = http.StatusNotFound
	default:
		message = "Internal server error"
		code = http.StatusInternalServerError
	}

	return
}
