package config

import (
	"strings"

	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Config struct {
}

func LoadConfig(path string) (*Config, error) {
	k := koanf.New(".")
	if err := k.Load(file.Provider(path), json.Parser()); err != nil {
		return nil, err
	}

	k.Load(env.Provider("BASD_", ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, "BASD_")), "_", ".", -1)
	}), nil)

	var c Config
	if err := k.Unmarshal("", &c); err != nil {
		return nil, err
	}

	return &c, nil
}
