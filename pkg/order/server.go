package order

import (
	"context"
	"fmt"
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

	var newOrderId, err = createNewOrder(o.orderConn, in.UserId)

	if err != nil {
		return nil, err
	}

	return &orderApi.CreateOrderResponse{OrderId: newOrderId}, nil
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
		return &orderApi.GetOrderResponse{}, nil
	}

	return &orderApi.GetOrderResponse{OrderId: order.OrderId[:],
		Paid:      order.Paid,
		UserId:    order.UserId,
		TotalCost: order.TotalCost,
		ItemIds:   order.Items}, nil
}

func (o orderServer) AddItem(ctx context.Context, in *orderApi.AddItemRequest) (*orderApi.EmptyMessage, error) {
	fmt.Println("Received an add item: ", in.ItemId, " request for order: ", in.OrderId)

	addItemToOrder(in.OrderId, in.ItemId)

	return &orderApi.EmptyMessage{}, nil
}

func (o orderServer) RemoveItem(ctx context.Context, in *orderApi.RemoveItemRequest) (*orderApi.EmptyMessage, error) {
	fmt.Println("Received a remove item: ", in.ItemId, " request for order: ", in.OrderId)

	removeItemFromOrder(in.OrderId, in.ItemId)

	return &orderApi.EmptyMessage{}, nil
}

func (o orderServer) Checkout(ctx context.Context, in *orderApi.CheckoutRequest) (*orderApi.EmptyMessage, error) {
	fmt.Println("Received a checkout request for order: ", in.OrderId)

	checkoutOrder(in.OrderId)

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

func createNewOrder(o *mongo2.OrdersConnection, userId []byte) ([]byte, error) {
	// TODO: Create a new order id for the given UserId. Add it to the DB and return the new order number

	res, err := o.EmptyOrder(userId)
	if err != nil {
		return []byte{}, err
	}
	return res, nil
}

func removeOrder(o *mongo2.OrdersConnection, orderId []byte) error {
	// TODO: Remove order from DB
	err := o.DeleteOrder(orderId)
	if err != nil {
		return err
	}
	return nil
}

func getOrder(orderId []byte) orderDetails {
	//TODO: Collect order details from database replace for holders
	var userId []byte
	var paid bool = false
	var totalCost float32 = 0
	var itemIds = [][]byte{{}}

	return orderDetails{
		orderId:   orderId,
		userId:    userId,
		paid:      paid,
		totalCost: totalCost,
		itemIds:   itemIds}
}

func addItemToOrder(orderId []byte, itemId []byte) {
	// TODO: Add item to order
}

func removeItemFromOrder(orderId []byte, itemId []byte) {
	// TODO: Remove item from order
}

func checkoutOrder(id []byte) {
	// TODO: Checkout order
}

type orderDetails struct {
	orderId   []byte
	userId    []byte
	paid      bool
	totalCost float32
	itemIds   [][]byte
}
