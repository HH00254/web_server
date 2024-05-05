package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respongWithError(w http.ResponseWriter, statusCode int, msg string) {
	if statusCode > 499 {
		log.Println("Responding with 5XX error:", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	responWithJSON(w, statusCode, errResponse{
		Error: msg,
	})
}

func responWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	byteData, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)

	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(byteData)
}
