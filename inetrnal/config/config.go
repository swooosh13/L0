package config

import (
	"fmt"
	"path/filepath"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	PostgresDB struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Database string `mapstructure:"database"`
		Timeout  int    `mapstructure:"timeout"`
		MaxConns int    `mapstructure:"max_conns"`
	} `mapstructure:"pg_db"`
	Listen struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
	} `mapstructure:"listen"`
	Stan struct {
		ReceiveChan string ``
		ClusterId   string `mapstructure:"cluster_id"`
		ClientId    string `mapstructure:"client_id"`
	} `mapstructure:"stan"`
}

var instance *Config
var once sync.Once

func GetConfig() (*Config, error) {
	once.Do(func() {
		cfgName := "local"

		viper.SetConfigName(cfgName)
		viper.SetConfigType("yml")

		dirPath, err := filepath.Abs("./configs")
		if err != nil {
			return
		}

		viper.AddConfigPath(dirPath)
		err = viper.ReadInConfig()
		if err != nil {
			return
		}

		instance = &Config{}

		err = viper.Unmarshal(instance)
		if err != nil {
			return
		}
	})

	if instance == nil {
		return nil, fmt.Errorf("error parsing config")
	}

	return instance, nil
}
