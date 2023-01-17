package helpers

import (
	"log"

	"github.com/joho/godotenv"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
		return
	}
}

func LoadEnv(args ...string) {

	envFile := ".env"

	if len(args) > 0 {
		envFile = args[0]
	}

	err := godotenv.Load(envFile)

	if err != nil {
		log.Fatal("Error loading .env file", err)
		return
	}
	log.Println(".env file loaded")

}
