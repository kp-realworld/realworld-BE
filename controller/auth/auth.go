package auth

import (
	"encoding/json"
	"fmt"
	"github.com/hotkimho/realworld-api/controller"
	"github.com/hotkimho/realworld-api/models"
	"github.com/hotkimho/realworld-api/repository"
	"github.com/hotkimho/realworld-api/responder"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {

	var signUpReq SignUpRequestDTO

	err := json.NewDecoder(r.Body).Decode(&signUpReq)
	if err != nil {
		responder.Response(w, http.StatusBadRequest, err.Error())
		return
	}

	err = controller.Validate.Struct(signUpReq)
	if err != nil {
		responder.Response(w, http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, err := HashPassword(signUpReq.Password)
	if err != nil {
		responder.Response(w, http.StatusInternalServerError, err.Error())
		return
	}

	user := models.User{
		Username: signUpReq.Username,
		Email:    signUpReq.Email,
		Password: hashedPassword,
	}

	_, err = repository.UserRepo.Create(repository.DB, user)
	if err != nil {
		responder.Response(w, http.StatusInternalServerError, err.Error())
		return
	}

	responder.Response(w, http.StatusOK, "success")
}

func SignIn(w http.ResponseWriter, r *http.Request) {

	var SignInReq SignInRequestDTO

	err := json.NewDecoder(r.Body).Decode(&SignInReq)
	if err != nil {
		responder.Response(w, http.StatusBadRequest, err.Error())
		return
	}

	err = controller.Validate.Struct(SignInReq)
	if err != nil {
		fmt.Println("???")
		responder.Response(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := repository.UserRepo.GetByEmail(repository.DB, SignInReq.Email)
	if err != nil {
		responder.Response(w, http.StatusInternalServerError, err.Error())
		return
	} else if user == nil {
		responder.Response(w, http.StatusNotFound, "user not found")
		return
	}

	if !CheckPasswordHash(SignInReq.Password, user.Password) {
		responder.Response(w, http.StatusUnauthorized, "password incorrect")
		return
	}

	responder.Response(w, http.StatusOK, "success")
}

func Heartbeat(w http.ResponseWriter, r *http.Request) {

	responder.Response(w, http.StatusOK, "success")
}
