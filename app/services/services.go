package services

import (
	"github.com/badfan/inno-taxi-user-service/app/models/sqlc"
	"github.com/badfan/inno-taxi-user-service/app/resources"
	"go.uber.org/zap"
)

type IService interface {
	SignUp(user sqlc.User) (int, error)
	SignIn(phone, password string) (string, error)
	ParseToken(accessToken string) (int, error)
	GetUserRating(id int) (float32, error)
}

type Service struct {
	resource resources.IResource
	logger   *zap.SugaredLogger
}

func NewService(resource resources.IResource, logger *zap.SugaredLogger) *Service {
	return &Service{resource: resource, logger: logger}
}
