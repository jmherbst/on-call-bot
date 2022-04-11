package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Mappings       map[string]string `yaml:"mappings"`
	PagerDutyToken string
	SlackToken     string
}

func GetConfig() (Config, error) {
	var config Config

	filename, _ := filepath.Abs("config.yml")

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return config, fmt.Errorf("GetConfig: unable to read config file: %w", err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return config, fmt.Errorf("GetConfig: unable to unmarshal config file: %w", err)
	}

	config.PagerDutyToken = os.Getenv("PAGERDUTY_TOKEN")
	config.SlackToken = os.Getenv("SLACK_TOKEN")

	return config, nil
}
