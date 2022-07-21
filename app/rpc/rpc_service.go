package rpc

import (
	"context"
	"fmt"
)

type IRPCService interface {
	SetUserRating(ctx context.Context, req *SetUserRatingRequest) (*EmptyResponse, error)
}

type RPCService struct {
	UserServiceServer UnimplementedUserServiceServer
}

func NewRPCService() *RPCService {
	return &RPCService{UserServiceServer: UnimplementedUserServiceServer{}}
}

// TODO : implement
func (s *RPCService) SetUserRating(ctx context.Context, req *SetUserRatingRequest) (*EmptyResponse, error) {
	rating := req.GetRating()

	fmt.Println(rating)

	return &EmptyResponse{}, nil
}
