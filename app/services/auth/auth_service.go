package auth //nolint:typecheck

import (
	"github.com/badfan/inno-taxi-user-service/app/apperrors"
	"github.com/pkg/errors"

	"github.com/badfan/inno-taxi-user-service/app/services/user"

	"github.com/badfan/inno-taxi-user-service/app/resources"
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
			return nil, errors.Wrap(apperrors.ErrInvalidToken, "invalid signing method")
		}

		return []byte(user.SigningKey), nil
	})
	if err != nil {
		return 0, errors.Wrapf(apperrors.ErrInvalidToken, "error occurred while parsing token: %s", err.Error())
	}

	claims, ok := token.Claims.(*user.TokenClaims)
	if !ok {
		return 0, errors.Wrap(apperrors.ErrInvalidToken, "token claims are not of type *tokenClaims")
	}

	return int(claims.ID), nil
}
