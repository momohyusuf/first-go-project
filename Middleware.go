package main

import (
	"fmt"
	"net/http"

	internal "github.com/rss/internal/auth"
	"github.com/rss/internal/database"
)

// Authentication middleware
type authHandler func(http.ResponseWriter, *http.Request, database.User)

// creating the middlewareAuth as a method on apiConfig in other to be able to interact with the database
func (dbConfig *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := internal.GetApiKey(r.Header)

		if err != nil {
			responseWithError(w, 400, fmt.Sprintf("An error occurred: %v", err))
			return
		}

		// find the user with the apikey in out database
		user, err := dbConfig.DB.FindUserByApiKey(r.Context(), apiKey)

		if err != nil {
			responseWithError(w, 400, fmt.Sprintf("An error occurred: %v", err))
			return
		}
		// after all the validations are done move on to the next function
		handler(w, r, user)
	}

}
