package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file: %s", err)
	}

	appConfigs, err := NewAppConfigs(
		os.Getenv("DB_CONNECTION_URL"),
		os.Getenv("DB_NAME"),
		os.Getenv("BASE_URL"),
	)

	if err != nil {
		log.Fatalf("error getting app configs: %s", appConfigs)
	}
	router, err := appConfigs.StartApp()

	if err != nil {
		log.Fatalf("error starting app: %s", err)
	}

	server := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           router,
	}

	fmt.Println("starting server on :8080")
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("error starting server: %s", err)
	}
}
