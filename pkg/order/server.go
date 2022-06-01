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
	mongo2 "shopping-cart/pkg/order/mongo"
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
		return nil, err
	}

	return &orderApi.GetOrderResponse{OrderId: order.OrderId.Hex(),
		Paid:      order.Paid,
		UserId:    order.UserId.Hex(),
		TotalCost: order.TotalCost,
		ItemIds: sf.Map(order.Items, func(t primitive.ObjectID) string {
			return t.Hex()
		})}, nil
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

	var err = checkoutOrder(in.OrderId)
	if err != nil {
		return nil, err
	}

	return &orderApi.EmptyMessage{}, nil
}

func RunGrpcServer(client *mongo.Client, port *int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		return err
	}

	orderConn := mongo2.Init(client)

	server := grpc.NewServer()
	orderApi.RegisterOrderServer(server, &orderServer{orderConn: orderConn})

	log.Printf("server listening at %v", lis.Addr())
	return server.Serve(lis)
}

// *********************** Server methods **********************

//func createNewOrder(o *mongo2.OrdersConnection, userId string) (string, error) {
//	// TODO: Create a new order id for the given UserId. Add it to the DB and return the new order number
//
//	res, err := o.EmptyOrder(userId)
//	if err != nil {
//		return "", err
//	}
//	return res, nil
//}
//
//func removeOrder(o *mongo2.OrdersConnection, orderId string) error {
//	// TODO: Remove order from DB
//	err := o.DeleteOrder(orderId)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func getOrder(orderId string) (orderDetails, error) {
//	// TODO: Collect order details from database replace for holders
//	var userId string = "Frodo"
//	var paid bool = false
//	var totalCost float32 = 0
//	var itemIds = []string{}
//
//	return orderDetails{
//		orderId:   orderId,
//		userId:    userId,
//		paid:      paid,
//		totalCost: totalCost,
//		itemIds:   itemIds}, nil
//}
//
//func addItemToOrder(orderId string, itemId string) error {
//	// TODO: Add item to order
//	return nil
//}
//
//func removeItemFromOrder(orderId string, itemId string) error {
//	// TODO: Remove item from order
//	return nil
//}

func checkoutOrder(id string) error {
	// TODO: Checkout order
	return nil
}

type orderDetails struct {
	orderId   string
	userId    string
	paid      bool
	totalCost float32
	itemIds   []string
}
