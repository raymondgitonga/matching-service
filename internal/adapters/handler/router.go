package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal("Healthy")
	if err != nil {
		fmt.Printf("error writing marshalling response: %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(response)
	if err != nil {
		fmt.Printf("error writing http response: %s", err)
	}
}
