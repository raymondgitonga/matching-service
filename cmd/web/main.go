package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file: %s", err)
	}

	appConfigs := NewAppConfigs(
		os.Getenv("DB_CONNECTION_URL"),
		os.Getenv("DB_NAME"),
		os.Getenv("BASE_URL"),
	)
	server, err := appConfigs.StartApp()

	if err != nil {
		log.Fatalf("error starting app: %s", err)
	}

	err = http.ListenAndServe(":8080", server)
	if err != nil {
		log.Fatalf("error starting server: %s", err)
	}
}
