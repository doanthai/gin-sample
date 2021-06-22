package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	App          App           `yaml:"app"`
	Database     Database      `yaml:"database"`
	TokenSetting TokenSettings `yaml:"token_settings"`
}

type Database struct {
	URI             string `yaml:"uri" env:"SQL_URI"`
	TimeOut         int    `yaml:"time_out" env:"SQL_TIME_OUT"`
	MaxIdleConns    int    `yaml:"max_idle_conns" env:"SQL_MAX_IDLE_CONNS"`
	MaxOpenConns    int    `yaml:"max_open_conns" env:"SQL_MAX_OPEN_CONNS"`
	ConnMaxLifeTime int64  `yaml:"conn_max_life_time" env:"SQL_CONN_MAX_LIFE_TIME"`
	Debug           bool   `yaml:"debug" env:"SQL_DEBUG"`
}

type TokenSettings struct {
	Issuer      string `yaml:"issuer" env:"TOKEN_ISSUER"`
	TimeExpired int64  `yaml:"time_expired" env:"TOKEN_TIME_EXPIRED"`
	SecretKey   string `yaml:"secret_key" env:"TOKEN_SECRET_KEY"`
}

type App struct {
	Port   string `yaml:"port"`
	Secret string `yaml:"secret"`
}

var config Config

func LoadConfig(env string, path string) {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	switch env {
	case "test", "local", "stg", "prod":
		viper.SetConfigName(env)
	default:
		panic(fmt.Sprint("Env name incorrect"))
	}
	err := viper.ReadInConfig()
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func GetConfig() Config {
	return config
}
