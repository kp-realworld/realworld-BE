package auth

import (
	"encoding/json"
	. "github.com/hotkimho/realworld-api/controller/dto/auth"
	"github.com/hotkimho/realworld-api/repository"
	"github.com/hotkimho/realworld-api/responder"
	"github.com/hotkimho/realworld-api/util"
	"net/http"
)

// @Summary 회원가입
// @Description 회원가입
// @Tags Auth tag
// @Accept json
// @Produce json
// @Param signUpReq body SignUpRequestDTO true "signUpReq"
// @Success 201 {object} SignUpResponseWrapperDTO "success"
// @Failure 400 {object} types.ErrorResponse "bad request"
// @Router /user/signup [post]
func SignUp(w http.ResponseWriter, r *http.Request) {

	var signUpReq SignUpRequestDTO

	err := json.NewDecoder(r.Body).Decode(&signUpReq)
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := ValidateSignUpRequestDTO(signUpReq); err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, err := HashPassword(signUpReq.Password)
	if err != nil {
		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	userID, err := repository.UserRepo.Create(repository.DB, SignUpDTOToUser(signUpReq, hashedPassword))
	if err != nil {
		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	responder.SignUpResponse(w, signUpReq, userID)
}

// @Summary 로그인
// @Description 로그인
// @Tags Auth tag
// @Accept json
// @Produce json
// @Param signInReq body SignInRequestDTO true "signInReq"
// @Success 200 {object} SignInResponseWrapperDTO "success"
// @Failure 400 {object} types.ErrorResponse "bad request"
// @Router /user/signin [post]
func SignIn(w http.ResponseWriter, r *http.Request) {

	var SignInReq SignInRequestDTO

	err := json.NewDecoder(r.Body).Decode(&SignInReq)
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := ValidateSignInRequestDTO(SignInReq); err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := repository.UserRepo.GetByEmail(repository.DB, SignInReq.Email)
	if err != nil {
		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	} else if user == nil {
		responder.ErrorResponse(w, http.StatusBadRequest, "user not found")
		return
	}

	if !CheckPasswordHash(SignInReq.Password, user.Password) {
		responder.ErrorResponse(w, http.StatusBadRequest, "password incorrect")
		return
	}

	token, err := util.IssueJWT(user.UserID)
	if err != nil {
		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	responder.SignInResponse(w, *user, token)
}

func Heartbeat(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
