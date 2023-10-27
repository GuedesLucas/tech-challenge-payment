package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadAppConfig() (AppConfig, error) {
	var appConfig AppConfig

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return appConfig, err
	}

	if err := viper.Unmarshal(&appConfig); err != nil {
		return appConfig, err
	}
	fmt.Println("AQUI")
	fmt.Println(appConfig)
	return appConfig, nil
}
