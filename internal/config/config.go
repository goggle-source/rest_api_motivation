package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Env    string     `mapstructure:"env"`
	Server HTTPServer `mapstructure:"App"`
	Db     DataBase   `mapstructure:"Database"`
}

type HTTPServer struct {
	ReadTimeout    time.Duration `mapstructure:"readtimeout"`
	WriteTimeout   time.Duration `mapstructure:"writetimeout"`
	RequestTimeout time.Duration `mapstructure:"requesttimeout"`
	Port           string        `mapstructure:"port"`
}

type DataBase struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	NameDB   string `mapstructure:"nameDB"`
	PortDB   string `mapstructure:"PortDB"`
}

func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("C:\\Users\\admin\\rest_api_motivation\\configs")
	viper.AddConfigPath("../../configs")
	var cfg Config
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	err := viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}
