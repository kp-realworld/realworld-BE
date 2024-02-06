package article

import (
	"encoding/json"
	"net/http"

	"github.com/hotkimho/realworld-api/controller/dto/article"
	"github.com/hotkimho/realworld-api/repository"
	"github.com/hotkimho/realworld-api/responder"
	"github.com/hotkimho/realworld-api/util"
)

// @Summary 기사 생성
// @Description 기사를 생성합니다.
// @Tags Article tag
// @Accept json
// @Produce json
// @Param authorization header string true "jwt token"
// @Param createArticleReq body articledto.CreateArticleRequestDTO true "기사 생성 내용"
// @Success 201 {object} articledto.CreateArticleResponseWrapperDTO "정상적으로 생성됨"
// @Failure 400 {object} types.ErrorResponse "입력값이 유효하지 않음"
// @Failure 422 {object} types.ErrorResponse "요청을 제대로 수행하지 못함"
// @Failure 500 {object} types.ErrorResponse "네트워크 에러"
// @Router /user/{user_id}/article [post]
func CreateArticle(w http.ResponseWriter, r *http.Request) {

	userID, err := util.GetIntegerParam[int64](r, "user_id")
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var createArticleReq articledto.CreateArticleRequestDTO

	err = json.NewDecoder(r.Body).Decode(&createArticleReq)
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := ValidateCreateArticleRequestDTO(createArticleReq); err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	article, err := repository.ArticleRepo.CreateWithTransaction(
		repository.DB,
		createArticleReq,
		userID)
	if err != nil {
		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	responder.CreateArticleResponse(w, *article, createArticleReq.TagList)
}

// @Summary Read article by id
// @Description Read article by id
// @Tags Article tag
// @Accept json
// @Produce json
// @Param user_id path int true "article author id(기사 소유자)"
// @Param article_id path int true "article id(기사 ID)"
// @Success 200 {object} articledto.ReadArticleResponseWrapperDTO "success"
// @Failure 400 {object} types.ErrorResponse "입력값이 유효하지 않음"
// @Failure 404 {object} types.ErrorResponse "기사를 찾을 수 없음"
// @Failure 422 {object} types.ErrorResponse "요청을 제대로 수행하지 못함"
// @Failure 500 {object} types.ErrorResponse "네트워크 에러"
// @Router /user/{user_id}/article/{article_id} [get]
func ReadArticleByID(w http.ResponseWriter, r *http.Request) {

	userID, err := util.GetIntegerParam[int64](r, "user_id")
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	articleID, err := util.GetIntegerParam[int64](r, "article_id")
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	article, err := repository.ArticleRepo.GetByID(repository.DB, articleID, userID)
	if err != nil {
		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	} else if article == nil {
		responder.ErrorResponse(w, http.StatusNotFound, "article not found")
		return
	}

	responder.ReadArticleByIDResponse(w, *article)
}

// @Summary Read article by offset
// @Description page, limit을 이용해서 article을 읽어옴(limit는 없는 경우 10으로 사용)
// @Tags Article tag
// @Accept json
// @Produce json
// @Param page header int false "page"
// @Param limit header int false "limit"
// @Success 200 {object} articledto.ReadArticleByOffsetResponseWrapperDTO "success"
// @Failure 400 {object} types.ErrorResponse "입력값이 유효하지 않음"
// @Failure 422 {object} types.ErrorResponse "요청을 제대로 수행하지 못함"
// @Failure 500 {object} types.ErrorResponse "네트워크 에러"
// @Router /articles [get]
func ReadArticleByOffset(w http.ResponseWriter, r *http.Request) {

	page, limit, err := util.GetOffsetAndLimit(r)
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	articles, err := repository.ArticleRepo.GetByOffset(repository.DB, page, limit)
	if err != nil {
		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	responder.ReadArticleByOffsetResponse(w, articles)
}

// @Summary Update article
// @Description Update article
// @Tags Article tag
// @Accept json
// @Produce json
// @Param authorization header string true "jwt token"
// @Param user_id path int true "article author id(기사 소유자)"
// @Param article_id path int true "article id(기사 ID)"
// @Param updateArticleReq body  articledto.UpdateArticleRequestDTO true "updateArticleReq"
// @Success 200 {object}  articledto.UpdateArticleResponseWrapperDTO "success"
// @Failure 400 {object} types.ErrorResponse "bad request"
// @Failure 422 {object} types.ErrorResponse "요청을 제대로 수행하지 못함"
// @Failure 500 {object} types.ErrorResponse "network error"
// @Router /user/{user_id}/article/{article_id} [put]
func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	userID, err := util.GetIntegerParam[int64](r, "user_id")
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	articleID, err := util.GetIntegerParam[int64](r, "article_id")
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var updateArticleReq articledto.UpdateArticleRequestDTO

	err = json.NewDecoder(r.Body).Decode(&updateArticleReq)
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := ValidateUpdateArticleRequestDTO(updateArticleReq); err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	article, err := repository.ArticleRepo.UpdateWithTransaction(repository.DB, updateArticleReq, userID, articleID)
	if err != nil {
		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	} else if article == nil {
		responder.ErrorResponse(w, http.StatusBadRequest, "invalid request")
		return
	}

	responder.UpdateArticleResponse(w, *article, updateArticleReq.TagList)
}

// @Summary Delete article
// @Description Delete article
// @Tags Article tag
// @Accept json
// @Produce json
// @Param authorization header string true "jwt token"
// @Param user_id path int true "article author id(기사 소유자)"
// @Param article_id path int true "article id(기사 ID)"
// @Success 204 "success"
// @Failure 400 {object} types.ErrorResponse "bad request"
// @Failure 422 {object} types.ErrorResponse "요청을 제대로 수행하지 못함"
// @Failure 500 {object} types.ErrorResponse "network error"
// @Router /user/{user_id}/article/{article_id} [delete]
func DeleteArticle(w http.ResponseWriter, r *http.Request) {

	userID, err := util.GetIntegerParam[int64](r, "user_id")
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	articleID, err := util.GetIntegerParam[int64](r, "article_id")
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = repository.ArticleRepo.DeleteByID(repository.DB, articleID, userID)
	if err != nil {
		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// @Summary Read my article by offset
// @Description page, limit을 이용해서 내가 작성한 article을 읽어옴(limit는 없는 경우 10으로 사용)
// @Tags Article tag
// @Accept json
// @Produce json
// @Param authorization header string true "jwt token"
// @Param user_id path int true "내 user ID"
// @Param page header int false "page"
// @Param limit header int false "limit"
// @Success 200 {object} articledto.ReadArticleByOffsetResponseWrapperDTO "success"
// @Failure 400 {object} types.ErrorResponse "입력값이 유효하지 않음"
// @Failure 422 {object} types.ErrorResponse "요청을 제대로 수행하지 못함"
// @Failure 500 {object} types.ErrorResponse "네트워크 에러"
// @Router /user/{user_id}/articles [get]
func ReadMyArticleByOffset(w http.ResponseWriter, r *http.Request) {
	userID, err := util.GetIntegerParam[int64](r, "user_id")
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	page, limit, err := util.GetOffsetAndLimit(r)
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	articles, err := repository.ArticleRepo.GetByUserAndOffset(repository.DB, page, limit, userID)
	if err != nil {
		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	responder.ReadArticleByOffsetResponse(w, articles)
}
