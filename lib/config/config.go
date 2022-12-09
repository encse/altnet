package config

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Twitter struct {
		AccessToken string `yaml:"accessToken" envconfig:"TWITTER_ACCESS_TOKEN"`
	} `yaml:"twitter"`
}

func Read(file string) (Config, error) {
	var cfg Config

	f, err := os.Open(file)
	if err != nil {
		return cfg, err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)

	if err != nil {
		return cfg, err
	}

	err = envconfig.Process("", &cfg)
	return cfg, err
}
