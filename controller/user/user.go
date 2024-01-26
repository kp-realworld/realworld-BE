package user

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/hotkimho/realworld-api/controller/dto/user"
	"github.com/hotkimho/realworld-api/repository"
	"github.com/hotkimho/realworld-api/responder"
	"net/http"
	"strconv"
)

// header에 authorization이 있어야 한다.
// @Summary Read user profile
// @Description Read user profile
// @Tags User tag
// @Accept json
// @Produce json
// @Param user_id path int true "user_id"
// @Param authorization header string true "jwt token"
// @Success 200 {object} userdto.ReadUserProfileResponseWrapperDTO "success"
// @Failure 400 {object} types.ErrorResponse "bad request"
// @Failure 404 {object} types.ErrorResponse "user not found"
// @Failure 422 {object} types.ErrorResponse "요청을 제대로 수행하지 못함"
// @Failure 500 {object} types.ErrorResponse "network error"
// @Router /user/{user_id}/profile [get]
func ReadUserProfile(w http.ResponseWriter, r *http.Request) {
	// @Header authorization string true "jwt token"
	vars := mux.Vars(r)
	userID, err := strconv.ParseInt(vars["user_id"], 10, 64)
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if userID <= 0 {
		responder.ErrorResponse(w, http.StatusBadRequest, "invalid user_id")
	}

	user, err := repository.UserRepo.GetByID(repository.DB, userID)
	if err != nil {
		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	} else if user == nil {
		responder.ErrorResponse(w, http.StatusNotFound, "user not found")
		return
	}

	responder.ReadUserProfileResponse(w, *user)
}

// @Summary Update user profile
// @Description Update user profile
// @Tags User tag
// @Accept json
// @Produce json
// @Param user_id path int true "user_id"
// @Param authorization header string true "jwt token"
// @Param updateUserProfileReq body userdto.UpdateUserProfileRequestDTO true "updateUserProfileReq"
// @Success 200 {object} userdto.UpdateUserProfileResponseWrapperDTO "success"
// @Failure 400 {object} types.ErrorResponse "bad request"
// @Failure 422 {object} types.ErrorResponse "요청을 제대로 수행하지 못함"
// @Failure 500 {object} types.ErrorResponse "network error"
// @Router /user/{user_id}/profile [put]
func UpdateUserProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.ParseInt(vars["user_id"], 10, 64)
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	var updateUserProfileReq userdto.UpdateUserProfileRequestDTO

	err = json.NewDecoder(r.Body).Decode(&updateUserProfileReq)
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if updateUserProfileReq.IsEmpty() {
		responder.ErrorResponse(w, http.StatusBadRequest, "invalid request")
		return
	}

	if err := ValidateUpdateUserProfileRequestDTO(updateUserProfileReq); err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := repository.UserRepo.UpdateUserProfileByID(repository.DB, updateUserProfileReq, userID)
	if err != nil {
		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	responder.UpdateUserProfileResponse(w, *user)
}
