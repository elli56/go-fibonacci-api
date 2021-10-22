package grpcserver

import (
	"context"

	"github.com/elli56/fibo-api/pkg/proto"
	"github.com/elli56/fibo-api/pkg/service"
)

type GRPCServer struct {
	proto.UnimplementedFibonacciSlicerServer
	services *service.Service
}

func NewGRPCServer(services *service.Service) *GRPCServer {
	return &GRPCServer{services: services}
}

func (s *GRPCServer) FibonacciSlice(ctx context.Context, req *proto.FiboRequest) (*proto.FiboResponse, error) {
	result, err := s.services.Calculation.FiboSlice(req.GetX(), req.GetY())
	return &proto.FiboResponse{Result: result}, err
}
