package follow

import (
	"github.com/hotkimho/realworld-api/repository"
	"github.com/hotkimho/realworld-api/responder"
	"github.com/hotkimho/realworld-api/util"
	"net/http"
)

// @Summary Follow user
// @Description Follow user
// @Tags Follow tag
// @Accept json
// @Produce json
// @Param authorization header string true "jwt token"
// @Param followed_id path int true "팔로우할 user id"
// @Success 200 {object} followdto.CreateFollowResponseWrapperDTO "이미 팔로우가 되어 있는 경우"
// @Success 201 {object} followdto.CreateFollowResponseWrapperDTO "팔로우한 유저 정보"
// @Failure 400 {object} types.ErrorResponse "followed id가 유효하지 않음"
// @Failure 404 {object} types.ErrorResponse "팔로울 할 유저가 존재하지 않음"
// @Failure 500 {object} types.ErrorResponse "network error"
// @Router /user/follow/{followed_id} [post]
func CreateFollow(w http.ResponseWriter, r *http.Request) {
	ctxUserID := r.Context().Value("ctx_user_id").(int64)

	followedUserID, err := util.GetIntegerParam[int64](r, "followed_id")
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// 팔로우할 유저가 존재하는지 확인
	user, err := repository.UserRepo.GetByID(repository.DB, followedUserID)
	if err != nil {
		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	} else if user == nil {
		responder.ErrorResponse(w, http.StatusNotFound, "followed user not found")
		return
	}

	// 이미 팔로우가 되어 있는지 확인
	isFollowed, err := repository.FollowRepo.IsFollowing(repository.DB, ctxUserID, followedUserID)
	if err != nil {
		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	} else if isFollowed { // 이미 팔로우가 되어 있는 경우
		responder.CreateFollowResponse(w, http.StatusOK, *user)
		return
	}

	err = repository.FollowRepo.Create(repository.DB, ctxUserID, followedUserID)
	if err != nil {
		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	responder.CreateFollowResponse(w, http.StatusCreated, *user)
}

// @Summary Unfollow user
// @Description Unfollow user
// @Tags Follow tag
// @Accept json
// @Produce json
// @Param authorization header string true "jwt token"
// @Param followed_id path int true "언팔로우할 user id"
// @Success 204 {object} followdto.DeleteFollowResponseWrapperDTO "언팔로우한 유저 정보"
// @Failure 400 {object} types.ErrorResponse "followed id가 유효하지 않음"
// @Failure 404 {object} types.ErrorResponse "언팔로우할 유저가 존재하지 않음"
// @Failure 500 {object} types.ErrorResponse "network error"
// @Router /user/follow/{followed_id} [delete]
func DeleteFollow(w http.ResponseWriter, r *http.Request) {
	ctxUserID := r.Context().Value("ctx_user_id").(int64)

	followedUserID, err := util.GetIntegerParam[int64](r, "followed_id")
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// 언팔로우할 유저가 존재하는지 확인
	user, err := repository.UserRepo.GetByID(repository.DB, followedUserID)
	if err != nil {
		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	} else if user == nil {
		responder.ErrorResponse(w, http.StatusNotFound, "followed user not found")
		return
	}

	err = repository.FollowRepo.Delete(repository.DB, ctxUserID, followedUserID)
	if err != nil {
		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	responder.DeleteFollowResponse(w, *user)
}
