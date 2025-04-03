package config

import (
	"encoding/json"
	"os"
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

var defaultConfif config = config{
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
		writeConfig(fileName, defaultConfif)
		return defaultConfif
	}

	err = json.Unmarshal(configAsString, &config)
	if err != nil {
		return defaultConfif
	}

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


