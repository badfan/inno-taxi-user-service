package handlers

import (
	"strings"

	"github.com/badfan/inno-taxi-user-service/app/apperrors"
	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
)

func (h *Handler) Middleware(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		apperrors.NewErrorResponse(c, errors.Wrap(apperrors.ErrInvalidToken, "empty authorization header"), h.logger)
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		apperrors.NewErrorResponse(c, errors.Wrap(apperrors.ErrInvalidToken, "invalid authorization header"), h.logger)
		return
	}

	id, err := h.authService.ParseToken(headerParts[1])
	if err != nil {
		apperrors.NewErrorResponse(c, err, h.logger)
		return
	}

	c.Set("userID", id)
}
