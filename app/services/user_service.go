package services

import (
	"errors"
	"github.com/badfan/inno-taxi-user-service/app/models/sqlc"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	signingKey = "daf#dfs@23df$321h7931g!4"
)

type tokenClaims struct {
	jwt.StandardClaims
	ID int32 `json:"id"`
}

func (s *Service) SignUp(user sqlc.User) (int, error) {
	if _, err := s.resource.GetUserIDByPhone(user.PhoneNumber); err == nil {
		return 0, errors.New("phone number is already taken")
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
			ExpiresAt: time.Now().Add(60 * time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *Service) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return int(claims.ID), nil
}

func (s *Service) GetUserRating(id int) (float32, error) {
	return s.resource.GetUserRatingByID(id)
}
