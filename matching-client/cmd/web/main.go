package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("error loading configs: %s", err)
		return
	}

	appConfig := NewAppConfigs(os.Getenv("BASE_URL"), os.Getenv("MATCHING_URL"))

	router, err := appConfig.StartApp()

	if err != nil {
		fmt.Printf("error starting app: %s", err)
		return
	}

	server := &http.Server{
		Addr:              ":8081",
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           router,
	}

	fmt.Println("starting server on :8081")
	err = server.ListenAndServe()
	if err != nil {
		fmt.Printf("error starting server: %s", err)
	}

}
