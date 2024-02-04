package auth

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"

	authdto "github.com/hotkimho/realworld-api/controller/dto/auth"
	"github.com/hotkimho/realworld-api/env"
	"github.com/hotkimho/realworld-api/repository"
	"github.com/hotkimho/realworld-api/responder"
	"github.com/hotkimho/realworld-api/types"
	"github.com/hotkimho/realworld-api/util"
)

func Heartbeat(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

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

// @Summary 토큰 갱신
// @Description 토큰 갱신
// @Tags Auth tag
// @Accept json
// @Produce json
// @Param authorization header string true "jwt token"
// @Success 200 {object} authdto.RefreshTokenResponseDTO "success"
// @Failure 400 {object} types.ErrorResponse "입력값이 유효하지 않음"
// @Failure 500 {object} types.ErrorResponse "network error"
// @Router /user/{user_id}/refresh [get]
func RefreshToken(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("Authorization")
	if token == "" {
		responder.ErrorResponse(w, http.StatusBadRequest, "token is empty")
		return
	}

	userID, err := util.GetIntegerParam[int64](r, "user_id")
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	parsedToken, err := jwt.ParseWithClaims(token, &types.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(env.Config.Auth.Secret), nil
	})
	if err != nil {
		if strings.Compare(err.Error(), types.ERR_EXPIRED_TOKEN) != 0 {
			responder.ErrorResponse(w, http.StatusUnauthorized, err.Error())
			return
		}
	}

	if claim, ok := parsedToken.Claims.(*types.JWTClaims); ok && parsedToken.Valid {
		if claim.UserID != userID {
			responder.ErrorResponse(w, http.StatusBadRequest, "invalid token and user_id")
			return
		}
	}

	jwtToken, err := util.IssueJWT(userID)
	if err != nil {
		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	responder.RefreshTokenResponse(w, jwtToken)
}

// @Summary 유저네임 중복 확인
// @Description 유저네임 중복 확인
// @Tags Auth tag
// @Accept json
// @Produce json
// @Param verifyUsernameReq body authdto.VerifyUsernameRequestDTO true "verifyUsernameReq"
// @Success 200 {object} authdto.VerifyUsernameResponseDTO "success"
// @Failure 400 {object} types.ErrorResponse "입력값이 유효하지 않음"
// @Failure 409 {object} types.ErrorResponse "이미 존재하는 username"
// @Failure 500 {object} types.ErrorResponse "network error"
// @Router /user/verify-username [post]
func VerifyUsername(w http.ResponseWriter, r *http.Request) {

	var verifyReq authdto.VerifyUsernameRequestDTO

	err := json.NewDecoder(r.Body).Decode(&verifyReq)
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := ValidateVerifyUsernameRequestDTO(verifyReq); err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := repository.UserRepo.GetByUsername(repository.DB, verifyReq.Username)
	if err != nil {
		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	} else if user != nil {
		responder.ErrorResponse(w, http.StatusConflict, "already username")
	}

	responder.VerifyUsernameResponse(w, verifyReq.Username)
}

// @Summary 이메일 중복 확인
// @Description 이메일 중복 확인
// @Tags Auth tag
// @Accept json
// @Produce json
// @Param verifyEmailReq body authdto.VerifyEmailRequestDTO true "verifyEmailReq"
// @Success 200 {object} authdto.VerifyEmailResponseDTO "success"
// @Failure 400 {object} types.ErrorResponse "입력값이 유효하지 않거나 이메일 형식이 아님"
// @Failure 409 {object} types.ErrorResponse "이미 존재하는 email"
// @Failure 500 {object} types.ErrorResponse "network error"
// @Router /user/verify-email [post]
func VerifyEmail(w http.ResponseWriter, r *http.Request) {

	var verifyReq authdto.VerifyEmailRequestDTO

	err := json.NewDecoder(r.Body).Decode(&verifyReq)
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := ValidateVerifyEmailRequestDTO(verifyReq); err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if !util.VerifyEmail(verifyReq.Email) {
		responder.ErrorResponse(w, http.StatusBadRequest, "invalid email")
		return
	}

	user, err := repository.UserRepo.GetByEmail(repository.DB, verifyReq.Email)
	if err != nil {
		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	} else if user != nil {
		responder.ErrorResponse(w, http.StatusConflict, "already email")
	}

	responder.VerifyEmailResponse(w, verifyReq.Email)
}
