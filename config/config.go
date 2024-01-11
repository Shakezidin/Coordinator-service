package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Host                string `mapstructure:"HOST"`
	User                string `mapstructure:"USER"`
	Password            string `mapstructure:"PASSWORD"`
	Database            string `mapstructure:"DATABASE"`
	Port                string `mapstructure:"PORT"`
	Sslmode             string `mapstructure:"SSLMODE"`
	GRPCCOORDINATORPORT string `mapstructure:"GRPCCOORDINATORPORT"`
	SID                 string `mapstructure:"SID"`
	TOKEN               string `mapstructure:"TOKEN"`
	SERVICETOKEN        string `mapstructure:"SERVICETOKEN"`
	SECRETKEY           string `mapstructure:"SECRETKEY"`
	REDISHOST           string `mapstructure:"REDISHOST"`
}

func LoadConfig() *Config {
	var config Config
	viper.SetConfigFile("../../.env")
	err := viper.ReadInConfig()
	err = viper.Unmarshal(&config)

	if err != nil {
		log.Fatal("Error while loading configure")
	}
	return &config
}
