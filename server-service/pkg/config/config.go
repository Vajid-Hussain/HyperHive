package config_server_service

import "github.com/spf13/viper"

type DataBasePostgres struct {
	Port               string `mapstructure:"PORT_SERVER_SERVICE"`
	DBConeectionString string `mapstructure:"POSTGRES_DB"`
	DBName             string `mapstructure:"DATABASENAME"`
	User               string `mapstructure:"USER"`
	UserPassword       string `mapstructure:"PASSWORD"`
	Host               string `mapstructure:"HOST"`
}

type Config struct {
	DB DataBasePostgres
}

func ConfigInit() (config *Config, err error) {
	viper.AddConfigPath("./")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config.DB)
	if err != nil {
		return nil, err
	}

	return config, nil
}
