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
