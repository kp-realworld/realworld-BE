package responder

import (
	"encoding/json"
	"github.com/hotkimho/realworld-api/controller/dto/auth"
	"github.com/hotkimho/realworld-api/models"
	"net/http"
)

func SignInResponse(w http.ResponseWriter, user models.User, token string) {

	wrapper := auth.SignInResponseWrapperDTO{
		User: auth.SignInResponseDTO{
			UserID:   user.UserID,
			Username: user.Username,
			Token:    token,
		},
	}

	jsonData, err := json.Marshal(wrapper)
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func SignUpResponse(w http.ResponseWriter, requestDTO auth.SignUpRequestDTO, userID int64) {

	wrapper := auth.SignUpResponseWrapperDTO{
		User: auth.SignUpResponseDTO{
			Username: requestDTO.Username,
			Email:    requestDTO.Email,
			UserID:   userID,
		},
	}

	jsonData, err := json.Marshal(wrapper)
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
