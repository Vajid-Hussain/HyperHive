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
	S3                  S3Bucket
}

type S3Bucket struct {
	AccessKeyID     string `mapstructure:"AccessKeyID"`
	AccessKeySecret string `mapstructure:"AccessKeySecret"`
	Region          string `mapstructure:"Region"`
	BucketName      string `mapstructure:"BucketName"`
}

type Redis struct {
	RedisPort     string `mapstructure:"REDIS_PORT"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
	RedisURL      string `mapstructure:"REDIS_HOST"`
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
	viper.Unmarshal(&c.S3)
	viper.Unmarshal(&c.RedisDB)

	return
}
