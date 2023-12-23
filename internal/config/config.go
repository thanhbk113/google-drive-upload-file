package config

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

var cfg Config

// Init read env config
func Init() {
	var (
		ctx = context.Background()
	)
	envFile := ".env"
	if err := godotenv.Load(envFile); err != nil {
		fmt.Println("Load env file err: ", err)
	}
	if err := envconfig.Process(ctx, &cfg); err != nil {
		log.Fatal("Assign env err: ", err)
	}
}

// InitWithENVFile ...
func InitWithENVFile(envFile string) {
	var (
		ctx = context.Background()
	)
	if err := godotenv.Load(envFile); err != nil {
		log.Fatal("Load env file err", err)
	}
	if err := envconfig.Process(ctx, &cfg); err != nil {
		log.Fatal("Assign env err: ", err)
	}
}

// GetEnv get env config
func GetEnv() *Config {
	return &cfg
}
