package config

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Global AppConfig

type AppConfig struct {
	SRV_HOST              string        `mapstructure:"SRV_HOST"`
	SRV_PORT              string        `mapstructure:"SRV_PORT"`
	LOCAL_ENV             string        `mapstructure:"LOCAL_ENV"`
	LOG_LEVEL             int           `mapstructure:"LOG_LEVEL"`
	DATABASE_PSQL_URL     string        `mapstructure:"DATABASE_PSQL_URL"`
	TOKEN_TYPE            string        `mapstructure:"TOKEN_TYPE"`
	TOKEN_SYMMETRIC_KEY   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	ACCESS_TOKEN_DURATION time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	PASS_SALT             string        `mapstructure:"ACCESS_TOKEN_DURATION"`
}

func LoadConfig(path string) (config AppConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("mauth")
	viper.SetConfigType("env")

	//Override config file with env variables
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		logrus.Errorf("Config file error: %q", err)
		return AppConfig{}, err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		logrus.Errorf("Config file error: %q", err)
		return AppConfig{}, err
	}

	return config, nil
}

func init() {
	Global, _ = LoadConfig("../../")
}
