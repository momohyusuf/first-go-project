package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/rss/internal/database"
	"github.com/rss/utils"
)

func (dbConfig *apiConfig) handleUser(w http.ResponseWriter, r *http.Request) {
	// since this is post route, it means it is expecting a json body from the frontend
	// lets parse the json body
	// A struct to rep what we are expecting
	type parameters struct {
		Name         string `json:"name"`
		Email        string `json:"email"`
		Phone_number string `json:"phone_number"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	fmt.Print(err)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Sorry an error occur Provide the following fields are missing: %v", err))
	}

	// create the user in your data base
	user, err := dbConfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:          uuid.New(),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Name:        params.Name,
		Email:       params.Email,
		ApiKey:      utils.GenerateApiKey(),
		PhoneNumber: params.Phone_number,
	})

	if err != nil {
		fmt.Println(err)
		responseWithError(w, 400, fmt.Sprintf("Sorry an error occur %v", err))
		return
	}
	responseWithJson(w, 201, updatedUserModelFieldToCustom(user))
}
