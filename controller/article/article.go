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
// @Router /article [post]
func CreateArticle(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("ctx_user_id").(int64)

	var createArticleReq articledto.CreateArticleRequestDTO
	err := json.NewDecoder(r.Body).Decode(&createArticleReq)
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
// @Param authorization header string false "로그인 한 경우, 토큰 전달(토큰이 없는 경우(로그아웃) 에러가 발생하지 않음)"
// @Param user_id path int true "article author id(기사 소유자)"
// @Param article_id path int true "article id(기사 ID)"
// @Success 200 {object} articledto.ReadArticleResponseWrapperDTO "success"
// @Failure 400 {object} types.ErrorResponse "입력값이 유효하지 않음"
// @Failure 404 {object} types.ErrorResponse "기사를 찾을 수 없음"
// @Failure 422 {object} types.ErrorResponse "요청을 제대로 수행하지 못함"
// @Failure 500 {object} types.ErrorResponse "네트워크 에러"
// @Router /user/{author_id}/article/{article_id} [get]
func ReadArticleByID(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("ctx_user_id").(int64)

	authorID, err := util.GetIntegerParam[int64](r, "author_id")
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

	// 팔로우 여부 확인(로그인한 경우에만)
	isFollowing, err := repository.FollowRepo.IsFollowing(repository.DB, userID, authorID)
	if err != nil {
		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// 게시글 좋아요 조회
	article.FavoriteCount, err = repository.ArticleLikeCountRepo.GetByArticle(repository.DB, article.ID)
	if err != nil {
		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	responder.ReadArticleByIDResponse(w, *article, isFollowing)
}

// @Summary Read article by offset
// @Description page, limit을 이용해서 article을 읽어옴(limit는 없는 경우 10으로 사용)
// @Tags Article tag
// @Accept json
// @Produce json
// @Param authorization header string false "로그인 한 경우, 토큰 전달(토큰이 없는 경우(로그아웃) 에러가 발생하지 않음)"
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Success 200 {object} articledto.ReadArticleByOffsetResponseWrapperDTO "success"
// @Failure 400 {object} types.ErrorResponse "입력값이 유효하지 않음"
// @Failure 422 {object} types.ErrorResponse "요청을 제대로 수행하지 못함"
// @Failure 500 {object} types.ErrorResponse "네트워크 에러"
// @Router /articles [get]
func ReadArticleByOffset(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value("ctx_user_id").(int64)

	page, limit, err := util.GetOffsetAndLimit(r)
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	articles, err := repository.ArticleRepo.GetByOffset(repository.DB, page, limit, userID)
	if err != nil {
		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	articleIDSlice := make([]int64, 0)
	for _, article := range articles {
		articleIDSlice = append(articleIDSlice, article.ID)
	}

	articleLikeCounts, err := repository.ArticleLikeCountRepo.GetByArticles(repository.DB, articleIDSlice)
	if err != nil {
		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	responder.ReadArticleByOffsetResponse(w, articles, articleLikeCounts)
}

// @Summary Update article
// @Description Update article
// @Tags Article tag
// @Accept json
// @Produce json
// @Param authorization header string true "jwt token"
// @Param article_id path int true "article id(기사 ID)"
// @Param updateArticleReq body  articledto.UpdateArticleRequestDTO true "updateArticleReq"
// @Success 200 {object}  articledto.UpdateArticleResponseWrapperDTO "success"
// @Failure 400 {object} types.ErrorResponse "bad request"
// @Failure 422 {object} types.ErrorResponse "요청을 제대로 수행하지 못함"
// @Failure 500 {object} types.ErrorResponse "network error"
// @Router /article/{article_id} [put]
func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("ctx_user_id").(int64)

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

	// 게시글 좋아요 조회
	article.FavoriteCount, err = repository.ArticleLikeCountRepo.GetByArticle(repository.DB, article.ID)
	if err != nil {
		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
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
// @Param article_id path int true "article id(기사 ID)"
// @Success 204 "success"
// @Failure 400 {object} types.ErrorResponse "bad request"
// @Failure 422 {object} types.ErrorResponse "요청을 제대로 수행하지 못함"
// @Failure 500 {object} types.ErrorResponse "network error"
// @Router /article/{article_id} [delete]
func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("ctx_user_id").(int64)

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
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Success 200 {object} articledto.ReadArticleByOffsetResponseWrapperDTO "success"
// @Failure 400 {object} types.ErrorResponse "입력값이 유효하지 않음"
// @Failure 422 {object} types.ErrorResponse "요청을 제대로 수행하지 못함"
// @Failure 500 {object} types.ErrorResponse "네트워크 에러"
// @Router /my/articles [get]
func ReadMyArticleByOffset(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("ctx_user_id").(int64)

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

	articleIDSlice := make([]int64, 0)
	for _, article := range articles {
		articleIDSlice = append(articleIDSlice, article.ID)
	}

	articleLikeCounts, err := repository.ArticleLikeCountRepo.GetByArticles(repository.DB, articleIDSlice)
	if err != nil {
		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	responder.ReadArticleByOffsetResponse(w, articles, articleLikeCounts)
}

// @Summary Read article by tag
// @Description tag를 이용해서 article을 읽어옴
// @Tags Article tag
// @Accept json
// @Produce json
// @Param authorization header string false "로그인 한 경우, 토큰 전달(토큰이 없는 경우(로그아웃) 에러가 발생하지 않음)"
// @Param tag query string true "tag"
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Success 200 {object} articledto.ReadArticleByOffsetResponseWrapperDTO "success"
// @Failure 400 {object} types.ErrorResponse "입력값이 유효하지 않음"
// @Failure 422 {object} types.ErrorResponse "요청을 제대로 수행하지 못함"
// @Failure 500 {object} types.ErrorResponse "네트워크 에러"
// @Router /articles/tag [get]
func ReadArticleByTag(w http.ResponseWriter, r *http.Request) {
	tag := r.URL.Query().Get("tag")

	page, limit, err := util.GetOffsetAndLimit(r)
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	articles, err := repository.ArticleRepo.GetByOffsetAndTag(repository.DB, page, limit, tag)
	if err != nil {
		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	articleIDSlice := make([]int64, 0)
	for _, article := range articles {
		articleIDSlice = append(articleIDSlice, article.ID)
	}

	articleLikeCounts, err := repository.ArticleLikeCountRepo.GetByArticles(repository.DB, articleIDSlice)
	if err != nil {
		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	responder.ReadArticleByOffsetResponse(w, articles, articleLikeCounts)
}

// @Summary Article like
// @Description Article like
// @Tags Article tag
// @Accept json
// @Produce json
// @Param authorization header string true "jwt token"
// @Param author_id path int true "article 작성자 ID"
// @Param article_id path int true "article id"
// @Success 200 {object} articledto.CreateArticleLikeResponseDTO "이미 좋아요한 경우(좋아요 처리)"
// @Success 201 {object} articledto.CreateArticleLikeResponseDTO "좋아요 성공 "
// @Failure 400 {object} types.ErrorResponse "user_id, article id가 유효하지 않음"
// @Failure 404 {object} types.ErrorResponse "기사를 찾지 못한 경우"
// @Failure 500 {object} types.ErrorResponse "network error"
// @Router /user/{author_id}/article/{article_id}/like [post]
func CreateArticleLike(w http.ResponseWriter, r *http.Request) {
	ctxUserID := r.Context().Value("ctx_user_id").(int64)

	userID, err := util.GetIntegerParam[int64](r, "author_id")
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	articleID, err := util.GetIntegerParam[int64](r, "article_id")
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = repository.ArticleRepo.ValidateArticleOwner(repository.DB, articleID, userID)
	if err != nil {
		responder.ErrorResponse(w, http.StatusNotFound, "article not found")
		return
	}

	// 이미 좋아요를 눌렀는지 확인
	like, err := repository.ArticleLikeRepo.IsLiked(repository.DB, articleID, ctxUserID)
	if err != nil {
		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if !like {
		err = repository.ArticleLikeRepo.CreateWithTransaction(repository.DB, articleID, ctxUserID)
		if err != nil {
			responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	article, err := repository.ArticleRepo.GetByID(repository.DB, articleID, ctxUserID)
	if err != nil {
		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	} else if article == nil {
		responder.ErrorResponse(w, http.StatusNotFound, "article not found")
		return
	}

	// 이미 좋아요를 누른 상태면 200, 아니면 201
	if like {
		responder.CreateArticleLikeResponse(w, http.StatusOK, *article)
	} else {
		responder.CreateArticleLikeResponse(w, http.StatusCreated, *article)
	}
}

// @Summary Article unlike
// @Description Article unlike
// @Tags Article tag
// @Accept json
// @Produce json
// @Param authorization header string true "jwt token"
// @Param user_id path int true "article 작성자 ID"
// @Param article_id path int true "article id"
// @Success 204 "success"
// @Failure 400 {object} types.ErrorResponse "user_id, article id가 유효하지 않음"
// @Failure 404 {object} types.ErrorResponse "기사를 찾지 못한 경우"
// @Failure 500 {object} types.ErrorResponse "network error"
// @Router /user/{author_id}/article/{article_id}/like [delete]
func DeleteArticleLike(w http.ResponseWriter, r *http.Request) {
	ctxUserID := r.Context().Value("ctx_user_id").(int64)

	userID, err := util.GetIntegerParam[int64](r, "author_id")
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	articleID, err := util.GetIntegerParam[int64](r, "article_id")
	if err != nil {
		responder.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = repository.ArticleRepo.ValidateArticleOwner(repository.DB, articleID, userID)
	if err != nil {
		responder.ErrorResponse(w, http.StatusNotFound, "article not found")
		return
	}

	err = repository.ArticleLikeRepo.DeleteWithTransaction(repository.DB, articleID, ctxUserID)
	if err != nil {
		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// @Summary 유저가 작성한 article 조회
// @Description 유저가 작성한 article 조회
// @Tags Article tag
// @Accept json
// @Produce json
// @Param authorization header string false "jwt token"
// @Param author_id path int true "article 작성자 ID"
// @Param page query int false "page number"
// @Param limit query int false "limit"
// @Success 200 {object} articledto.ReadArticlesByUserResponseWrapperDTO "success"
// @Failure 400 {object} types.ErrorResponse "user_id가 유효하지 않음"
// @Failure 500 {object} types.ErrorResponse "network error"
// @Router /user/{author_id}/articles [get]
func ReadArticlesByUserID(w http.ResponseWriter, r *http.Request) {
	userID, err := util.GetIntegerParam[int64](r, "author_id")
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

	articleIDSlice := make([]int64, 0)
	for _, article := range articles {
		articleIDSlice = append(articleIDSlice, article.ID)
	}

	articleLikeCounts, err := repository.ArticleLikeCountRepo.GetByArticles(repository.DB, articleIDSlice)
	if err != nil {
		responder.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	responder.ReadArticleByOffsetResponse(w, articles, articleLikeCounts)
}
