package services

import (
	"context"

	"github.com/badfan/inno-taxi-user-service/app/models"
	"github.com/badfan/inno-taxi-user-service/app/resources"
	"go.uber.org/zap"
)

type IService interface {
	SignUp(ctx context.Context, user *models.User) (int, error)
	SignIn(ctx context.Context, phone, password string) (string, error)
	ParseToken(accessToken string) (int, error)
	GetUserRating(ctx context.Context, id int) (float32, error)
}

type Service struct {
	resource resources.IResource
	logger   *zap.SugaredLogger
}

func NewService(resource resources.IResource, logger *zap.SugaredLogger) *Service {
	return &Service{resource: resource, logger: logger}
}
