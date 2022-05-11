package stock

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	stockApi "shopping-cart/api/proto/stock"
)

type stockServer struct {
	stockApi.StockServer
}

func (o stockServer) Ping(ctx context.Context, in *stockApi.PingRequest) (*stockApi.PingResponse, error) {
	fmt.Println("Received ping")
	return &stockApi.PingResponse{Message: "order"}, nil
}

func RunGrpcServer(port *int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	stockApi.RegisterStockServer(server, &stockServer{})

	log.Printf("server listening at %v", lis.Addr())
	return server.Serve(lis)
}
