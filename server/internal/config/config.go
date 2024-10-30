package config

import (
	"strings"

	_ "github.com/go-playground/validator"
	"github.com/spf13/viper"
	"github.com/subliker/server/internal/logger"
	"github.com/subliker/server/internal/validation"
)

type Config struct {
	ItemStore  ItemStore  `validate:"required" mapstructure:"item"`
	PhotoStore PhotoStore `validate:"required" mapstructure:"photo"`
	App        App        `validate:"required" mapstructure:"app"`
}

type ItemStore struct {
	Name     string `validate:"required" mapstructure:"name"`
	Port     int    `validate:"required" mapstructsure:"port"`
	Host     string `validate:"required" mapstructure:"host"`
	User     string `validate:"required" mapstructure:"user"`
	Password string `validate:"required" mapstructure:"password"`
}

type PhotoStore struct {
	Access string `validate:"required" mapstructure:"access"`
	Secret string `validate:"required" mapstructure:"secret"`
}

type App struct {
	Server Server `validate:"required" mapstructure:"server"`
}

type Server struct {
	Port int `validate:"required" mapstructure:"port"`
}

func init() {
	// viper setup
	viper.SetConfigFile("config.toml")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.SafeWriteConfig()

	// to env work
	viper.SetDefault("APP.SERVER.PORT", 8080)

	viper.SetDefault("ITEM.PORT", "")
	viper.SetDefault("ITEM.NAME", 3306)
	viper.SetDefault("ITEM.HOST", "")
	viper.SetDefault("ITEM.USER", "")
	viper.SetDefault("ITEM.PASSWORD", "")

	viper.SetDefault("PHOTO.ACCESS", "")
	viper.SetDefault("PHOTO.SECRET", "")

	viper.SetEnvPrefix("SERVER")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()

	// reading config
	if err := viper.ReadInConfig(); err != nil {
		logger.Zap.Fatalf("error reading config: %s", err)
	}
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
