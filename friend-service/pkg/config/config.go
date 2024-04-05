package config_friend_server

import (
	"github.com/spf13/viper"
)

type Frend_service struct {
	Friend_Service_Port string `mapstructure:"FRIEND_SERVICE_PORT"`
}

type Auth_service struct {
	Auth_Service_port string `mapstructure:"AUTH_SERVICE_PORT"`
}

type DataBase struct {
	Port               string `mapstructure:"PORT_AUTH_SVC"`
	DBConeectionString string `mapstructure:"DBCONNECTION"`
	DBName             string `mapstructure:"DATABASENAME"`
	User               string `mapstructure:"USER"`
	UserPassword       string `mapstructure:"PASSWORD"`
	Host               string `mapstructure:"HOST"`
}

type Kafka struct {
	KafkaPort  string `mapstructure:"KAFKAPORT"`
	KafkaTopic string `mapstrucuture:"TOPIC"`
}

type Config struct {
	Friend Frend_service
	DB     DataBase
	Auth   Auth_service
	Kafka  Kafka
}

func InitConfig() (*Config, error) {
	var config = Config{}
	viper.AddConfigPath("./")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	viper.Unmarshal(&config.DB)
	viper.Unmarshal(&config.Friend)
	viper.Unmarshal(&config.Auth)

	return &config, nil
}
