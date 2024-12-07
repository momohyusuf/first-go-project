package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/rss/internal/database"
)

func (dbConfig *apiConfig) handleCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	// since this is post route, it means it is expecting a json body from the frontend
	// lets parse the json body
	// A struct to rep what we are expecting
	type parameters struct {
		Url  string `json:"url"`
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Sorry an error occur Provide the following fields are missing: %v", err))
	}

	// check if email address is valid
	if params.Url == "" {
		responseWithError(w, 400, fmt.Sprintf("Please provide your feed url: %v", params.Url))
		return
	}

	// create the user in your data base
	feed, err := dbConfig.DB.CreateNewFeed(r.Context(), database.CreateNewFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		UserID:    user.ID,
		Url:       params.Url,
	})

	if err != nil {
		fmt.Println(err)
		responseWithError(w, 400, fmt.Sprintf("Sorry an error occur %v", err))
		return
	}
	responseWithJson(w, 201, updatedFeedModelFieldToCustom(feed))
}
