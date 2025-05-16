package config

import (
	"github.com/spf13/viper"
	"log"
)

var (
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	JWTSecret  string
)

func LoadEnv() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Can not load env file %v", err)
	}

	DBHost = viper.GetString("DB_HOST")
	DBUser = viper.GetString("DB_USER")
	DBPassword = viper.GetString("DB_PASSWORD")
	DBName = viper.GetString("DB_NAME")
	DBPort = viper.GetString("DB_PORT")
	JWTSecret = viper.GetString("JWT_SECRET")

	log.Println("Success to config env file")
}
