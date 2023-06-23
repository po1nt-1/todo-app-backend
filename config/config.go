package config

import (
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

var (
	k      = koanf.New(".")
	parser = yaml.Parser()
)

type Config struct {
	App struct {
		Name string `koanf:"name"`
	} `koanf:"app"`
	HTTP struct {
		Host string `koanf:"host"`
		Port string `koanf:"port"`
	} `koanf:"http"`
}

func NewConfig() (*Config, error) {
	err := k.Load(file.Provider("config/config.yaml"), parser)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	// k.Unmarshal("", &cfg)
	err = k.UnmarshalWithConf("", &cfg, koanf.UnmarshalConf{Tag: "koanf"})
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
