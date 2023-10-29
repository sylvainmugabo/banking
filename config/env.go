package config

import (
	"github.com/spf13/viper"
	"github.com/sylvainmugabo/microservices-lib/logger"
)

type EnvConfigs struct {
	ServerAddress   string `mapstructure:"SERVER_ADDRESS"`
	ServerPort      string `mapstructure:"SERVER_PORT"`
	DatabaseUser    string `mapstructure:"DB_USER"`
	DatabasePwd     string `mapstructure:"DB_PASSWD"`
	DatabaseAddress string `mapstructure:"DB_ADDR"`
	DatabasePort    string `mapstructure:"DB_PORT"`
	DatabaseName    string `mapstructure:"DB_NAME"`

	/*

		DB_USER = "root"
		DB_PASSWD = "mugabo"
		DB_ADDR = "localhost"
		DB_PORT = "3306"
		DB_NAME = "banking"*/
}

func LoadEnvVariables() (config *EnvConfigs) {

	var conf *EnvConfigs

	viper.AddConfigPath(".")

	// Tell viper the name of your file
	viper.SetConfigName("app")

	// Tell viper the type of your file
	viper.SetConfigType("env")

	// Viper reads all the variables from env file and log error if any found
	if err := viper.ReadInConfig(); err != nil {
		logger.Fatal("Error reading env file")
	}

	// Viper unmarshals the loaded env varialbes into the struct
	if err := viper.Unmarshal(&conf); err != nil {
		logger.Fatal("Unable to Unmarshal")
	}
	return conf
}
