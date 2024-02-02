package responder

import (
	"encoding/json"
	authdto "github.com/hotkimho/realworld-api/controller/dto/auth"
	"github.com/hotkimho/realworld-api/models"
	"net/http"
)

func SignInResponse(w http.ResponseWriter, user models.User, token string) {

	wrapper := authdto.SignInResponseWrapperDTO{
		User: authdto.SignInResponseDTO{
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

func SignUpResponse(w http.ResponseWriter, requestDTO authdto.SignUpRequestDTO, userID int64) {

	wrapper := authdto.SignUpResponseWrapperDTO{
		User: authdto.SignUpResponseDTO{
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
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonData)
}
func RefreshTokenResponse(w http.ResponseWriter, token string) {
	wrapper := authdto.RefreshTokenResponseDTO{
		Token: token,
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
