package stock

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	stockApi "shopping-cart/api/proto/stock"
)

var (
	client stockApi.StockClient
)

func Connect(address *string) *grpc.ClientConn {
	conn, err := grpc.Dial(*address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("couldn't connect to stock service: %v", err)
	}

	client = stockApi.NewStockClient(conn)
	return conn
}
