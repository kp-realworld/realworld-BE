package responder

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"

	"github.com/hotkimho/realworld-api/types"
)

func ErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errorResponse := BuildErrorResponse(statusCode, message)

	if statusCode >= 500 {
		logrus.Error(message)
		sentry.CaptureException(errors.New(message))
	}
	err := json.NewEncoder(w).Encode(errorResponse)
	if err != nil {
		logrus.Error(err)
		sentry.CaptureException(err)
	}
}

func BuildErrorResponse(statusCode int, message string) types.ErrorResponse {
	return types.ErrorResponse{
		Errors: types.ErrorDetail{
			Message: message,
			Code:    statusCode,
		},
	}
}
