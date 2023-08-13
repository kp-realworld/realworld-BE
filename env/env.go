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

type ConfigSetting struct {
	Database DatabaseSetting `toml:"database"`
}

var Config ConfigSetting

func SetConfig(filepath string) error {
	//path, _ := os.Getwd()

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
