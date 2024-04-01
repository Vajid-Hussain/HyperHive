package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	PORT              string `mapstructure:"PORT"`
	Auth_service_port string `mapstructure:"AUTH_SERVICE_PORT"`
	Friend_service_Port string `mapstructure:"FRIEND_SERVICE_PORT"`
}

func InitConfig() (c *Config, err error) {

	viper.AddConfigPath("./")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	viper.Unmarshal(&c)
	return
}
