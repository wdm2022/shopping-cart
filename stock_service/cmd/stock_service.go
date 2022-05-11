package main

import (
	"flag"
)

var (
	orderServiceAddr = flag.String("order-service-addr", "localhost:50000", "address of the order service")
)

func main() {
	//flag.Parse()
	//
	//orderServiceConn, err := grpc.Dial(*orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	//if err != nil {
	//	log.Fatalf("did not connect: %v", err)
	//}
	//defer orderServiceConn.Close()
	//
	//orderClient := order.NewOrderClient(orderServiceConn)
	//
	//// Contact the server and print out its response.
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	//
	//_, err = orderClient.Ping(ctx, &order.PingRequest{})
	//if err != nil {
	//	log.Fatalf("could not ping: %v", err)
	//}
	//log.Printf("Received response")

}
