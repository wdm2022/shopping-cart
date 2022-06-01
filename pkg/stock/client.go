package stock

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	stockApi "shopping-cart/api/proto/stock"
	"time"
)

var (
	client stockApi.StockClient
)

func Connect(address string) *grpc.ClientConn {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("couldn't connect to stock service: %v", err)
	}

	client = stockApi.NewStockClient(conn)
	return conn
}

func createContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Second*10)

}

func Find(request *stockApi.FindRequest) (*stockApi.FindResponse, error) {
	ctx, cel := createContext()
	defer cel()
	return client.Find(ctx, request)
}

func Subtract(request *stockApi.SubtractRequest) (*stockApi.EmptyMessage, error) {
	ctx, cel := createContext()
	defer cel()
	return client.Subtract(ctx, request)
}

func Add(request *stockApi.AddRequest) (*stockApi.EmptyMessage, error) {
	ctx, cel := createContext()
	defer cel()
	return client.Add(ctx, request)
}

func Create(request *stockApi.CreateRequest) (*stockApi.CreateResponse, error) {
	ctx, cel := createContext()
	defer cel()
	return client.Create(ctx, request)
}
