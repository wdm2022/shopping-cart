package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	orderApi "shopping-cart/api/proto/order"
	paymentApi "shopping-cart/api/proto/payment"
	stockApi "shopping-cart/api/proto/stock"
	"time"
)

var (
	orderServiceAddr   = flag.String("order-service-addr", "localhost:50000", "address of the order service")
	paymentServiceAddr = flag.String("payment-service-addr", "localhost:50001", "address of the payment service")
	stockServiceAddr   = flag.String("stock-service-addr", "localhost:50002", "address of the stock service")
)

func main() {
	flag.Parse()

	orderServiceConn, err := grpc.Dial(*orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("couldn't connect to order service: %v", err)
	}

	paymentServiceConn, err := grpc.Dial(*paymentServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("couldn't connect to payment service: %v", err)
	}

	stockServiceConn, err := grpc.Dial(*stockServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("couldn't connect to stock service: %v", err)
	}

	defer orderServiceConn.Close()
	defer paymentServiceConn.Close()
	defer stockServiceConn.Close()

	orderClient := orderApi.NewOrderClient(orderServiceConn)
	paymentClient := paymentApi.NewPaymentClient(paymentServiceConn)
	stockClient := stockApi.NewStockClient(stockServiceConn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	orderResponse, err := orderClient.Ping(ctx, &orderApi.PingRequest{})
	if err != nil {
		log.Fatalf("could not ping: %v", err)
	}
	log.Printf("Received response: %s", orderResponse.Message)

	paymentResponse, err := paymentClient.Ping(ctx, &paymentApi.PingRequest{})
	if err != nil {
		log.Fatalf("could not ping: %v", err)
	}
	log.Printf("Received response: %s", paymentResponse.Message)

	stockResponse, err := stockClient.Ping(ctx, &stockApi.PingRequest{})
	if err != nil {
		log.Fatalf("could not ping: %v", err)
	}
	log.Printf("Received response: %s", stockResponse.Message)
}
