package config

import (
	"strings"

	_ "github.com/go-playground/validator"
	"github.com/spf13/viper"
	"github.com/subliker/server/internal/logger"
	"github.com/subliker/server/internal/validation"
)

type Config struct {
	DB  DBConfig  `validate:"required" mapstructure:"db"`
	App AppConfig `validate:"required" mapstructure:"app"`
}

type DBConfig struct {
	Name     string `validate:"required" mapstructure:"name"`
	Port     int    `validate:"required" mapstructsure:"port"`
	Host     string `validate:"required" mapstructure:"host"`
	User     string `validate:"required" mapstructure:"user"`
	Password string `validate:"required" mapstructure:"password"`
}

type AppConfig struct {
	Server ServerConfig `validate:"required" mapstructure:"server"`
}

type ServerConfig struct {
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

	viper.SetDefault("DB.PORT", "")
	viper.SetDefault("DB.NAME", 3306)
	viper.SetDefault("DB.HOST", "")
	viper.SetDefault("DB.USER", "")
	viper.SetDefault("DB.PASSWORD", "")

	viper.SetEnvPrefix("rest")
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
