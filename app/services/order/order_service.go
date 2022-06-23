package order

import (
	"fmt"

	pb "github.com/badfan/inno-taxi-user-service/app/rpc"
	"google.golang.org/grpc"
)

type OrderService struct {
	pb.OrderServiceClient
}

func NewOrderService(conn *grpc.ClientConn) *OrderService {
	return &OrderService{pb.NewOrderServiceClient(conn)}
}

func GRPCOrdersConvert(source []*pb.Order) []string {
	var orders []string
	for _, item := range source {
		orders = append(orders, fmt.Sprintf("User UUID: %s\nDriver UUID: %s\nOrigin: %s\nDestination: %s\nTaxi type: %s\nDate: %s\n"+
			"Duration: %s", item.GetUserUuid(), item.GetDriverUuid(), item.GetOrigin(), item.GetDestination(),
			item.GetTaxiType(), item.GetDate(), item.GetDuration()))
	}

	return orders
}

func GetOrderHistoryRequestConvert(source string) *pb.GetOrderHistoryRequest {
	return &pb.GetOrderHistoryRequest{Uuid: source}
}

func SetDriverRatingRequestConvert(source int32) *pb.SetDriverRatingRequest {
	return &pb.SetDriverRatingRequest{Rating: source}
}

func GetTaxiRequestConvert(rating float32, userUUID, origin, destination, taxiType string) *pb.GetTaxiForUserRequest {
	return &pb.GetTaxiForUserRequest{
		UserUuid:    userUUID,
		UserRating:  rating,
		Origin:      origin,
		Destination: destination,
		TaxiType:    taxiType,
	}
}
