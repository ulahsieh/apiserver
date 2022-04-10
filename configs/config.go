package configs

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Configuration struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port int
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Sid      string
}

// Load config.yaml 載入與讀取config檔
func LoadConfig() *Configuration {

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	// set config path
	viper.AddConfigPath(".")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath("./configs/config")

	var config Configuration

	// read config

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Read Config Failed!")
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Error("Unable to decode config file to struct")
	}

	return &config
}
