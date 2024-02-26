package router

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/hotkimho/realworld-api/controller/follow"

	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/hotkimho/realworld-api/controller/comment"

	"github.com/hotkimho/realworld-api/controller/article"
	"github.com/hotkimho/realworld-api/controller/auth"
	"github.com/hotkimho/realworld-api/controller/user"
	"github.com/hotkimho/realworld-api/env"
	"github.com/hotkimho/realworld-api/responder"
	"github.com/hotkimho/realworld-api/types"
)

type Router struct {
	Server        *mux.Router
	CorsHandle    *cors.Cors
	SentryHandler *sentryhttp.Handler
}

func (m *Router) Init() {
	m.Server = mux.NewRouter()
	m.SentryHandler = sentryhttp.New(sentryhttp.Options{})
	//
	//m.Server.HandleFunc("/foo", sentryHandler.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
	//	var aa *string
	//	println(aa)
	//	w.Header().Set("Content-Type", "application/json")
	//	w.WriteHeader(http.StatusInternalServerError)
	//}))

	m.InitSwagger()
	m.InitCORS()
	m.AddRoute(ArticleRouter)
	m.AddRoute(CommentRouter)
	m.AddRoute(FollowRouter)
	m.AddRoute([][]*Route{
		{
			{
				Method:      "POST",
				Path:        "/user/signup",
				HandlerFunc: auth.SignUp,
				Middleware:  []Middleware{LoggingMiddleware},
			},
			{
				Method:      "POST",
				Path:        "/user/signin",
				HandlerFunc: auth.SignIn,
				Middleware:  []Middleware{LoggingMiddleware},
			},
			{
				Method:      "GET",
				Path:        "/heartbeat",
				HandlerFunc: auth.Heartbeat,
				Middleware:  []Middleware{LoggingMiddleware},
			},
			//{
			//	Method:      "GET",
			//	Path:        "/my/profile",
			//	HandlerFunc: user.ReadMyProfile,
			//	Middleware:  []Middleware{LoggingMiddleware, UserAuthMiddleware},
			//},
			{
				Method:      "GET",
				Path:        "/profile",
				HandlerFunc: user.ReadUserProfile,
				Middleware:  []Middleware{LoggingMiddleware, UserAuthMiddlewareWithoutVerify},
			},
			{
				Method:      "PUT",
				Path:        "/profile",
				HandlerFunc: user.UpdateUserProfile,
				Middleware:  []Middleware{LoggingMiddleware, UserAuthMiddleware},
			},
			{
				Method:      "GET",
				Path:        "/token-refresh",
				HandlerFunc: auth.RefreshToken,
				Middleware:  []Middleware{LoggingMiddleware},
			},
			{
				Method:      "POST",
				Path:        "/user/verify-email",
				HandlerFunc: auth.VerifyEmail,
				Middleware:  []Middleware{LoggingMiddleware},
			},
			{
				Method:      "POST",
				Path:        "/user/verify-username",
				HandlerFunc: auth.VerifyUsername,
				Middleware:  []Middleware{LoggingMiddleware},
			},
		},
	})
}

var ArticleRouter = [][]*Route{
	{
		{
			Method:      "POST",
			Path:        "/article",
			HandlerFunc: article.CreateArticle,
			Middleware:  []Middleware{LoggingMiddleware, UserAuthMiddleware},
		},
		{
			Method:      "GET",
			Path:        "/user/{author_id}/article/{article_id}",
			HandlerFunc: article.ReadArticleByID,
			Middleware:  []Middleware{LoggingMiddleware, UserAuthMiddlewareWithoutVerify},
		},
		{
			Method:      "PUT",
			Path:        "/article/{article_id}",
			HandlerFunc: article.UpdateArticle,
			Middleware:  []Middleware{LoggingMiddleware, UserAuthMiddleware},
		},
		{
			Method:      "DELETE",
			Path:        "/article/{article_id}",
			HandlerFunc: article.DeleteArticle,
			Middleware:  []Middleware{LoggingMiddleware, UserAuthMiddleware},
		},
		{
			Method:      "GET",
			Path:        "/articles",
			HandlerFunc: article.ReadArticleByOffset,
			Middleware:  []Middleware{LoggingMiddleware, UserAuthMiddlewareWithoutVerify},
		},
		{
			Method:      "GET",
			Path:        "/my/articles",
			HandlerFunc: article.ReadMyArticleByOffset,
			Middleware:  []Middleware{LoggingMiddleware, UserAuthMiddleware},
		},
		{
			Method:      "GET",
			Path:        "/articles/tag",
			HandlerFunc: article.ReadArticleByTag,
			Middleware:  []Middleware{LoggingMiddleware},
		},
		{
			Method:      "POST",
			Path:        "/user/{author_id}/article/{article_id}/like",
			HandlerFunc: article.CreateArticleLike,
			Middleware:  []Middleware{LoggingMiddleware, UserAuthMiddleware},
		},
		{
			Method:      "DELETE",
			Path:        "/user/{author_id}/article/{article_id}/like",
			HandlerFunc: article.DeleteArticleLike,
			Middleware:  []Middleware{LoggingMiddleware, UserAuthMiddleware},
		},
		{
			Method:      "GET",
			Path:        "/user/{author_id}/articles",
			HandlerFunc: article.ReadArticlesByUserID,
			Middleware:  []Middleware{LoggingMiddleware, UserAuthMiddlewareWithoutVerify},
		},
	},
}

var CommentRouter = [][]*Route{
	{
		{
			Method:      "POST",
			Path:        "/user/{author_id}/article/{article_id}/comment",
			HandlerFunc: comment.CreateComment,
			Middleware:  []Middleware{LoggingMiddleware, UserAuthMiddleware},
		},
		{
			Method:      "GET",
			Path:        "/user/{author_id}/article/{article_id}/comments",
			HandlerFunc: comment.ReadComments,
			Middleware:  []Middleware{LoggingMiddleware},
		},
		{
			Method:      "PUT",
			Path:        "/user/{author_id}/article/{article_id}/comment/{comment_id}",
			HandlerFunc: comment.UpdateComment,
			Middleware:  []Middleware{LoggingMiddleware, UserAuthMiddleware},
		},
		{
			Method:      "DELETE",
			Path:        "/user/{author_id}/article/{article_id}/comment/{comment_id}",
			HandlerFunc: comment.DeleteComment,
			Middleware:  []Middleware{LoggingMiddleware, UserAuthMiddleware},
		},
	},
}

var FollowRouter = [][]*Route{
	{
		{
			Method:      "POST",
			Path:        "/user/follow/{followed_id}",
			HandlerFunc: follow.CreateFollow,
			Middleware:  []Middleware{LoggingMiddleware, UserAuthMiddleware},
		},
		{
			Method:      "DELETE",
			Path:        "/user/follow/{followed_id}",
			HandlerFunc: follow.DeleteFollow,
			Middleware:  []Middleware{LoggingMiddleware, UserAuthMiddleware},
		},
	},
}

func (m *Router) AddRoute(routeMaps [][]*Route) {
	for _, routeMap := range routeMaps {
		for _, route := range routeMap {

			handler := route.HandlerFunc

			// 미드웨어가 있는 경우
			for i := len(route.Middleware) - 1; i >= 0; i-- {
				handler = route.Middleware[i](handler)
			}

			m.Server.Handle(route.Path, m.SentryHandler.Handle(handler)).Methods(route.Method)
		}
	}
}

func (m *Router) InitSwagger() {
	fmt.Println(env.Config.Swagger.Host)
	m.Server.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL(env.Config.Swagger.Host), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)
}

// 5173
func (m *Router) InitCORS() {
	m.CorsHandle = cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:3000",
			"https://localhost:3000",
			"http://localhost:5173",
			"https://localhost:5173",
			"http://kp-realworld.com",
			"https://kp-realworld.com",
			"http://*.kp-realworld.com",
			"https://*.kp-realworld.com",
		},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowedHeaders: []string{
			"Accept",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
			"X-CSRF-Token",
		},
		AllowCredentials: true,
	})

}

type Middleware func(handlerFunc http.HandlerFunc) http.HandlerFunc
type Route struct {
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
	Middleware  []Middleware
}

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		// 요청 본문 로깅을 위해 Body를 읽고 복사
		var bodyBytes []byte
		if r.Body != nil {
			bodyBytes, _ = io.ReadAll(r.Body) // 수정됨
		}
		// Body 내용을 복사한 후 원본 요청에 다시 쓰기
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // 수정됨

		// 요청 처리 전 로깅
		logrus.WithFields(logrus.Fields{
			"domain":   r.Host,
			"path":     r.URL.Path,
			"method":   r.Method,
			"body":     string(bodyBytes),
			"query":    r.URL.Query(),
			"clientIP": r.RemoteAddr,
		}).Info("Request received")
		logrus.Trace("Trace log")

		next.ServeHTTP(w, r)

		// 요청 처리 후 로깅 (예: 응답 시간)
		logrus.WithFields(logrus.Fields{
			"path":         r.URL.Path,
			"method":       r.Method,
			"responseTime": time.Since(start).String(),
		}).Info("Request processed")
	}
}

func UserAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")
		if token == "" {
			responder.ErrorResponse(w, http.StatusUnauthorized, "token is empty")
			return
		}

		parsedToken, err := jwt.ParseWithClaims(token, &types.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(env.Config.Auth.Secret), nil
		})
		if err != nil {
			// 토큰 만료
			if strings.Compare(err.Error(), types.ERR_EXPIRED_TOKEN) == 0 {
				responder.ErrorResponse(w, http.StatusUnauthorized, "Token expired")
				return
			}
			responder.ErrorResponse(w, http.StatusUnauthorized, "token is invalid")
			return
		}

		var userID int64
		if claim, ok := parsedToken.Claims.(*types.JWTClaims); ok && parsedToken.Valid {
			if claim.UserID <= 0 {
				responder.ErrorResponse(w, http.StatusUnauthorized, "token is invalid")
				return
			}
			userID = claim.UserID
			// 필요 시, api path의 user_id 와 토큰의 user_id 를 비교
			//vars := mux.Vars(r)
			//val, ok := vars["user_id"]
			//if !ok {
			//	responder.ErrorResponse(w, http.StatusBadRequest, "user_id is empty")
			//	return
			//}
			//
			//currentUserID, err := strconv.ParseInt(val, 10, 64)
			//if err != nil {
			//	responder.ErrorResponse(w, http.StatusBadRequest, "user_id is invalid")
			//	return
			//}
		} else {
			responder.ErrorResponse(w, http.StatusUnauthorized, "token is invalid")
			w.WriteHeader(401)
			return
		}

		if userID <= 0 {
			responder.ErrorResponse(w, http.StatusUnauthorized, "token is invalid")
			return
		}

		ctx := context.WithValue(r.Context(), "ctx_user_id", userID)
		next(w, r.WithContext(ctx))
	}
}

// 토큰이 없어도 인증 검사를 하지 않음

func UserAuthMiddlewareWithoutVerify(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var userID int64
		token := r.Header.Get("Authorization")
		if token != "" {
			parsedToken, err := jwt.ParseWithClaims(token, &types.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(env.Config.Auth.Secret), nil
			})
			if err == nil {
				if claim, ok := parsedToken.Claims.(*types.JWTClaims); ok && parsedToken.Valid {
					userID = claim.UserID
				}
			}
		}

		ctx := context.WithValue(r.Context(), "ctx_user_id", userID)

		next(w, r.WithContext(ctx))
	}
}
