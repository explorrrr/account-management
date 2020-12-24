package config

import (
	"log"
	"path/filepath"
	"github.com/spf13/viper"
)

var config *viper.Viper

// Init read config file by environment name
func Init(env string) {
	var err error
	config = viper.New()
	config.SetConfigType("yml")
	config.SetConfigName(env)
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")
	err = config.ReadInConfig()

	if err != nil {
		log.Fatal("error on parsing configuration file")
	}
}

func relativePath(basedir string, path *string) {
	p := *path
	if len(p) > 0 && p[0] != '/' {
		*path = filepath.Join(basedir, p)
	}
}

// GetConfig returns viper config
func GetConfig() *viper.Viper {
	return config
}
