package main

import (
	"flag"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/hotkimho/realworld-api/redis"

	//_ "github.com/swaggo/http-swagger/example/gorilla/docs"
	"net/http"

	"github.com/sirupsen/logrus"

	_ "github.com/hotkimho/realworld-api/docs"
	"github.com/hotkimho/realworld-api/env"
	"github.com/hotkimho/realworld-api/repository"
	"github.com/hotkimho/realworld-api/router"
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

	// 환경 설정
	if err := env.SetConfig(*config); err != nil {
		fmt.Println(err)
		return
	}

	// Sentry 설정
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              env.Config.Sentry.DSN,
		EnableTracing:    true,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// 타임존 설정
	if err := env.InitTimeZone(); err != nil {
		fmt.Println(err)
		return
	}

	// RDB 설정
	if err := repository.Init(); err != nil {
		fmt.Println(err)
		return
	}

	// Redis 설정
	if err := redis.Init(); err != nil {
		fmt.Println(err)
		return
	}

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)

	var router router.Router
	router.Init()

	fmt.Println("start server :8080")
	http.ListenAndServe(":8080", router.CorsHandle.Handler(router.Server))
	fmt.Println("end server :8080")
	//fmt.Println(viper.GetString("database.host"))
}
