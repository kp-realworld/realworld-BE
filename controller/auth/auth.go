package auth

import (
	"encoding/json"
	authdto "github.com/hotkimho/realworld-api/controller/dto/auth"
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
// @Param signUpReq body authdto.SignUpRequestDTO true "signUpReq"
// @Success 201 {object} authdto.SignUpResponseWrapperDTO "success"
// @Failure 400 {object} types.ErrorResponse "입력값이 유효하지 않음"
// @Failure 422 {object} types.ErrorResponse "이미 존재하는 계정"
// @Router /user/signup [post]
func SignUp(w http.ResponseWriter, r *http.Request) {

	var signUpReq authdto.SignUpRequestDTO

	err := json.NewDecoder(r.Body).Decode(&signUpReq)
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := ValidateSignUpRequestDTO(signUpReq); err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if !repository.UserRepo.CheckEmailOrUsername(repository.DB, signUpReq.Email, signUpReq.Username) {
		responder.ErrorResponse(w, http.StatusUnprocessableEntity, "already username or email")
		return
	}

	hashedPassword, err := HashPassword(signUpReq.Password)
	if err != nil {
		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	createdUser, err := repository.UserRepo.Create(repository.DB, SignUpDTOToUser(signUpReq, hashedPassword))
	if err != nil {
		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	responder.SignUpResponse(w, signUpReq, createdUser.UserID)
}

// @Summary 로그인
// @Description 로그인
// @Tags Auth tag
// @Accept json
// @Produce json
// @Param signInReq body authdto.SignInRequestDTO true "signInReq"
// @Success 200 {object} authdto.SignInResponseWrapperDTO "로그인 성공"
// @Failure 400 {object} types.ErrorResponse "입력값이 유효하지 않음"
// @Failure 422 {object} types.ErrorResponse "유저가 존재하지 않거나 비밀번호가 틀림"
// @Failure 500 {object} types.ErrorResponse "network error"
// @Router /user/signin [post]
func SignIn(w http.ResponseWriter, r *http.Request) {

	var SignInReq authdto.SignInRequestDTO

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
		responder.ErrorResponse(w, http.StatusUnprocessableEntity, "user not found")
		return
	}

	if !CheckPasswordHash(SignInReq.Password, user.Password) {
		responder.ErrorResponse(w, http.StatusUnprocessableEntity, "password incorrect")
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
