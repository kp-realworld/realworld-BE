package env

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type DatabaseSetting struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

type RedisSetting struct {
	Host     string
	Port     int
	Password string
	DB       int
}

type AuthSetting struct {
	Secret             string
	RefreshTokenExpire int
	AccessTokenExpire  int
	Issuer             string
}

type SwaggerSetting struct {
	Host string
}

type SentrySetting struct {
	DSN string
}

type ConfigSetting struct {
	Database DatabaseSetting `toml:"database"`
	Redis    RedisSetting    `toml:"redis"`
	Auth     AuthSetting     `toml:"auth"`
	Swagger  SwaggerSetting  `toml:"swagger"`
	Sentry   SentrySetting   `toml:"sentry"`
}

var Config ConfigSetting

var Seoul *time.Location

func SetConfig(filepath string) error {
	//path, _ := os.Getwd()
	fmt.Println(filepath)
	if filepath == "" {
		filepath = "config/local-env.toml"
	}

	viper.SetConfigFile(filepath)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	// set config
	if err := viper.Unmarshal(&Config); err != nil {
		return err
	}

	return nil
}

func InitTimeZone() error {
	var err error
	Seoul, err = time.LoadLocation("Asia/Seoul")
	if err != nil {
		return err
	}

	return nil
}
