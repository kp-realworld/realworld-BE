package responder

import (
	"encoding/json"
	"github.com/hotkimho/realworld-api/types"
	"github.com/sirupsen/logrus"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errorResponse := BuildErrorResponse(statusCode, message)

	err := json.NewEncoder(w).Encode(errorResponse)
	if err != nil {
		logrus.Error(err)
		// report sentry
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
