package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
		//os.Exit(1)
	}
}

func GetEnv(key string) string {
	loadEnv()
	return os.Getenv(key)
}
