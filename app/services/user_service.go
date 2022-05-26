package services

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/badfan/inno-taxi-user-service/app/models"
	"github.com/dgrijalva/jwt-go"
)

const (
	hashSalt   = "jfkdsfjhs16743v213fdskjf"
	signingKey = "daf#dfs@23df$321h7931g!4"
)

type tokenClaims struct {
	jwt.StandardClaims
	ID int32 `json:"id"`
}

func (s *Service) SignUp(ctx context.Context, user *models.User) (int, error) {
	if _, err := s.resource.GetUserIDByPhone(ctx, user.PhoneNumber); err == nil {
		return 0, errors.New("phone number is already taken")
	}

	user.Password = generatePasswordHash(user.Password)

	res, err := s.resource.CreateUser(ctx, user)
	if err != nil {
		return 0, err
	}

	return res, err
}

func (s *Service) SignIn(ctx context.Context, phone, password string) (string, error) {
	user, err := s.resource.GetUserByPhoneAndPassword(ctx, phone, password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(60 * time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *Service) GetUserRating(ctx context.Context, id int) (float32, error) {
	return s.resource.GetUserRatingByID(ctx, id)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(hashSalt)))
}
