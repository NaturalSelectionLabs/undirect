package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Debug  bool   `yaml:"debug"`
	Listen string `yaml:"listen"`
}

var GlobalConfig Config

func LoadConfig(filename string) error {

	data, err := os.ReadFile(filename)

	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, &GlobalConfig); err != nil {
		return err
	}

	return nil

}
