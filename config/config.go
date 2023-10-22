package config

import (
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	ServiceName string `yaml:"serviceName" default:"pesa-api"`
	Environment string `yaml:"env" default:"dev"`
	Port        int    `yaml:"port" default:"8080" required:"true"`
	Host        string `yaml:"host" default:"localhost"`
	DBHost      string `yaml:"dBHost" default:"localhost"`
	DBPort      int    `yaml:"dBPort" default:"5432"`
	DBUserName  string `yaml:"dBUserName" default:"lawrence"`
	DBPassword  string `yaml:"dBPassword" default:"lawrence"`
	DBName      string `yaml:"dBName" default:"pesa_api"`
	DBLogMode   int    `yaml:"dBLogMode" default:"3"`
}

func New() *Config {
	configFile := "config.yml"
	confPath, err := os.Getwd()
	if err != nil {
		log.Error("error getting a working directory:%v", err)
		return nil
	}
	configPath := fmt.Sprintf("%s/config/%s", confPath, configFile)
	viper.SetConfigFile(configPath)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		return nil
	}

	cfg := Config{}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
		return nil
	}
	return &cfg
}

func GetDatabaseConnection() string {
	c := New()
	conn := fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%d",
		c.DBHost,
		c.DBName,
		c.DBUserName,
		c.DBPassword,
		c.DBPort)
	return conn
}
