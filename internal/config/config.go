package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	HTTP HTTPConfig
}

type HTTPConfig struct {
	Host               string        `mapstructure:"host"`
	Port               string        `mapstructure:"port"`
	ReadTimeout        time.Duration `mapstructure:"readTimeout"`
	WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
	MaxHeaderMegabytes int           `mapstructure:"maxHeaderBytes"`
}

func Init(path string) (*Config, error) {
	viper.SetConfigName("main")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./configs/")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	var cfg Config
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return nil, err
	}
	return &cfg, nil
}