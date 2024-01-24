package auth

import (
	"encoding/json"
	"fmt"
	uuid2 "github.com/google/uuid"
	"github.com/hotkimho/realworld-api/controller"
	"github.com/hotkimho/realworld-api/models"
	"github.com/hotkimho/realworld-api/repository"
	"github.com/hotkimho/realworld-api/responder"
	"github.com/hotkimho/realworld-api/util"
	"net/http"
)

// @Summary 회원가입 sum
// @Description 회원가입 des
// @Tags Auth tag
// @Accept json
// @Produce json
// @Param signUpReq body SignUpRequestDTO true "signUpReq"
// @Success 200 {string} string "success"
// @Failure 400 {string} string "bad request"
// @Router /user/signup [post]
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
		ID:       uuid2.New().String(),
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

	fmt.Println("????????")
	token, err := util.IssueJWT(user.ID)
	if err != nil {
		responder.Response(w, http.StatusInternalServerError, err.Error())
		return
	}

	signInRes := SignInResponseDTO{
		Email:    user.Email,
		Username: user.Username,
		Token:    token,
	}

	resJson, _ := json.Marshal(signInRes)
	responder.Response(w, http.StatusOK, string(resJson))
}

func Heartbeat(w http.ResponseWriter, r *http.Request) {

	responder.Response(w, http.StatusOK, "success")
}
