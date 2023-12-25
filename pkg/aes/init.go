package aes

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	gorandom "github.com/zekiahmetbayar/go-random"
)

func init() {
	// Check is app key exists
	if os.Getenv("APP_KEY") != "" {
		return
	}
	// Read .env file
	dotenv, err := godotenv.Read("./.env")
	if err != nil {
		log.Println("Can not found .env file")
	}
	// Create random string
	random, err := gorandom.String(true, false, false, 32)
	if err != nil {
		log.Printf("error when creating random string, got err %s", err.Error())
	}
	dotenv["APP_KEY"] = random
	// Write app key to .env
	if err := godotenv.Write(dotenv, "./.env"); err != nil {
		log.Fatalln("Can not write to .env file")
	}
	// Set app key
	os.Setenv("APP_KEY", random)
}
