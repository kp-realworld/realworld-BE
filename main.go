package main

import (
	"flag"
	"fmt"
	_ "github.com/hotkimho/realworld-api/docs"
	"github.com/hotkimho/realworld-api/env"
	"github.com/hotkimho/realworld-api/repository"
	"github.com/hotkimho/realworld-api/router"
	//_ "github.com/swaggo/http-swagger/example/gorilla/docs"
	"net/http"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {

	config := flag.String("config", "config/local-env.toml", "config file path")
	flag.Parse()

	if err := env.SetConfig(*config); err != nil {
		fmt.Println(err)
		return
	}

	if err := repository.Init(); err != nil {
		fmt.Println(err)
		return
	}

	var router router.Router
	router.Init()

	fmt.Println("start server :8080")
	http.ListenAndServe(":8080", router.Server)
	fmt.Println("end server :8080")
	//fmt.Println(viper.GetString("database.host"))
}
