package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/rss/internal/database"
)

type User struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Api_key     string    `json:"api_key"`
}

func updatedUserModelFieldToCustom(dbUser database.User) User {
	user := User{
		ID:          dbUser.ID,
		CreatedAt:   dbUser.CreatedAt,
		UpdatedAt:   dbUser.UpdatedAt,
		Name:        dbUser.Name,
		Email:       dbUser.Email,
		Api_key:     dbUser.ApiKey,
		PhoneNumber: dbUser.PhoneNumber,
	}
	return user
}
