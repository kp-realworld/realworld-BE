package responder

import (
	"encoding/json"
	"net/http"

	authdto "github.com/hotkimho/realworld-api/controller/dto/auth"
	"github.com/hotkimho/realworld-api/models"
)

func SignInResponse(w http.ResponseWriter, user models.User, accessToken, refreshToken string) {

	wrapper := authdto.SignInResponseWrapperDTO{
		User: authdto.SignInResponseDTO{
			UserID:       user.UserID,
			Username:     user.Username,
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
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

func VerifyUsernameResponse(w http.ResponseWriter, username string) {
	wrapper := authdto.VerifyUsernameResponseDTO{
		Username: username,
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

func VerifyEmailResponse(w http.ResponseWriter, email string) {
	wrapper := authdto.VerifyEmailResponseDTO{
		Email: email,
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
