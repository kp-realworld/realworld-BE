package article

import (
	"encoding/json"
	"net/http"

	articledto "github.com/hotkimho/realworld-api/controller/dto/article"
	"github.com/hotkimho/realworld-api/repository"
	"github.com/hotkimho/realworld-api/responder"
	"github.com/hotkimho/realworld-api/util"
)

// @Summary Create article
// @Description Create article
// @Tags Article tag
// @Accept json
// @Produce json
// @Param authorization header string true "jwt token"
// @Param createArticleReq body CreateArticleRequestDTO true "createArticleReq"
// @Success 201 {object} CreateArticleResponseWrapperDTO "success"
// @Failure 400 {object} types.ErrorResponse "bad request"
// @Failure 422 {object} types.ErrorResponse "요청을 제대로 수행하지 못함"
// @Failure 500 {object} types.ErrorResponse "network error"
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

	_, err = repository.ArticleRepo.CreateWithTransaction(
		repository.DB,
		createArticleReq,
		userID)
	if err != nil {
		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	responder.CreateArticleResponse(w, createArticleReq)
}

// @Summary Read article by id
// @Description Read article by id
// @Tags Article tag
// @Accept json
// @Produce json
// @Param authorization header string true "jwt token"
// @Param user_id path int true "user id"
// @Param article_id path int true "article id"
// @Success 200 {object} articledto.ReadArticleResponseWrapperDTO "success"
// @Failure 400 {object} types.ErrorResponse "bad request"
// @Failure 404 {object} types.ErrorResponse "article not found"
// @Failure 422 {object} types.ErrorResponse "요청을 제대로 수행하지 못함"
// @Failure 500 {object} types.ErrorResponse "network error"
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
