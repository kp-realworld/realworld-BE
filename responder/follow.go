package responder

import (
	"encoding/json"
	followdto "github.com/hotkimho/realworld-api/controller/dto/follow"
	"github.com/hotkimho/realworld-api/models"
	"net/http"
)

func CreateFollowResponse(w http.ResponseWriter, code int, user models.User) {

	wrapper := followdto.CreateFollowResponseWrapperDTO{
		Profile: followdto.CreateFollowResponseDTO{
			Username:     user.Username,
			Bio:          user.Bio,
			ProfileImage: user.ProfileImage,
			Following:    true,
		},
	}

	jsonData, err := json.Marshal(wrapper)
	if err != nil {
		ErrorResponse(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(jsonData)
}

func DeleteFollowResponse(w http.ResponseWriter, user models.User) {

	wrapper := followdto.DeleteFollowResponseWrapperDTO{
		Profile: followdto.DeleteFollowResponseDTO{
			Username:     user.Username,
			Bio:          user.Bio,
			ProfileImage: user.ProfileImage,
			Following:    false,
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
