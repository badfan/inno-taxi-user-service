package models

import (
	"time"

	"github.com/badfan/inno-taxi-user-service/app/models/sqlc"
	"github.com/google/uuid"
)

type User struct {
	ID          int32     `json:"id"`
	UserUuid    uuid.UUID `json:"user_uuid"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	UserRating  float32   `json:"user_rating"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func SqlcUserConvert(source *sqlc.User) *User {
	res := &User{
		ID:          source.ID,
		UserUuid:    source.UserUuid,
		Name:        source.Name,
		PhoneNumber: source.PhoneNumber,
		Email:       source.Email,
		Password:    source.Password,
		UserRating:  source.UserRating,
		CreatedAt:   source.CreatedAt,
		UpdatedAt:   source.UpdatedAt,
	}

	return res
}
