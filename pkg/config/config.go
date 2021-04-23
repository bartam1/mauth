package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var GConfig AppConfig

type AppConfig struct {
	SRV_HOST          string `mapstructure:"SRV_HOST"`
	SRV_PORT          string `mapstructure:"SRV_PORT"`
	LOCAL_ENV         string `mapstructure:"LOCAL_ENV"`
	LOG_LEVEL         string `mapstructure:"LOG_LEVEL"`
	DATABASE_PSQL_URL string `mapstructure:"DATABASE_PSQL_URL"`
}

func LoadConfig(path string) (config AppConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	//Override config file with env variables
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		logrus.Errorf("Config file error: %q", err)
		return
	}
}

func Init() {
	GConfig, _ = LoadConfig("../../")
}
