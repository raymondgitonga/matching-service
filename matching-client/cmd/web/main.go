package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	appConfig := NewAppConfigs("/match")

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
	err = server.ListenAndServe()
	if err != nil {
		fmt.Printf("error starting server: %s", err)
	}

}
