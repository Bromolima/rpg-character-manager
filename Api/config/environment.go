package config

import (
	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
)

type Environment struct {
	ConnectionString string `env:"CONNECTION_STRING"`
}

var Envi Environment

func LoadEnvironments() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	_, err = env.UnmarshalFromEnviron(&Envi)
	if err != nil {
		panic(err)
	}

}
