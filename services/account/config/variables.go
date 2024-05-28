package config

import (
	"github.com/Diasisme/asssesment-march-ihsan.git/models"
	"github.com/spf13/viper"
)

func NewEnvVar(viper *viper.Viper) (envVar models.VarEnviroment) {

	envVar = models.VarEnviroment{
		Host:        viper.GetString("POSTGRES_HOST"),
		Port:        viper.GetInt32("POSTGRES_DB_PORT"),
		User:        viper.GetString("POSTGRES_USER"),
		Pass:        viper.GetString("POSTGRES_PASSWORD"),
		DB:          viper.GetString("POSTGRES_DB"),
		ServicePort: viper.GetString("SVC_PORT"),
		Service:     viper.GetString("CONTAINER_ID_NAME"),
		RedisHost:   viper.GetString("REDIS_HOST"),
		RedisPort:   viper.GetString("REDIS_PORT"),
		RedisUser:   viper.GetString("REDIS_USER"),
		RedisPass:   viper.GetString("REDIS_PASS"),
	}

	return
}
