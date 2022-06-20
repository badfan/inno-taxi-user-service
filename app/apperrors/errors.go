package apperrors

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	ErrNotFound                  = errors.New("user not found")
	ErrPhoneNumberIsAlreadyTaken = errors.New("phone number is already taken")
	ErrInternalServer            = errors.New("internal server error")
	ErrInvalidToken              = errors.New("invalid auth token")
	ErrBadRequest                = errors.New("incorrect input")
)

type ErrorMessage struct {
	Message string
}

func NewErrorResponse(c *gin.Context, err error, logger *zap.SugaredLogger) {
	var statusCode int

	logger.Error(err.Error())
	unwrappedErr := errors.Unwrap(err)

	switch unwrappedErr {
	case ErrPhoneNumberIsAlreadyTaken:
		statusCode = http.StatusBadRequest
	case ErrNotFound:
		statusCode = http.StatusNotFound
	case ErrInvalidToken:
		statusCode = http.StatusUnauthorized
	case ErrBadRequest:
		statusCode = http.StatusBadRequest
	case ErrInternalServer:
		statusCode = http.StatusInternalServerError
	}

	c.AbortWithStatusJSON(statusCode, ErrorMessage{Message: unwrappedErr.Error()})
}
