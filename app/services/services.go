package services

import (
	"github.com/badfan/inno-taxi-user-service/app/models/sqlc"
	"github.com/badfan/inno-taxi-user-service/app/resources"
	"go.uber.org/zap"
)

type ServiceI interface {
	SignUp(user sqlc.User) (int, error)
	SignIn(phone, password string) (string, error)
}

type Service struct {
	resource resources.ResourceI
	logger   *zap.SugaredLogger
}

func NewService(resource resources.ResourceI, logger *zap.SugaredLogger) *Service {
	return &Service{resource: resource, logger: logger}
}
