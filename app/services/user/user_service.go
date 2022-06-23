package user

import (
	"context"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"time"

	pb "github.com/badfan/inno-taxi-user-service/app/rpc"
	"github.com/badfan/inno-taxi-user-service/app/services/order"

	"github.com/badfan/inno-taxi-user-service/app/apperrors"
	"github.com/pkg/errors"

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
	SetDriverRating(ctx context.Context, rating int) error
	GetOrderHistory(ctx context.Context, id int) ([]string, error)
	FindTaxi(ctx context.Context, id int, origin, destination, taxiType string) (string, float32, error)
}

type UserService struct {
	resource     resources.IResource
	orderService *order.OrderService
	apiConfig    *app.APIConfig
	logger       *zap.SugaredLogger
}

func NewUserService(resource resources.IResource, orderService *order.OrderService, apiConfig *app.APIConfig, logger *zap.SugaredLogger) *UserService {
	return &UserService{resource: resource, orderService: orderService, apiConfig: apiConfig, logger: logger}
}

type TokenClaims struct {
	jwt.StandardClaims       //nolint:typecheck
	ID                 int32 `json:"id"`
}

func (s *UserService) SignUp(ctx context.Context, user *models.User) (int, error) {
	if _, err := s.resource.GetUserIDByPhone(ctx, user.PhoneNumber); err == nil {
		return 0, errors.Wrapf(apperrors.ErrPhoneNumberIsAlreadyTaken, "error occurred while verifying phone number: %s", err.Error())
	}

	user.Password = generatePasswordHash(user.Password)

	res, err := s.resource.CreateUser(ctx, user)
	if err != nil {
		return 0, errors.Wrapf(err, "error occurred while creating user")
	}

	return res, err
}

func (s *UserService) SignIn(ctx context.Context, phone, password string) (string, error) {
	password = generatePasswordHash(password)

	user, err := s.resource.GetUserByPhoneAndPassword(ctx, phone, password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.Wrapf(apperrors.ErrUserNotFound,
				"error occurred while verifying phone number and password: %s", err.Error())
		}
		return "", errors.Wrap(err, "error occurred while verifying phone number and password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{ //nolint:typecheck
		StandardClaims: jwt.StandardClaims{ //nolint:typecheck
			ExpiresAt: time.Now().Add(time.Duration(s.apiConfig.TokenTTL) * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		ID: user.ID,
	})

	return token.SignedString([]byte(SigningKey))
}

func (s *UserService) GetUserRating(ctx context.Context, id int) (float32, error) {
	user, err := s.resource.GetUserRatingByID(ctx, id)
	if err != nil {
		return 0, errors.Wrap(err, "error occurred while getting user rating")
	}

	return user, nil
}

func (s *UserService) SetDriverRating(ctx context.Context, rating int) error {
	_, err := s.orderService.SetDriverRating(ctx, &pb.SetDriverRatingRequest{Rating: int32(rating)})
	if err != nil {
		return errors.Wrap(err, "error occurred while setting driver rating")
	}

	return nil
}

func (s *UserService) GetOrderHistory(ctx context.Context, id int) ([]string, error) {
	uuid, err := s.resource.GetUserUUIDByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "error occurred while getting user UUID")
	}

	ordersResponse, err := s.orderService.GetOrderHistory(ctx, &pb.GetOrderHistoryRequest{Uuid: uuid.String()})
	if err != nil {
		return nil, errors.Wrap(err, "error occurred while getting orders history from grpc server")
	}

	orderHistory := order.GRPCOrdersConvert(ordersResponse.Orders)

	return orderHistory, nil
}

func (s *UserService) FindTaxi(ctx context.Context, id int, origin, destination, taxiType string) (string, float32, error) {
	userUUID, rating, err := s.resource.GetUserUUIDAndRatingByID(ctx, id)
	if err != nil {
		return "", 0, errors.Wrap(err, "error occurred while getting user UUID and rating")
	}

	driverInfo, err := s.orderService.GetTaxiForUser(ctx, &pb.GetTaxiForUserRequest{
		UserUuid:    userUUID.String(),
		UserRating:  rating,
		Origin:      origin,
		Destination: destination,
		TaxiType:    taxiType,
	})
	if err != nil {
		return "", 0, errors.Wrap(err, "error occurred while finding taxi for user")
	}

	return driverInfo.GetDriverUuid(), driverInfo.GetDriverRating(), nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(hashSalt)))
}
