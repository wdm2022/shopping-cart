package order

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	orderApi "shopping-cart/api/proto/order"
)

var (
	client orderApi.OrderClient
)

func Connect(address string) *grpc.ClientConn {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("couldn't connect to order service: %v", err)
	}

	client = orderApi.NewOrderClient(conn)
	return conn
}
