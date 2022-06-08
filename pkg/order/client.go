package order

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	orderApi "shopping-cart/api/proto/order"
	"time"
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

func createContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Second*10)

}

func CreateOrder(request *orderApi.CreateOrderRequest) (*orderApi.CreateOrderResponse, error) {
	ctx, cel := createContext()
	defer cel()
	return client.CreateOrder(ctx, request)
}

func RemoveOrder(request *orderApi.RemoveOrderRequest) (*orderApi.EmptyMessage, error) {
	ctx, cel := createContext()
	defer cel()
	return client.RemoveOrder(ctx, request)
}

func GetOrder(request *orderApi.GetOrderRequest) (*orderApi.GetOrderResponse, error) {
	ctx, cel := createContext()
	defer cel()
	return client.GetOrder(ctx, request)
}

func AddItem(request *orderApi.AddItemRequest) (*orderApi.EmptyMessage, error) {
	ctx, cel := createContext()
	defer cel()
	return client.AddItem(ctx, request)
}

func RemoveItem(request *orderApi.RemoveItemRequest) (*orderApi.EmptyMessage, error) {
	ctx, cel := createContext()
	defer cel()
	return client.RemoveItem(ctx, request)
}

func Checkout(request *orderApi.CheckoutRequest) (*orderApi.EmptyMessage, error) {
	ctx, cel := createContext()
	defer cel()
	return client.Checkout(ctx, request)
}
