package config

import (
	"os"

	"github.com/encse/altnet/lib/fs"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Twitter struct {
		AccessToken string `yaml:"accessToken" envconfig:"TWITTER_ACCESS_TOKEN"`
	} `yaml:"twitter"`
	Dfrotz struct {
		Location string `yaml:"location" envconfig:"DFROTZ"`
	} `yaml:"dfrotz"`
	Memcached struct {
		Server string `yaml:"server"`
	} `yaml:"memcached"`
}

func Get() Config {
	var cfg Config

	f, err := os.Open(fs.WithAppRoot("config.yml"))
	if err != nil {
		return cfg
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)

	if err != nil {
		return cfg
	}

	err = envconfig.Process("", &cfg)
	return cfg
}
