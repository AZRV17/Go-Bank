package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
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
}

func NewConfig(configPath string) (*Config, error) {
	config := &Config{}

	viper.SetConfigFile(configPath)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config.Server.Host = viper.GetString("http.host")
	config.Server.Port = viper.GetString("http.port")
	config.Postgres.Host = viper.GetString("postgres-dev.host")
	config.Postgres.Port = viper.GetString("postgres-dev.port")
	config.Postgres.User = viper.GetString("postgres-dev.user")
	config.Postgres.Password = viper.GetString("postgres-dev.password")
	config.Postgres.Db = viper.GetString("postgres-dev.db")

	return config, nil
}
