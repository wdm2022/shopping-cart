package payment

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	paymentApi "shopping-cart/api/proto/payment"
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
