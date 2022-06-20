package handlers

import (
	"github.com/badfan/inno-taxi-user-service/app/apperrors"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (h *Handler) ErrorLogger(c *gin.Context, err error) {
	h.logger.Error(err.Error())
	unwrappedErr := errors.Unwrap(err)

	apperrors.ErrorResponse(c, unwrappedErr)
}
