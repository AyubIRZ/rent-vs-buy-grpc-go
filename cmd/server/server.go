package server

import (
	v1 "github.com/ayubirz/rent-vs-buy-grpc-go/internal/handlers/v1"
	"github.com/ayubirz/rent-vs-buy-grpc-go/internal/protos/v1/breakeven"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//RegisterGrpcServices registers all of gRPC API versions and services
func RegisterGrpcServices(server *grpc.Server) {
	v1Server := v1.Server{}
	breakeven.RegisterBreakevenServiceServer(server, &v1Server)

	reflection.Register(server)
}