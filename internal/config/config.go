package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	HTTP struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}

	GRPC struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}

	Postgres struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Db       string `yaml:"db"`
	}

	Redis struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Password string `yaml:"password"`
		Db       int    `yaml:"db"`
	}
}

func NewConfig(configPath string) (*Config, error) {
	config := &Config{}

	viper.SetConfigFile(configPath)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config.HTTP.Host = viper.GetString("http.host")
	config.HTTP.Port = viper.GetString("http.port")
	config.GRPC.Host = viper.GetString("grpc.host")
	config.GRPC.Port = viper.GetString("grpc.port")
	config.Postgres.Host = viper.GetString("postgres-prod.host")
	config.Postgres.Port = viper.GetString("postgres-prod.port")
	config.Postgres.User = viper.GetString("postgres-prod.user")
	config.Postgres.Password = viper.GetString("postgres-prod.password")
	config.Postgres.Db = viper.GetString("postgres-prod.db")
	config.Redis.Host = viper.GetString("redis-prod.host")
	config.Redis.Port = viper.GetString("redis-prod.port")
	config.Redis.Password = viper.GetString("redis-prod.password")
	config.Redis.Db = viper.GetInt("redis-prod.db")

	return config, nil
}
