package payment

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	paymentApi "shopping-cart/api/proto/payment"
)

type paymentServer struct {
	paymentApi.PaymentServer
}

func (o paymentServer) Ping(ctx context.Context, in *paymentApi.PingRequest) (*paymentApi.PingResponse, error) {
	fmt.Println("Received ping")
	return &paymentApi.PingResponse{Message: "order"}, nil
}

func RunGrpcServer(port *int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	paymentApi.RegisterPaymentServer(server, &paymentServer{})

	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		return err
	}

	return nil
}
