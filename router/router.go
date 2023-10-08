package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hotkimho/realworld-api/controller"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

type Router struct {
	Server *mux.Router
}

func (m *Router) Init() {
	m.Server = mux.NewRouter()

	m.InitSwagger()

	m.AddRoute([][]*Route{
		{
			{
				Method:      "POST",
				Path:        "/heartbeat",
				HandlerFunc: controller.TestFunc,
				Middleware:  authMiddleware,
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

type Middleware func(http.HandlerFunc) http.Handler
type Route struct {
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
	Middleware  Middleware
}

func authMiddleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("auth middleware")
		next(w, r)
	})
}
