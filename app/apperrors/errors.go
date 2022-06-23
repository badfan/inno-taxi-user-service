package apperrors

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrUserNotFound              = errors.New("user not found")
	ErrPhoneNumberIsAlreadyTaken = errors.New("phone number is already taken")
	ErrInvalidToken              = errors.New("invalid auth token")
)

type ErrorMessage struct {
	Message string
}

func ErrorResponse(c *gin.Context, err error) {
	var statusCode int

	switch err {
	case ErrPhoneNumberIsAlreadyTaken:
		statusCode = http.StatusBadRequest
	case ErrUserNotFound:
		statusCode = http.StatusNotFound
	case ErrInvalidToken:
		statusCode = http.StatusUnauthorized
	default:
		statusCode = http.StatusInternalServerError
	}

	c.AbortWithStatusJSON(statusCode, ErrorMessage{Message: err.Error()})
}
