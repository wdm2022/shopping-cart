package order

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	orderApi "shopping-cart/api/proto/order"
)

type orderServer struct {
	orderApi.OrderServer
}

func (o orderServer) Ping(ctx context.Context, in *orderApi.PingRequest) (*orderApi.PingResponse, error) {
	fmt.Println("Received ping")
	return &orderApi.PingResponse{Message: "order"}, nil
}

func RunGrpcServer(port *int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	orderApi.RegisterOrderServer(server, &orderServer{})

	log.Printf("server listening at %v", lis.Addr())
	return server.Serve(lis)
}
