package middleware

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

func LoadEnvVariables() {

	dir, err := os.Getwd()

	err = godotenv.Load(filepath.Join(dir, ".env.example"))

	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("Error loading .env file!")
	}
}
