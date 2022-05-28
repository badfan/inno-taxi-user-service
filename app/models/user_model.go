package models

import (
	"time"

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
