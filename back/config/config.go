package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

var Config = ReadConfig("config.json")

type config struct {
	ApiPath string
	Database databaseConfig
	Server   serverConfig
}

type databaseConfig struct {
	Type     string
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type serverConfig struct {
	Address string
	Port    int
}

var defaultConfig config = config{
	ApiPath: "/api/v1",
	Server: serverConfig{
		Address: "0.0.0.0",
		Port: 7777,		
	},
	Database: databaseConfig{
		Type: "",
		Host: "",
		Port: 0,
		Username: "",
		Password: "",
		Database: "",
	},
}

func ReadConfig(fileName string) config {
	var config config
	configAsString, err := os.ReadFile(fileName)
	if err != nil {
		writeConfig(fileName, defaultConfig)
		updateConfigFromEnv(&defaultConfig)
		return defaultConfig
	}

	err = json.Unmarshal(configAsString, &config)
	if err != nil {
		return defaultConfig
	}

	updateConfigFromEnv(&config)
	return config
}

func writeConfig(filename string, config config) error {
	jsonData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}

func updateConfigFromEnv(conf *config){
	updateConfigField("DATABASE_HOST", func(value string){
		conf.Database.Host = value
	})

	updateConfigField("DATABASE_USERNAME", func(value string){
		conf.Database.Username = value
	})

	updateConfigField("DATABASE_PASSWORD", func(value string){
		conf.Database.Password = value
	})

	updateConfigField("DATABASE_NAME", func(value string){
		conf.Database.Database = value
	})

	updateConfigField("DATABASE_PORT", func(value string){
		if valueAsInteger, err := strconv.Atoi(value); err != nil{
			conf.Database.Port = valueAsInteger
			fmt.Print("err")
		}
	})
}

func updateConfigField(envKey string, updateFunc func(string)){
	if envValue := os.Getenv(envKey); envKey != "" {
		updateFunc(envValue)
	}
}