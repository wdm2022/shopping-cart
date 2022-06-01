package stock

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"log"
	"net"
	stockApi "shopping-cart/api/proto/stock"
	mongo2 "shopping-cart/pkg/stock/mongo"
)

type stockServer struct {
	stockApi.StockServer
	stockConn *mongo2.StockConnection
}

// **************** Interface methods *********************

func (o stockServer) Ping(ctx context.Context, in *stockApi.EmptyMessage) (*stockApi.PingResponse, error) {
	fmt.Println("Received ping")
	return &stockApi.PingResponse{Message: "Stock"}, nil
}

func (o stockServer) Find(ctx context.Context, in *stockApi.FindRequest) (*stockApi.FindResponse, error) {
	fmt.Println("Received a find stock request for item: ", in.ItemId)

	stock, err := o.stockConn.FindStock(in.ItemId)
	if err != nil {
		return nil, err
	}

	return &stockApi.FindResponse{Stock: stock.Amount, Price: stock.Price}, nil
}

func (o stockServer) Subtract(ctx context.Context, in *stockApi.SubtractRequest) (*stockApi.EmptyMessage, error) {
	fmt.Println("Received a subtract from  stock request for item: ", in.ItemId, ", amount: ", in.Amount)

	err := o.stockConn.SubtractStock(in.ItemId, in.Amount)
	if err != nil {
		return nil, err
	}

	return &stockApi.EmptyMessage{}, nil
}

func (o stockServer) Add(ctx context.Context, in *stockApi.AddRequest) (*stockApi.EmptyMessage, error) {
	fmt.Println("Received an add to stock request for item: ", in.ItemId, ", amount: ", in.Amount)

	//var err = AddToStock(o.stockConn, in.ItemId, in.Amount)

	err := o.stockConn.AddStock(in.ItemId, in.Amount)
	if err != nil {
		return nil, err
	}

	return &stockApi.EmptyMessage{}, nil
}

func (o stockServer) Create(ctx context.Context, in *stockApi.CreateRequest) (*stockApi.CreateResponse, error) {
	fmt.Println("Received an create stock request with price: ", in.Price)

	item, err := o.stockConn.NewItem(in.Price)
	if err != nil {
		return nil, err
	}

	return &stockApi.CreateResponse{ItemId: item}, nil
}

func (o stockServer) TotalCost(ctx context.Context, in *stockApi.TotalCostRequest) (*stockApi.TotalCostResponse, error) {
	fmt.Println("Received a total cost request for the following items: ", in.ItemIds)

	//TODO: Create a db call which returns the price for each item
	totalCost, err := o.stockConn.CalculateTotalCost(in.ItemIds)
	if err != nil {
		return nil, err
	}

	return &stockApi.TotalCostResponse{TotalCost: totalCost}, nil
}

func (o stockServer) SubtractBatch(ctx context.Context, in *stockApi.SubtractBatchRequest) (*stockApi.EmptyMessage, error) {
	fmt.Println("Received a total cost request for the following items: ", in.ItemIds)

	err := o.stockConn.SubtractBatchStock(in.ItemIds)
	if err != nil {
		return nil, err
	}

	return &stockApi.EmptyMessage{}, nil
}

func (o stockServer) AddBatch(ctx context.Context, in *stockApi.AddBatchRequest) (*stockApi.EmptyMessage, error) {
	fmt.Println("Received a total cost request for the following items: ", in.ItemIds)

	err := o.stockConn.AddBatchStock(in.ItemIds)
	if err != nil {
		return nil, err
	}

	return &stockApi.EmptyMessage{}, nil
}

func RunGrpcServer(client *mongo.Client, port *int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		return err
	}

	stockConn := mongo2.Init(client)

	server := grpc.NewServer()
	stockApi.RegisterStockServer(server, &stockServer{stockConn: stockConn})

	log.Printf("server listening at %v", lis.Addr())
	return server.Serve(lis)
}

// *********************** Server methods **********************
//
//func CreateStock(conn *mongo2.OrdersConnection, price float32) (string, error) {
//	// TODO
//	return "Brownie", nil
//}
//
//func AddToStock(conn *mongo2.OrdersConnection, id string, amount uint32) error {
//	// TODO
//	return nil
//}
//
//func SubtractFromStock(conn *mongo2.OrdersConnection, id string, amount uint32) error {
//	// TODO
//	return nil
//}
//
//func FindStock(conn *mongo2.OrdersConnection, id string) (uint32, float32, error) {
//	// TODO
//	return 0, 0.0, nil
//}
