package config

import (
	"github.com/Diasisme/asssesment-march-ihsan.git/models"
	"github.com/spf13/viper"
)

// SERVICE :=
// 	DB_HOST :=
// 	DB_USER :=
// 	DB_PASSWORD :=
// 	DB_PORT :=
// 	DB_DATABASE := os.Getenv("POSTGRES_DB")
// 	SVC_PORT :=

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
	}

	return
}
