package main

import (
	"fmt"

	"github.com/Netflix/go-env"
)

type Config struct {
	DBName     string `env:"DB_NAME"`
	DBPort     string `env:"DB_PORT"`
	DBHost     string `env:"DB_HOST"`
	DBUsername string `env:"DB_USERNAME"`
	DBPassword string `env:"DB_PASSWORD"`
	DBParams   string `env:"DB_PARAMS"`
}

func main() {
	var config Config

	_, err := env.UnmarshalFromEnviron(&config)
	handleErr(err)
	fmt.Printf("\nConfig: %+v\n", config)

}
func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
