package comment

import (
	"encoding/json"
	commentdto "github.com/hotkimho/realworld-api/controller/dto/comment"
	"github.com/hotkimho/realworld-api/repository"
	"github.com/hotkimho/realworld-api/responder"
	"github.com/hotkimho/realworld-api/util"
	"net/http"
)

// @Summary 댓글 생성
// @Description 댓글을 생성합니다.
// @Tags Comment tag
// @Accept json
// @Produce json
// @Param authorization header string true "jwt token"
// @Param user_id path int true "article author id(기사 소유자)"
// @Param article_id path int true "article id(기사 ID)"
// @Param createCommentReq body commentdto.CreateCommentRequestDTO true "댓글 내용"
// @Success 201 {object} commentdto.CreateCommentResponseDTO
// @Failure 400 {object} types.ErrorResponse "잘못된 값을 요청한 경우"
// @Failure 422 {object} types.ErrorResponse "요청을 처리하지 못한 경우"
// @Failure 500 {object} types.ErrorResponse "네트워크 에러"
// @Router /user/{user_id}/article/{article_id}/comment [post]
func CreateComment(w http.ResponseWriter, r *http.Request) {

	ctxUserID := r.Context().Value("ctx_user_id").(int64)
	authorID, err := util.GetIntegerParam[int64](r, "user_id")
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	articleID, err := util.GetIntegerParam[int64](r, "article_id")
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var createCommentReq commentdto.CreateCommentRequestDTO

	err = json.NewDecoder(r.Body).Decode(&createCommentReq)
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := ValidateCreateCommentRequestDTO(createCommentReq); err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// article author가 맞는지 확인
	if err := repository.ArticleRepo.ValidateArticleOwner(repository.DB, articleID, authorID); err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, "invalid article author")
		return
	}

	comment, err := repository.CommentRepo.Create(repository.DB, createCommentReq, ctxUserID, articleID)
	if err != nil {
		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// response
	responder.CreateCommentResponse(w, *comment)
}

// @Summary 댓글 목록 조회
// @Description 기사의 댓글 목록을 조회합니다.
// @Tags Comment tag
// @Accept json
// @Produce json
// @Param user_id path int true "article author id(기사 소유자)"
// @Param article_id path int true "article id(기사 ID)"
// @Success 200 {object} commentdto.ReadCommentsResponseWrapperDTO
// @Failure 400 {object} types.ErrorResponse "잘못된 값을 요청한 경우"
// @Failure 500 {object} types.ErrorResponse "네트워크 에러"
// @Router /user/{user_id}/article/{article_id}/comments [get]
func ReadComments(w http.ResponseWriter, r *http.Request) {
	authorID, err := util.GetIntegerParam[int64](r, "user_id")
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	articleID, err := util.GetIntegerParam[int64](r, "article_id")
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// article author가 맞는지 확인
	if err := repository.ArticleRepo.ValidateArticleOwner(repository.DB, articleID, authorID); err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, "invalid article author")
		return
	}

	comments, err := repository.CommentRepo.GetByArticle(repository.DB, articleID)
	if err != nil {
		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// response
	responder.ReadCommentsResponse(w, comments)
}

// @Summary 댓글 수정
// @Description 댓글을 수정합니다.
// @Tags Comment tag
// @Accept json
// @Produce json
// @Param authorization header string true "jwt token"
// @Param user_id path int true "article author id(기사 소유자)"
// @Param article_id path int true "article id(기사 ID)"
// @Param comment_id path int true "comment id(댓글 ID)"
// @Param updateCommentReq body commentdto.UpdateCommentRequestDTO true "updateCommentReq"
// @Success 200 {object} commentdto.UpdateCommentResponseDTO
// @Failure 400 {object} types.ErrorResponse "잘못된 값을 요청한 경우"
// @Failure 422 {object} types.ErrorResponse "요청을 처리하지 못한 경우"
// @Failure 500 {object} types.ErrorResponse "네트워크 에러"
// @Router /user/{user_id}/article/{article_id}/comment/{comment_id} [put]
func UpdateComment(w http.ResponseWriter, r *http.Request) {
	ctxUserID := r.Context().Value("ctx_user_id").(int64)
	authorID, err := util.GetIntegerParam[int64](r, "user_id")
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	articleID, err := util.GetIntegerParam[int64](r, "article_id")
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	commentID, err := util.GetIntegerParam[int64](r, "comment_id")
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var updateCommentReq commentdto.UpdateCommentRequestDTO

	err = json.NewDecoder(r.Body).Decode(&updateCommentReq)
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := ValidateUpdateCommentRequestDTO(updateCommentReq); err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// article author가 맞는지 확인
	if err := repository.ArticleRepo.ValidateArticleOwner(repository.DB, articleID, authorID); err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, "invalid article author")
		return
	}

	comment, err := repository.CommentRepo.UpdateByID(repository.DB, updateCommentReq, commentID, ctxUserID, articleID)
	if err != nil {
		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	} else if comment == nil {
		responder.ErrorResponse(w, http.StatusBadRequest, "invalid comment author")
		return
	}

	// response
	responder.UpdateCommentResponse(w, *comment)
}

// @Summary 댓글 삭제
// @Description 댓글을 삭제합니다.
// @Tags Comment tag
// @Accept json
// @Produce json
// @Param authorization header string true "jwt token"
// @Param user_id path int true "article author id(기사 소유자)"
// @Param article_id path int true "article id(기사 ID)"
// @Param comment_id path int true "comment id(댓글 ID)"
// @Success 204 {object} string "정상적으로 삭제됨"
// @Failure 400 {object} types.ErrorResponse "잘못된 값을 요청한 경우"
// @Failure 500 {object} types.ErrorResponse "네트워크 에러"
// @Router /user/{user_id}/article/{article_id}/comment/{comment_id} [delete]
func DeleteComment(w http.ResponseWriter, r *http.Request) {
	ctxUserID := r.Context().Value("ctx_user_id").(int64)
	authorID, err := util.GetIntegerParam[int64](r, "user_id")
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	articleID, err := util.GetIntegerParam[int64](r, "article_id")
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	commentID, err := util.GetIntegerParam[int64](r, "comment_id")
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// article author가 맞는지 확인
	if err := repository.ArticleRepo.ValidateArticleOwner(repository.DB, articleID, authorID); err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, "invalid article author")
		return
	}

	err = repository.CommentRepo.DeleteByID(repository.DB, commentID, ctxUserID, articleID)
	if err != nil {
		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// response
	w.WriteHeader(http.StatusNoContent)
}
