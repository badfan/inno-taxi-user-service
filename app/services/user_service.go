package services

import (
	"fmt"
	"github.com/badfan/inno-taxi-user-service/app/models/sqlc"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	signingKey = "daf#dfs@23df$321h7931g!4"
)

type tokenClaims struct {
	jwt.StandardClaims
	id int `json:"id"`
}

func (s *Service) SignUp(user sqlc.User) (int, error) {
	if _, err := s.resource.GetUserIDByPhone(user.PhoneNumber); err == nil {
		return 0, fmt.Errorf("phone number is already taken")
	}

	res, err := s.resource.CreateUser(user)
	if err != nil {
		return 0, err
	}

	return res, err
}

func (s *Service) SignIn(phone, password string) (string, error) {
	user, err := s.resource.GetUserByPhoneAndPassword(phone, password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		int(user.ID),
	})

	return token.SignedString([]byte(signingKey))
}
