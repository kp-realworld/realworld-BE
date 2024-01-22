package auth

import (
	"encoding/json"
	"github.com/hotkimho/realworld-api/controller"
	"github.com/hotkimho/realworld-api/models"
	"github.com/hotkimho/realworld-api/repository"
	"github.com/hotkimho/realworld-api/responder"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var userReq CreateUserRequestDTO

	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		responder.Response(w, http.StatusBadRequest, err.Error())
		return
	}

	err = controller.Validate.Struct(userReq)
	if err != nil {
		responder.Response(w, http.StatusBadRequest, err.Error())
		return
	}

	user := models.User{
		Username: userReq.Username,
		Email:    userReq.Email,
		Password: userReq.Password,
		Name:     "wjoi",
	}

	_, err = repository.UserRepo.Create(repository.DB, user)
	if err != nil {
		responder.Response(w, http.StatusInternalServerError, err.Error())
		return
	}

	responder.Response(w, http.StatusOK, "success")
}
