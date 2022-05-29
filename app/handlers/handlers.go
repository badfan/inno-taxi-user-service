package handlers

import (
	"github.com/badfan/inno-taxi-user-service/app/services/auth"
	"github.com/badfan/inno-taxi-user-service/app/services/user"
	"go.uber.org/zap"
)

type Handler struct {
	authService auth.IAuthenticationService
	userService user.IUserService
	logger      *zap.SugaredLogger
}

func NewHandler(authService auth.IAuthenticationService, userService user.IUserService, logger *zap.SugaredLogger) *Handler {
	return &Handler{authService: authService, userService: userService, logger: logger}
}
