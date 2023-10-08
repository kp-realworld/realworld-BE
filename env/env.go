package env

import (
	"fmt"
	"github.com/spf13/viper"
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

type ConfigSetting struct {
	Database DatabaseSetting `toml:"database"`
	Redis    RedisSetting    `toml:"redis"`
}

var Config ConfigSetting

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

	fmt.Println(Config.Database)
	return nil
}
