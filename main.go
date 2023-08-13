package main

import (
	"flag"
	"fmt"
	"github.com/hotkimho/realworld-api/domain"
	"github.com/hotkimho/realworld-api/env"
)

func main() {

	config := flag.String("config", "env/local-env.toml", "config file path")
	flag.Parse()

	fmt.Println(*config)
	if err := env.SetConfig(*config); err != nil {
		fmt.Println(err)
		return
	}

	if err := domain.Init(); err != nil {
		fmt.Println(err)
		return
	}
	//viper.SetConfigFile("env/local-env.toml")
	////viper.AddConfigPath("env")
	////viper.SetConfigType("toml")
	//
	//if err := viper.ReadInConfig(); err != nil {
	//	panic(fmt.Errorf("Fatal error config file: %s", err))
	//}

	fmt.Println("에러 안남 ㅋㅋ")
	//fmt.Println(viper.GetString("database.host"))
}
