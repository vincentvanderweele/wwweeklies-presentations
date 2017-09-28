package presentation

import (
	"strings"

	"github.com/vincentvanderweele/wwweeklies-presentations/generated/model"
	"github.com/vincentvanderweele/wwweeklies-presentations/generated/router"
)

type errorTransformer struct{}

// NewErrorTransformer returns a new error transformer
func NewErrorTransformer() router.ErrorTransformer {
	return &errorTransformer{}
}

func (t *errorTransformer) ValidationErrorsToInvalidDataError(errs []string) model.InvalidDataError {
	return model.InvalidDataError(strings.Join(errs, "\n"))
}

func (t *errorTransformer) ErrorToString(err error) string {
	return err.Error()
}
