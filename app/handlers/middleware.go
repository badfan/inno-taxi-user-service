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
		h.ErrorLogger(c, errors.Wrap(apperrors.ErrInvalidToken, "empty authorization header"))
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		h.ErrorLogger(c, errors.Wrap(apperrors.ErrInvalidToken, "invalid authorization header"))
		return
	}

	id, err := h.authService.ParseToken(headerParts[1])
	if err != nil {
		h.ErrorLogger(c, err)
		return
	}

	c.Set("userID", id)
}
