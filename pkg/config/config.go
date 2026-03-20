package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config represents the application configuration
// @Description config model
// @receiver c
// @return Config
// @example {"port": 8080, "database": {"username": "user", "password": "pass", "name": "db"}}
type Config struct {
	Port     int    `json:"port"`
	Database Database `json:"database"`
}

// Database represents the database configuration
// @Description database model
// @receiver d
// @return Database
type Database struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

// InitConfig initializes the configuration
func InitConfig() {
	cfg := &Config{}
	cfgBytes, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(cfgBytes, cfg)
	if err != nil {
		fmt.Println(err)
		return
	}
}