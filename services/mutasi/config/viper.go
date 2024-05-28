package config

import "github.com/spf13/viper"

func NewViper() *viper.Viper {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("env")
	config.AddConfigPath(".")
	config.AddConfigPath("/etc/booking/")
	config.AutomaticEnv()
	err := config.ReadInConfig()
	if err != nil {
		panic(err)
	}

	return config
}
