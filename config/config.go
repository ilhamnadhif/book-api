package config

import (
	"github.com/BurntSushi/toml"
	"github.com/go-playground/validator/v10"
	"log"
)

func Init(validate *validator.Validate) Config {
	var config Config
	_, err := toml.DecodeFile("config/config.toml", &config)

	if err != nil {
		log.Fatalln(err)
	}

	ok := validateConfig(validate, &config)
	if !ok {
		log.Fatalln("Incomplete configuration file")
	}

	return config
}

func validateConfig(validate *validator.Validate, config *Config) bool {
	if config.Server.HostPort == "" {
		config.Server.HostPort = ":8000"
	}

	err := validate.Struct(*config)
	if err != nil {
		log.Fatalln(err.Error())
		return false
	} else {
		return true
	}

}
