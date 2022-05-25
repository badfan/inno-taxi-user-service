package handlers

import (
	"github.com/badfan/inno-taxi-user-service/app/services"
	"go.uber.org/zap"
)

type Handler struct {
	service services.IService
	logger  *zap.SugaredLogger
}

func NewHandler(service services.IService, logger *zap.SugaredLogger) *Handler {
	return &Handler{service: service, logger: logger}
}
