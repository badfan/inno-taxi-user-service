package user //nolint:typecheck

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/badfan/inno-taxi-user-service/app"

	"github.com/badfan/inno-taxi-user-service/app/models"
	"github.com/badfan/inno-taxi-user-service/app/resources"
	"github.com/dgrijalva/jwt-go" //nolint:typecheck
	"go.uber.org/zap"
)

const (
	hashSalt   = "jfkdsfjhs16743v213fdskjf"
	SigningKey = "daf#dfs@23df$321h7931g!4"
)

type IUserService interface {
	SignUp(ctx context.Context, user *models.User) (int, error)
	SignIn(ctx context.Context, phone, password string) (string, error)
	GetUserRating(ctx context.Context, id int) (float32, error)
}

type UserService struct {
	resource  resources.IResource
	apiConfig *app.APIConfig
	logger    *zap.SugaredLogger
}

func NewUserService(resource resources.IResource, apiConfig *app.APIConfig, logger *zap.SugaredLogger) *UserService {
	return &UserService{resource: resource, apiConfig: apiConfig, logger: logger}
}

type TokenClaims struct {
	jwt.StandardClaims       //nolint:typecheck
	ID                 int32 `json:"id"`
}

func (s *UserService) SignUp(ctx context.Context, user *models.User) (int, error) {
	if _, err := s.resource.GetUserIDByPhone(ctx, user.PhoneNumber); err == nil {
		return 0, errors.New("phone number is already taken")
	}

	user.Password = generatePasswordHash(user.Password)

	res, err := s.resource.CreateUser(ctx, user)
	if err != nil {
		return 0, err
	}

	return res, err
}

func (s *UserService) SignIn(ctx context.Context, phone, password string) (string, error) {
	password = generatePasswordHash(password)

	user, err := s.resource.GetUserByPhoneAndPassword(ctx, phone, password)
	if err != nil {
		return "", err
	}

	tokenTTL, _ := strconv.ParseInt(s.apiConfig.TokenTTL, 10, 64)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{ //nolint:typecheck
		StandardClaims: jwt.StandardClaims{ //nolint:typecheck
			ExpiresAt: time.Now().Add(time.Duration(tokenTTL) * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		ID: user.ID,
	})

	return token.SignedString([]byte(SigningKey))
}

func (s *UserService) GetUserRating(ctx context.Context, id int) (float32, error) {
	return s.resource.GetUserRatingByID(ctx, id)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(hashSalt)))
}
