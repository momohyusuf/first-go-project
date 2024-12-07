package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error converting code to byte %v", payload)
		w.WriteHeader(500)
		w.Write([]byte("Sorry an error occurred"))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)

}

func responseWithError(w http.ResponseWriter, code int, msg string) {

	if code > 499 {
		fmt.Println("Error occur")

	}

	type errorResponse struct {
		Error string `jason:"Error"`
	}
	responseWithJson(w, code, errorResponse{Error: msg})

}
