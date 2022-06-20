package user

import (
	"context"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"io"
	"time"

	"github.com/badfan/inno-taxi-user-service/app/apperrors"
	"github.com/pkg/errors"

	"github.com/badfan/inno-taxi-user-service/app"
	"github.com/badfan/inno-taxi-user-service/app/models"
	"github.com/badfan/inno-taxi-user-service/app/resources"
	"github.com/badfan/inno-taxi-user-service/app/services/proto"
	pb "github.com/badfan/inno-taxi-user-service/app/services/proto"
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
	resource   resources.IResource
	grpcClient proto.OrderServiceClient
	apiConfig  *app.APIConfig
	logger     *zap.SugaredLogger
}

func NewUserService(resource resources.IResource, grpcClient proto.OrderServiceClient, apiConfig *app.APIConfig, logger *zap.SugaredLogger) *UserService {
	return &UserService{resource: resource, grpcClient: grpcClient, apiConfig: apiConfig, logger: logger}
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
		return 0, errors.Wrapf(apperrors.ErrInternalServer, "error occurred while creating user: %s", err.Error())
	}

	return res, err
}

func (s *UserService) SignIn(ctx context.Context, phone, password string) (string, error) {
	password = generatePasswordHash(password)

	user, err := s.resource.GetUserByPhoneAndPassword(ctx, phone, password)
	if err != nil {
		var apperr error
		switch err {
		case sql.ErrNoRows:
			apperr = apperrors.ErrNotFound
		default:
			apperr = apperrors.ErrInternalServer
		}
		return "", errors.Wrapf(apperr,
			"error occurred while verifying phone number and password: %s", err.Error())
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
		return 0, errors.Wrapf(apperrors.ErrInternalServer, "error occurred while getting user rating: %s", err.Error())
	}
	return user, nil
}

func (s *UserService) SetDriverRating(ctx context.Context, rating int) error {
	_, err := s.grpcClient.SetDriverRating(ctx, &proto.SetDriverRatingRequest{Rating: int32(rating)})
	if err != nil {
		return errors.Wrapf(apperrors.ErrInternalServer, "error occurred while setting driver rating: %s", err.Error())
	}
	return nil
}

func (s *UserService) GetOrderHistory(ctx context.Context, id int) ([]string, error) {
	uuid, err := s.resource.GetUserUUIDByID(ctx, id)
	if err != nil {
		return nil, errors.Wrapf(apperrors.ErrInternalServer, "error occurred while getting user UUID: %s", err.Error())
	}

	stream, err := s.grpcClient.GetOrderHistory(ctx, &proto.GetOrderHistoryRequest{Uuid: uuid.String()})
	if err != nil {
		return nil, errors.Wrapf(apperrors.ErrInternalServer, "error occurred while getting grpc stream: %s", err.Error())

	}

	var orderHistory []string
	for {
		order, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, errors.Wrapf(apperrors.ErrInternalServer, "error occurred while getting orders from stream: %s", err.Error())
		}
		convertedOrder := grpcOrderConvert(order)
		orderHistory = append(orderHistory, convertedOrder)
	}

	return orderHistory, nil
}

func (s *UserService) FindTaxi(ctx context.Context, id int, origin, destination, taxiType string) (string, float32, error) {
	userUUID, rating, err := s.resource.GetUserUUIDAndRatingByID(ctx, id)
	if err != nil {
		return "", 0, errors.Wrapf(apperrors.ErrInternalServer, "error occurred while getting user UUID and rating: %s", err.Error())
	}

	driverInfo, err := s.grpcClient.GetTaxiForUser(ctx, &pb.GetTaxiForUserRequest{
		UserUuid:    userUUID.String(),
		UserRating:  rating,
		Origin:      origin,
		Destination: destination,
		TaxiType:    taxiType,
	})
	if err != nil {
		return "", 0, errors.Wrapf(apperrors.ErrInternalServer, "error occurred while finding taxi: %s", err.Error())
	}

	return driverInfo.GetDriverUuid(), driverInfo.GetDriverRating(), nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(hashSalt)))
}

func grpcOrderConvert(source *pb.Order) string {
	return fmt.Sprintf("User UUID: %s\nDriver UUID: %s\nOrigin: %s\nDestination: %s\nTaxi type: %s\nDate: %s\n"+
		"Duration: %s", source.GetUserUuid(), source.GetDriverUuid(), source.GetOrigin(), source.GetDestination(),
		source.GetTaxiType(), source.GetDate(), source.GetDuration())
}
