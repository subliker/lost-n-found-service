package config

import (
	"strings"

	"github.com/spf13/viper"
	"github.com/subliker/server/internal/logger"
	"github.com/subliker/server/internal/validation"
)

type Config struct {
	ItemStore  ItemStore  `validate:"required" mapstructure:"item"`
	PhotoStore PhotoStore `validate:"required" mapstructure:"photo"`
	Server     Server     `validate:"required" mapstructure:"server"`
}

type ItemStore struct {
	Name      string `validate:"required" mapstructure:"name"`
	Port      int    `validate:"required" mapstructure:"port"`
	Host      string `validate:"required" mapstructure:"host"`
	User      string `validate:"required" mapstructure:"user"`
	Password  string `validate:"required" mapstructure:"password"`
	Migration bool   `validate:"required" mapstructure:"migration"`
}

type PhotoStore struct {
	Access string `validate:"required" mapstructure:"access"`
	Secret string `validate:"required" mapstructure:"secret"`
}

type Server struct {
	Port int `validate:"required" mapstructure:"port"`
}

func init() {
	// config file setup
	viper.SetConfigFile("./config/config.toml")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")
	viper.SafeWriteConfig()

	// reading config
	if err := viper.ReadInConfig(); err != nil {
		logger.Zap.Fatalf("error reading config: %s", err)
	}

	// reading environments
	viper.AutomaticEnv()

	// env and default binding
	viper.SetDefault("server.port", 8080)

	viper.SetDefault("item.port", 3306)
	viper.BindEnv("item.name")
	viper.SetDefault("item.host", "localhost")
	viper.BindEnv("item.user")
	viper.BindEnv("item.password")
	viper.SetDefault("item.migration", false)

	viper.BindEnv("photo.access")
	viper.BindEnv("photo.secret")

	// env setup
	viper.SetEnvPrefix("SERVER")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

}

func Get() Config {
	cfg := Config{}
	if err := viper.Unmarshal(&cfg); err != nil {
		logger.Zap.Fatalf("config unmarshal error: %s", err)
	}

	// config validation
	err := validation.V.Struct(cfg)
	if err != nil {
		logger.Zap.Fatalf("config validation error: %s", err)
	}

	return cfg
}
