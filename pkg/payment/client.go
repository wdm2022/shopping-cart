package payment

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	paymentApi "shopping-cart/api/proto/payment"
	"time"
)

var (
	client paymentApi.PaymentClient
)

func Connect(address string) *grpc.ClientConn {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("couldn't connect to payment service: %v", err)
	}

	client = paymentApi.NewPaymentClient(conn)
	return conn
}

func createContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Second*10)

}

func Pay(request *paymentApi.PayRequest) (*paymentApi.PayResponse, error) {
	ctx, cel := createContext()
	defer cel()
	return client.Pay(ctx, request)
}

func Cancel(request *paymentApi.CancelRequest) (*paymentApi.EmptyMessage, error) {
	ctx, cel := createContext()
	defer cel()
	return client.Cancel(ctx, request)
}

func Status(request *paymentApi.StatusRequest) (*paymentApi.StatusResponse, error) {
	ctx, cel := createContext()
	defer cel()
	return client.Status(ctx, request)
}

func AddFunds(request *paymentApi.AddFundsRequest) (*paymentApi.AddFundsResponse, error) {
	ctx, cel := createContext()
	defer cel()
	return client.AddFunds(ctx, request)
}

func CreateUser(request *paymentApi.EmptyMessage) (*paymentApi.CreateUserResponse, error) {
	ctx, cel := createContext()
	defer cel()
	return client.CreateUser(ctx, request)
}

func FindUser(request *paymentApi.FindUserRequest) (*paymentApi.FindUserResponse, error) {
	ctx, cel := createContext()
	defer cel()
	return client.FindUser(ctx, request)
}
