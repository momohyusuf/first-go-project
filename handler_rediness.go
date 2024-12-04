package main

import "net/http"

type Payload struct {
	Message string `json:"message"`
}

func handleServerReadiness(w http.ResponseWriter, r *http.Request) {
	responseObj := Payload{
		Message: "Hello and welcome",
	}
	responseWithJson(w, 200, responseObj)

}
