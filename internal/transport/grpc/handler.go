package grpc

import (
	"context"
	"log"
	"net"

	"github.com/misua/go-grpc-svc/internal/rocket"
	"google.golang.org/grpc"
)

type RocketService interface {
	GetRocketByID(ctx context.Context, id string) (rocket.Rocket, error)
	InsertRocket(ctx context.Context, rkt rocket.Rocket) (rocket.Rocket, error)
	DeleteRocket(ctx context.Context, id string) error
}

//NOTE will handle incoming gRPC request
type Handler struct {
	RocketService RocketService
}

func New(rktService RocketService) Handler {
	return Handler{
		RocketService: rktService,
	}
}

func (h Handler) Serve() error {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen : %v,err")
	}

	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
		return err

	}
	return nil
}

func (h Handler) GeRocket(ctx context.Context, req *rkt.GetRocketRequest) (*rkt.GetRocketResponse, error) {
	return &rkt.GetRocketResponse{}, nil
}
