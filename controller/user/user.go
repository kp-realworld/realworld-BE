package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hotkimho/realworld-api/controller/auth"

	"github.com/hotkimho/realworld-api/controller/dto/user"
	"github.com/hotkimho/realworld-api/repository"
	"github.com/hotkimho/realworld-api/responder"
)

// @Summary 내 프로필 조회
// @Description 내 프로필 조회
// @Tags Profile tag
// @Accept json
// @Produce json
// @Param authorization header string true "jwt token"
// @Success 200 {object} userdto.ReadMyProfileResponseDTO "success"
// @Failure 404 {object} types.ErrorResponse "user 정보가 없는 경우"
// @Failure 500 {object} types.ErrorResponse "network error"
// @Router /my/profile [get]
//func ReadMyProfile(w http.ResponseWriter, r *http.Request) {
//	ctxUserID := r.Context().Value("ctx_user_id").(int64)
//
//	user, err := repository.UserRepo.GetByID(repository.DB, ctxUserID)
//	if err != nil {
//		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
//		return
//	} else if user == nil {
//		responder.ErrorResponse(w, http.StatusNotFound, "user not found")
//		return
//	}
//
//	responder.ReadMyProfileResponse(w, *user)
//}

// header에 authorization이 있어야 한다.
// @Summary Read user profile
// @Description Read user profile
// @Tags Profile tag
// @Accept json
// @Produce json
// @Param username query string true "username"
// @Param authorization header string false "jwt token"
// @Success 200 {object} userdto.ReadUserProfileResponseWrapperDTO "success"
// @Failure 400 {object} types.ErrorResponse "bad request"
// @Failure 404 {object} types.ErrorResponse "user not found"
// @Failure 500 {object} types.ErrorResponse "network error"
// @Router /profile [get]
func ReadUserProfile(w http.ResponseWriter, r *http.Request) {
	ctxUserID := r.Context().Value("ctx_user_id").(int64)

	username := r.URL.Query().Get("username")
	if username == "" {
		responder.ErrorResponse(w, http.StatusBadRequest, "invalid username")
		return
	}

	fmt.Println("username: ", username)
	user, err := repository.UserRepo.GetByUsername(repository.DB, username)
	if err != nil {
		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	} else if user == nil {
		responder.ErrorResponse(w, http.StatusNotFound, "user not found")
		return
	}

	var isFollowed bool
	if ctxUserID > 0 {
		followed, err := repository.FollowRepo.IsFollowing(repository.DB, ctxUserID, user.UserID)
		if err != nil {
			responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		isFollowed = followed
	}

	responder.ReadUserProfileResponse(w, *user, isFollowed)
}

// @Summary Update user profile
// @Description Update user profile
// @Tags Profile tag
// @Accept json
// @Produce json
// @Param username query string true "username"
// @Param authorization header string false "jwt token"
// @Param updateUserProfileReq body userdto.UpdateUserProfileRequestDTO true "updateUserProfileReq"
// @Success 200 {object} userdto.UpdateUserProfileResponseWrapperDTO "success"
// @Failure 400 {object} types.ErrorResponse "bad request"
// @Failure 422 {object} types.ErrorResponse "요청을 제대로 수행하지 못함"
// @Failure 500 {object} types.ErrorResponse "network error"
// @Router /profile [put]
func UpdateUserProfile(w http.ResponseWriter, r *http.Request) {
	ctxUserID := r.Context().Value("ctx_user_id").(int64)

	var updateUserProfileReq userdto.UpdateUserProfileRequestDTO

	err := json.NewDecoder(r.Body).Decode(&updateUserProfileReq)
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if updateUserProfileReq.IsEmpty() {
		responder.ErrorResponse(w, http.StatusBadRequest, "invalid request")
		return
	}

	if updateUserProfileReq.Username == nil || *updateUserProfileReq.Username == "" {
		responder.ErrorResponse(w, http.StatusBadRequest, "invalid username")
		return
	}

	if err := ValidateUpdateUserProfileRequestDTO(updateUserProfileReq); err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var password *string
	if updateUserProfileReq.Password != nil {
		hashedPass, err := auth.HashPassword(*updateUserProfileReq.Password)
		if err != nil {
			responder.ErrorResponse(w, http.StatusUnprocessableEntity, err.Error())
			return
		}
		password = &hashedPass
	}

	user, err := repository.UserRepo.UpdateUserProfileByUsernameAndUserID(repository.DB, updateUserProfileReq, *updateUserProfileReq.Username, password, ctxUserID)
	if err != nil {
		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	responder.UpdateUserProfileResponse(w, *user)
}
