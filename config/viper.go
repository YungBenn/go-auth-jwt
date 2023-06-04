package config

import (
	"log"

	"github.com/spf13/viper"
)

type EnvVars struct {
	Port  string `mapstructure:"PORT"`
	DBurl string `mapstructure:"DB_URL"`
}

func LoadConfig() (config EnvVars, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal(err)
	}

	return
}
