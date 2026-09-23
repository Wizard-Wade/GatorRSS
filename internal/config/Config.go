package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	Current_user_name string `json:"user_name"`
	Url               string `json:"db_url"`
}

func (c *Config) SetUser(user_name string) error {
	c.Current_user_name = user_name
	err := write(c)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func Read() (Config, error) {
	c := Config{}
	path, err := getConfigFilepath()
	if err != nil {
		return c, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return c, err
	}

	err = json.Unmarshal(data, &c)
	if err != nil {
		return Config{}, err
	}

	return c, nil
}

func write(cfg *Config) error {
	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	var path string
	path, err = getConfigFilepath()
	if err != nil {
		return err
	}
	err = os.WriteFile(path, jsonData, 0644)
	return nil
}

func getConfigFilepath() (string, error) {
	dir, ok := os.UserHomeDir()
	if ok != nil {
		return "", fmt.Errorf("cound not find user home directory")
	}

	return dir + "//" + configFileName, nil
}
