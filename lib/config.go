package lib

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DBConfig string `json:"db_config"`
	Port     int    `json:"port"`
}

func ParseConfig() *Config {
	file, err := os.Open("config/development/config.json")
	if err != nil {
		panic(fmt.Sprintf("failed to get config file, err: %s", err.Error()))
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	cfg := &Config{}
	if err := decoder.Decode(cfg); err != nil {
		panic(fmt.Sprintf("failed to parse config, err: %s", err.Error()))
	}
	return cfg
}
