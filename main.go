package main

import (
	"flag"
	"fmt"
	"github.com/hotkimho/realworld-api/domain"
	"github.com/hotkimho/realworld-api/env"
	"github.com/hotkimho/realworld-api/router"
	"net/http"
)

func main() {

	config := flag.String("config", "env/local-env.toml", "config file path")
	flag.Parse()

	if err := env.SetConfig(*config); err != nil {
		fmt.Println(err)
		return
	}

	if err := domain.Init(); err != nil {
		fmt.Println(err)
		return
	}

	var router router.Router
	router.Init()

	//router.Handle("/heartbeat", authMiddleware(TestFunc)).Methods("GET")

	http.ListenAndServe(":8080", router.Server)
	//fmt.Println(viper.GetString("database.host"))
}
