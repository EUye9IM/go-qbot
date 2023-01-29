package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type ConfigConnect struct {
	Host          string
	Port          uint
	Qq            string
	Verify_key    string
	Retry_seconds uint
	Max_retry     uint
}
type Config struct {
	Connect ConfigConnect
}

var config Config

func Conf() Config {
	return config
}

func init() {
	// read file
	config_bytes, err := os.ReadFile("setting.yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(config_bytes, &config)
	if err != nil {
		panic(err)
	}
}
