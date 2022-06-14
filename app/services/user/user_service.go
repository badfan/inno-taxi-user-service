package user

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"time"

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
	GetOrderHistory(ctx context.Context, id int) ([]*pb.Order, error)
	GetTaxi(ctx context.Context, id int, origin, destination, taxiType string) (*pb.GetTaxiForUserResponse, error)
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
	return s.resource.GetUserRatingByID(ctx, id)
}

func (s *UserService) SetDriverRating(ctx context.Context, rating int) error {
	_, err := s.grpcClient.SetDriverRating(ctx, &proto.SetDriverRatingRequest{Rating: int32(rating)})
	return err
}

func (s *UserService) GetOrderHistory(ctx context.Context, id int) ([]*pb.Order, error) {
	uuid, err := s.resource.GetUserUUIDByID(ctx, id)
	if err != nil {
		return nil, err
	}

	stream, err := s.grpcClient.GetOrderHistory(ctx, &proto.GetOrderHistoryRequest{Uuid: uuid.String()})
	if err != nil {
		return nil, err
	}

	var orderHistory []*pb.Order
	for {
		order, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		orderHistory = append(orderHistory, order)
	}

	return orderHistory, nil
}

func (s *UserService) GetTaxi(ctx context.Context, id int, origin, destination, taxiType string) (*pb.GetTaxiForUserResponse, error) {
	userUUID, rating, err := s.resource.GetUserUUIDAndRatingByID(ctx, id)
	if err != nil {
		return nil, err
	}

	driverInfo, err := s.grpcClient.GetTaxiForUser(ctx, &pb.GetTaxiForUserRequest{
		UserUuid:    userUUID.String(),
		UserRating:  rating,
		Origin:      origin,
		Destination: destination,
		TaxiType:    taxiType,
	})
	if err != nil {
		return nil, err
	}

	return driverInfo, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(hashSalt)))
}
