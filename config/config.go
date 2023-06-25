package config

import (
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
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
	Config struct {
		Path string `koanf:"path"`
	} `koanf:"config"`
}

func prettyEnv(s string) string {
	s = strings.TrimPrefix(s, "TODO_")
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, "_", ".")
	return s
}

func NewConfig() (*Config, error) {
	err := k.Load(structs.Provider(Config{
		App: struct {
			Name string `koanf:"name"`
		}{"todo-app"},
		HTTP: struct {
			Host string `koanf:"host"`
			Port string `koanf:"port"`
		}{"0.0.0.0", "8080"},
		Config: struct {
			Path string `koanf:"path"`
		}{"config/config.yaml"},
	}, "koanf"), nil)
	if err != nil {
		return nil, err
	}

	err = k.Load(env.Provider("TODO_", ".", prettyEnv), nil)
	if err != nil {
		return nil, err
	}

	if k.String("config.path") != "config/config.yaml" {
		err = k.Load(file.Provider(k.String("config.path")), parser)
		if err != nil {
			return nil, err
		}
	}

	cfg := &Config{}
	err = k.Unmarshal("", &cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
