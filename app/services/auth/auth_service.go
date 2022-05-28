package auth //nolint:typecheck

import (
	"errors"

	"github.com/badfan/inno-taxi-user-service/app/resources"
	"github.com/badfan/inno-taxi-user-service/app/services/user"
	"github.com/dgrijalva/jwt-go" //nolint:typecheck
	"go.uber.org/zap"
)

type IAuthenticationService interface {
	ParseToken(accessToken string) (int, error)
}

type AuthenticationService struct {
	resource resources.IResource
	logger   *zap.SugaredLogger
}

func NewAuthenticationService(resource resources.IResource, logger *zap.SugaredLogger) *AuthenticationService {
	return &AuthenticationService{resource: resource, logger: logger}
}

func (s *AuthenticationService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &user.TokenClaims{}, func(token *jwt.Token) (interface{}, error) { //nolint:typecheck
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(user.SigningKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*user.TokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return int(claims.ID), nil
}
