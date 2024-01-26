package responder

import (
	"encoding/json"
	"github.com/hotkimho/realworld-api/controller/dto/user"
	"github.com/hotkimho/realworld-api/models"
	"net/http"
)

func ReadUserProfileResponse(w http.ResponseWriter, user models.User) {

	wrapper := userdto.ReadUserProfileResponseWrapperDTO{
		User: userdto.ReadUserProfileResponseDTO{
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
