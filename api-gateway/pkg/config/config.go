package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	PORT                string `mapstructure:"PORT"`
	Auth_service_port   string `mapstructure:"AUTH_SERVICE_PORT"`
	Friend_service_Port string `mapstructure:"FRIEND_SERVICE_PORT"`
	Server_service_port string `mapstructure:"SERVER_SERVICE_PORT"`
	KafkaPort           string `mapstructure:"KAFKAPORT"`
	KafkaTopic          string `mapstructure:"TOPIC"`
	KafkaServerTopic    string `mapstructure:"KAFKASERVERTOPIC"`
	RedisDB             Redis
}

type Redis struct {
	RedisPort     string `mapstructure:"REDIS_PORT"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
	RedisURL      string `mapstructure:"REDIS_URL"`
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
