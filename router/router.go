package router

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"github.com/hotkimho/realworld-api/controller/auth"
	"github.com/hotkimho/realworld-api/controller/user"
	"github.com/hotkimho/realworld-api/env"
	"github.com/hotkimho/realworld-api/responder"
	"github.com/hotkimho/realworld-api/types"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

type Router struct {
	Server     *mux.Router
	CorsHandle *cors.Cors
}

func (m *Router) Init() {
	m.Server = mux.NewRouter()

	m.InitSwagger()
	m.InitCORS()

	m.AddRoute([][]*Route{
		{
			{
				Method:      "POST",
				Path:        "/user/signup",
				HandlerFunc: auth.SignUp,
			},
			{
				Method:      "POST",
				Path:        "/user/signin",
				HandlerFunc: auth.SignIn,
			},
			{
				Method:      "GET",
				Path:        "/heartbeat",
				HandlerFunc: auth.Heartbeat,
			},
			{
				Method:      "GET",
				Path:        "/user/{user_id}/profile",
				HandlerFunc: user.ReadUserProfile,
				Middleware:  []Middleware{AuthMiddleware},
			},
			{
				Method:      "PUT",
				Path:        "/user/{user_id}/profile",
				HandlerFunc: user.UpdateUserProfile,
				//Middleware:  []Middleware{authMiddleware},
			},
		},
	})
}

func (m *Router) AddRoute(routeMaps [][]*Route) {
	for _, routeMap := range routeMaps {
		for _, route := range routeMap {

			handler := route.HandlerFunc

			// 미드웨어가 있는 경우
			for i := len(route.Middleware) - 1; i >= 0; i-- {
				handler = route.Middleware[i](handler)
			}

			m.Server.Handle(route.Path, handler).Methods(route.Method)
		}
	}
}

func (m *Router) InitSwagger() {
	m.Server.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
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

		// todo : logging

		next(w, r)
	}
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")
		if token == "" {
			responder.ErrorResponse(w, http.StatusUnauthorized, "token is empty")
			//w.WriteHeader(401)
			return
		}

		parsedToken, err := jwt.ParseWithClaims(token, &types.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(env.Config.Auth.Secret), nil
		})
		if err != nil {
			responder.ErrorResponse(w, http.StatusUnauthorized, "token is invalid")
			return
		}

		if _, ok := parsedToken.Claims.(*types.JWTClaims); ok && parsedToken.Valid {

		} else {
			responder.ErrorResponse(w, http.StatusUnauthorized, "token is invalid")
			w.WriteHeader(401)
			return
		}

		next(w, r)
	}
}

func testMiddlewareA(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("a")

		next(w, r)
	}
}

func testMiddlewareB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("b")
		next(w, r)
	}

}

func testMiddlewareC(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("c")
		next(w, r)
	}
}
