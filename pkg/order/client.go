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

func CreateOrder(userId string) (*orderApi.CreateOrderResponse, error) {
	ctx, cel := context.WithTimeout(context.Background(), time.Second*3)
	defer cel()
	dto := orderApi.CreateOrderRequest{
		UserId: userId,
	}
	return client.CreateOrder(ctx, &dto)
}

func RemoveOrder(orderId string) (*orderApi.EmptyMessage, error) {
	ctx, cel := context.WithTimeout(context.Background(), time.Second*3)
	defer cel()
	dto := orderApi.RemoveOrderRequest{
		OrderId: orderId,
	}
	return client.RemoveOrder(ctx, &dto)
}

func GetOrder(orderId string) (*orderApi.GetOrderResponse, error) {
	ctx, cel := context.WithTimeout(context.Background(), time.Second*3)
	defer cel()
	dto := orderApi.GetOrderRequest{
		OrderId: orderId,
	}
	return client.GetOrder(ctx, &dto)
}

func AddItem(orderId string, itemId string) (*orderApi.EmptyMessage, error) {
	ctx, cel := context.WithTimeout(context.Background(), time.Second*3)
	defer cel()
	dto := orderApi.AddItemRequest{
		OrderId: orderId,
		ItemId:  itemId,
	}
	return client.AddItem(ctx, &dto)
}

func RemoveItem(orderId string, itemId string) (*orderApi.EmptyMessage, error) {
	ctx, cel := context.WithTimeout(context.Background(), time.Second*3)
	defer cel()
	dto := orderApi.RemoveItemRequest{
		OrderId: orderId,
		ItemId:  itemId,
	}
	return client.RemoveItem(ctx, &dto)
}

func Checkout(orderId string) (*orderApi.EmptyMessage, error) {
	ctx, cel := context.WithTimeout(context.Background(), time.Second*3)
	defer cel()
	dto := orderApi.CheckoutRequest{
		OrderId: orderId,
	}
	return client.Checkout(ctx, &dto)
}
