package config

import (
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

func NewConfig() (*koanf.Koanf, error) {
	k := koanf.New(".")
	err := k.Load(file.Provider("config/config.yaml"), yaml.Parser())
	if err != nil {
		return nil, err
	}
	return k, nil
}
