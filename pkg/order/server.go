package order

import (
	"context"
	"fmt"
	sf "github.com/sa-/slicefunk"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"log"
	"net"
	orderApi "shopping-cart/api/proto/order"
	paymentApi "shopping-cart/api/proto/payment"
	stockApi "shopping-cart/api/proto/stock"
	mongo2 "shopping-cart/pkg/order/mongo"
	"shopping-cart/pkg/payment"
	"shopping-cart/pkg/stock"
)

type orderServer struct {
	orderApi.OrderServer
	orderConn *mongo2.OrdersConnection
}

// **************** Interface methods *********************

func (o orderServer) Ping(ctx context.Context, in *orderApi.EmptyMessage) (*orderApi.PingResponse, error) {
	fmt.Println("Received ping")
	return &orderApi.PingResponse{Message: "order"}, nil
}

func (o orderServer) CreateOrder(ctx context.Context, in *orderApi.CreateOrderRequest) (*orderApi.CreateOrderResponse, error) {
	fmt.Println("Received a new order request for user: ", in.UserId)

	id, err := o.orderConn.EmptyOrder(in.UserId)

	if err != nil {
		return nil, err
	}

	return &orderApi.CreateOrderResponse{OrderId: id}, nil
}

func (o orderServer) RemoveOrder(ctx context.Context, in *orderApi.RemoveOrderRequest) (*orderApi.EmptyMessage, error) {
	fmt.Println("Received a remove order request for order: ", in.OrderId)

	err := o.orderConn.DeleteOrder(in.OrderId)
	if err != nil {
		return nil, err
	}

	return &orderApi.EmptyMessage{}, nil
}

func (o orderServer) GetOrder(ctx context.Context, in *orderApi.GetOrderRequest) (*orderApi.GetOrderResponse, error) {
	fmt.Println("Received a get order detail request for order: ", in.OrderId)

	order, err := o.orderConn.FindOrder(in.OrderId)
	if err != nil {
		fmt.Println("err when getting total cost", err)
		return nil, err
	}

	itemList := sf.Map(order.Items, func(t primitive.ObjectID) string {
		return t.Hex()
	})

	totalCost, err := stock.TotalCost(&stockApi.TotalCostRequest{ItemIds: itemList})
	if err != nil {
		fmt.Println("err when getting total cost", err)
		return nil, err
	}
	return &orderApi.GetOrderResponse{OrderId: order.OrderId.Hex(),
		Paid:      order.Paid,
		UserId:    order.UserId.Hex(),
		TotalCost: totalCost.TotalCost,
		ItemIds:   itemList,
	}, nil
}

func (o orderServer) AddItem(ctx context.Context, in *orderApi.AddItemRequest) (*orderApi.EmptyMessage, error) {
	fmt.Println("Received an add item: ", in.ItemId, " request for order: ", in.OrderId)

	err := o.orderConn.AddItem(in.OrderId, in.ItemId)
	if err != nil {
		return nil, err
	}

	return &orderApi.EmptyMessage{}, nil
}

func (o orderServer) RemoveItem(ctx context.Context, in *orderApi.RemoveItemRequest) (*orderApi.EmptyMessage, error) {
	fmt.Println("Received a remove item: ", in.ItemId, " request for order: ", in.OrderId)

	err := o.orderConn.RemoveItem(in.OrderId, in.ItemId)
	if err != nil {
		return nil, err
	}

	return &orderApi.EmptyMessage{}, nil
}

func (o orderServer) Checkout(ctx context.Context, in *orderApi.CheckoutRequest) (*orderApi.EmptyMessage, error) {
	fmt.Println("Received a checkout request for order: ", in.OrderId)

	//TODO - Implement the DB calls in the corresponding microservices. #Rahim :)

	// TODO - Add checks which see if everything worked and add logic which maybe reserves the items for this order?
	// use random as an txid
	txId := primitive.NewObjectID().Hex()

	err := o.orderConn.StartTransaction(txId)
	if err != nil {
		return nil, err
	}

	// Get the order details from the db
	order, orderErr := o.orderConn.FindOrder(in.OrderId)
	userId := order.UserId.Hex()
	itemIds := sf.Map(order.Items, func(t primitive.ObjectID) string {
		return t.Hex()
	})
	if orderErr != nil {
		log.Println("could not find order")
		return nil, orderErr
	}

	// Calculate the total cost of the order
	totalCost, stockErr := stock.TotalCost(&stockApi.TotalCostRequest{ItemIds: itemIds})
	if stockErr != nil {
		log.Println("could not calculate total cost")
		return nil, stockErr
	}

	// Process payment
	_, payErr := payment.Pay(&paymentApi.PayRequest{TxId: txId, UserId: userId, OrderId: in.OrderId, Amount: totalCost.TotalCost})
	if payErr != nil {
		fmt.Println("cost ", totalCost)
		fmt.Println("could not pay", userId, payErr)
		return nil, payErr
	}

	// Remove items in the order from stock
	_, stockErr2 := stock.SubtractBatch(&stockApi.SubtractBatchRequest{TxId: txId, ItemIds: itemIds})
	if stockErr2 != nil {
		fmt.Println("could not subtract", stockErr2)
		// Something went wrong while subtracting the batch, payment has to be reverted
		_, rollbackErr := payment.Rollback(&paymentApi.RollbackRequest{TxId: txId})
		if rollbackErr != nil {
			stockErr2 = rollbackErr
		}
		return nil, stockErr2
	}

	err = o.orderConn.EndTransaction(txId)
	if err != nil {
		log.Println("could not end tx ", err)
		return nil, err
	}

	// TODO: add succes/fail error messages on return
	return &orderApi.EmptyMessage{}, nil
}

func RunGrpcServer(client *mongo.Client, port *int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		return err
	}

	orderConn := mongo2.Init(client)

	transactions, err := orderConn.FindOpenTransactions()
	if err != nil {
		return err
	}
	fmt.Println(transactions)

	server := grpc.NewServer()
	orderApi.RegisterOrderServer(server, &orderServer{orderConn: orderConn})

	log.Printf("server listening at %v", lis.Addr())
	return server.Serve(lis)
}

type orderDetails struct {
	orderId   string
	userId    string
	paid      bool
	totalCost float32
	itemIds   []string
}
