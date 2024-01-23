package router

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"github.com/hotkimho/realworld-api/controller/auth"
	"github.com/hotkimho/realworld-api/env"
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
		},
	})
}

func (m *Router) AddRoute(routeMaps [][]*Route) {
	for _, routeMap := range routeMaps {
		for _, route := range routeMap {
			// 미들웨어 있는 경우 처리
			if route.Middleware != nil {
				m.Server.Handle(route.Path, route.Middleware(route.HandlerFunc)).Methods(route.Method)
			} else {
				m.Server.HandleFunc(route.Path, route.HandlerFunc).Methods(route.Method)
			}
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

func (m *Router) InitCORS() {
	m.CorsHandle = cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:3000",
			"https://localhost:3000",
			"http://kp-realworld.com",
			"https://kp-realworld.com",
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

type Middleware func(http.HandlerFunc) http.Handler
type Route struct {
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
	Middleware  Middleware
}

func authMiddleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")
		if token == "" {
			w.WriteHeader(401)
			return
		}

		parsedToken, err := jwt.ParseWithClaims(token, &types.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(env.Config.Auth.Secret), nil
		})
		if err != nil {
			w.WriteHeader(401)
			return
		}

		if _, ok := parsedToken.Claims.(*types.JWTClaims); ok && parsedToken.Valid {

		} else {
			w.WriteHeader(401)
			return
		}

		next(w, r)
	})
}
