package config

import (
	_ "embed"
	"os"

	"github.com/spf13/viper"
)



type Conf struct {
	Module      string `yaml:"module"`
	Env         string `yaml:"env"`
	Debug       bool   `yaml:"debug"`
	ServerHost  string `yaml:"serverhost,omitempty" default:"localhost"`
	ServerPort  int32  `yaml:"serverport"`
}

var appConf *Conf

// Config 读取app应用配置

func Config() (*Conf, error) {
	// Environment variables
	// viper.AutomaticEnv()
	viper.AddConfigPath("./config/")
	// viper.SetEnvPrefix("APP")
	// viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	// Configuration file
	viper.SetConfigType("yml")

	// Read configuration
	if err := viper.ReadInConfig(); err != nil && !os.IsNotExist(err) {
		panic(err)
	  }

	// Unmarshal the configuration
	
	if err := viper.Unmarshal(&appConf); err != nil {
		return nil, err
	}
	return appConf, nil
}

