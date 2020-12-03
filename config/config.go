package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/yrisob/database_migrations/utils"
)

type Config struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	SslMode  string `json:"sslmode"`
	Sources  string `json:"sources"`
}

const configFilePath string = "./database_migration.json"
const dbDriver string = "postgres"

var config *Config

func GetConfig() (*Config, error) {
	config = &Config{}
	if utils.FileExists(configFilePath) {
		buffer, err := ioutil.ReadFile(configFilePath)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(buffer, config)
		if err != nil {
			config = nil
			return nil, err
		}
		return config, nil
	}
	return nil, errors.New("config file doesn't exist")
}

func (c *Config) GetConnectionString() string {
	return fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=%s", dbDriver, c.User, c.Password, c.Host, c.Port, c.Database, c.SslMode)
}
