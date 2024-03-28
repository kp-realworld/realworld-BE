package responder

import (
	"encoding/json"
	"github.com/hotkimho/realworld-api/controller/dto/user"
	"github.com/hotkimho/realworld-api/models"
	"net/http"
)

func ReadMyProfileResponse(w http.ResponseWriter, user models.User) {

	wrapper := userdto.ReadMyProfileResponseWrapperDTO{
		User: userdto.ReadMyProfileResponseDTO{
			Username:     user.Username,
			Bio:          user.Bio,
			ProfileImage: user.ProfileImage,
		},
	}

	jsonData, err := json.Marshal(wrapper)
	if err != nil {
		ErrorResponse(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func ReadUserProfileResponse(w http.ResponseWriter, user models.User, isFollowing bool) {

	wrapper := userdto.ReadUserProfileResponseWrapperDTO{
		User: userdto.ReadUserProfileResponseDTO{
			UserID:       user.UserID,
			Username:     user.Username,
			Bio:          user.Bio,
			ProfileImage: user.ProfileImage,
			Following:    &isFollowing,
			Email:        user.Email,
		},
	}

	jsonData, err := json.Marshal(wrapper)
	if err != nil {
		ErrorResponse(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func UpdateUserProfileResponse(w http.ResponseWriter, user models.User) {

	wrapper := userdto.UpdateUserProfileResponseWrapperDTO{
		User: userdto.UpdateUserProfileResponseDTO{
			Username:     user.Username,
			Bio:          user.Bio,
			ProfileImage: user.ProfileImage,
		},
	}

	jsonData, err := json.Marshal(wrapper)
	if err != nil {
		ErrorResponse(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
