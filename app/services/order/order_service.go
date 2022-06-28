package order

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"google.golang.org/grpc"

	"github.com/badfan/inno-taxi-user-service/app/rpc"
)

type IOrderService interface {
	SetDriverRating(ctx context.Context, rating int) error
	GetOrderHistory(ctx context.Context, uuid uuid.UUID) ([]string, error)
	GetTaxiForUser(ctx context.Context, uuid uuid.UUID, rating float32, origin, destination, taxiType string) (string, float32, error)
}

type OrderService struct {
	Client rpc.OrderServiceClient
}

func NewOrderService(conn *grpc.ClientConn) *OrderService {
	return &OrderService{Client: rpc.NewOrderServiceClient(conn)}
}

func (s *OrderService) SetDriverRating(ctx context.Context, rating int) error {
	_, err := s.Client.SetDriverRating(ctx, &rpc.SetDriverRatingRequest{Rating: int32(rating)})
	return err
}

func (s *OrderService) GetOrderHistory(ctx context.Context, uuid uuid.UUID) ([]string, error) {
	ordersResponse, err := s.Client.GetOrderHistory(ctx, &rpc.GetOrderHistoryRequest{Uuid: uuid.String()})
	if err != nil {
		return nil, err
	}

	orderHistory := grpcOrdersConvert(ordersResponse.GetOrders())

	return orderHistory, nil
}

func (s *OrderService) GetTaxiForUser(
	ctx context.Context,
	uuid uuid.UUID,
	rating float32,
	origin, destination, taxiType string) (string, float32, error) {

	driverInfo, err := s.Client.GetTaxiForUser(ctx,
		&rpc.GetTaxiForUserRequest{
			UserUuid:    uuid.String(),
			Origin:      origin,
			Destination: destination,
			TaxiType:    taxiType,
			UserRating:  rating,
		})
	if err != nil {
		return "", 0, err
	}

	return driverInfo.GetDriverUuid(), driverInfo.GetDriverRating(), nil
}

func grpcOrdersConvert(source []*rpc.Order) []string {
	var orders []string
	for _, item := range source {
		orders = append(orders, fmt.Sprintf("User UUID: %s\nDriver UUID: %s\nOrigin: %s\nDestination: %s\n"+
			"Taxi type: %s\nDate: %s\nDuration: %s", item.GetUserUuid(), item.GetDriverUuid(), item.GetOrigin(),
			item.GetDestination(), item.GetTaxiType(), item.GetDate(), item.GetDuration()))
	}

	return orders
}
