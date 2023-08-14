package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	Server *mux.Router
}

func (m *Router) Init() {
	m.Server = mux.NewRouter()

	m.AddRoute([][]*Route{
		{
			{
				Method:      "GET",
				Path:        "/heartbeat",
				HandlerFunc: TestFunc,
				Middleware:  authMiddleware,
			},
		},
	})
}

func (m *Router) AddRoute(routeMaps [][]*Route) {
	for _, routeMap := range routeMaps {
		for _, route := range routeMap {
			if route.Middleware != nil {
				m.Server.Handle(route.Path, route.Middleware(route.HandlerFunc)).Methods(route.Method)
			} else {
				m.Server.HandleFunc(route.Path, route.HandlerFunc).Methods(route.Method)
			}
		}
	}
}

type Middleware func(http.HandlerFunc) http.Handler
type Route struct {
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
	Middleware  Middleware
}

func TestFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test")
}

func authMiddleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("auth middleware")
		next(w, r)
	})
}
