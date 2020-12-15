package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Name     string `yaml:"dbName"`
	} `yaml:"database"`
	RapidAPI struct {
		Host string `yaml:"host"`
		Key  string `yaml:"key"`
	} `yaml:"rapidApi"`
}

func (c *Config) parse(data []byte) error {
	return yaml.Unmarshal(data, c)
}

func Read(configFile string) (Config, error) {
	config := Config{}

	if len(configFile) > 0 {
		if _, err := os.Stat(configFile); err != nil {
			return config, err
		}

		data, err := ioutil.ReadFile(configFile)
		if err != nil {
			return config, err
		}

		if err := config.parse(data); err != nil {
			return config, err
		}
	}

	return config, nil
}
