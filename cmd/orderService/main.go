package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"shoppingcart/api/proto/order"
)

var (
	port = flag.Int("port", 50000, "The gRPC server port")
)

type orderserver struct {
	order.OrderServer
}

func (o orderserver) Ping(ctx context.Context, in *order.PingRequest) (*order.PingResponse, error) {
	fmt.Println("Received ping")
	return &order.PingResponse{}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	order.RegisterOrderServer(server, &orderserver{})
	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
