package order

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	orderApi "shopping-cart/api/proto/order"
)

type orderServer struct {
	orderApi.OrderServer
}

// **************** Interface methods *********************

func (o orderServer) Ping(ctx context.Context, in *orderApi.EmptyMessage) (*orderApi.PingResponse, error) {
	fmt.Println("Received ping")
	return &orderApi.PingResponse{Message: "order"}, nil
}

func (o orderServer) CreateOrder(ctx context.Context, in *orderApi.CreateOrderRequest) (*orderApi.CreateOrderResponse, error) {
	fmt.Println("Received a new order request for user: ", in.UserId)

	var newOrderId = createNewOrder(in.UserId)

	return &orderApi.CreateOrderResponse{OrderId: newOrderId}, nil
}

func (o orderServer) RemoveOrder(ctx context.Context, in *orderApi.RemoveOrderRequest) (*orderApi.EmptyMessage, error) {
	fmt.Println("Received a remove order request for order: ", in.OrderId)

	removeOrder(in.OrderId)

	return &orderApi.EmptyMessage{}, nil
}

func (o orderServer) GetOrder(ctx context.Context, in *orderApi.GetOrderRequest) (*orderApi.GetOrderResponse, error) {
	fmt.Println("Received a get order detail request for order: ", in.OrderId)

	var order = getOrder(in.OrderId)

	return &orderApi.GetOrderResponse{OrderId: order.orderId,
		Paid:      order.paid,
		UserId:    order.userId,
		TotalCost: order.totalCost,
		ItemIds:   order.itemIds}, nil
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

func RunGrpcServer(port *int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	orderApi.RegisterOrderServer(server, &orderServer{})

	log.Printf("server listening at %v", lis.Addr())
	return server.Serve(lis)
}

// *********************** Server methods **********************

func createNewOrder(userId uint32) uint32 {
	// TODO: Create a new order id for the given UserId. Add it to the DB and return the new order number
	return 1
}

func removeOrder(orderId uint32) {
	// TODO: Remove order from DB
}

func getOrder(orderId uint32) orderDetails {
	//TODO: Collect order details from database replace for holders
	var userId uint32 = 1
	var paid bool = false
	var totalCost float32 = 0
	var itemIds []uint32 = []uint32{1}

	return orderDetails{
		orderId:   orderId,
		userId:    userId,
		paid:      paid,
		totalCost: totalCost,
		itemIds:   itemIds}
}

func addItemToOrder(orderId uint32, itemId uint32) {
	// TODO: Add item to order
}

func removeItemFromOrder(orderId uint32, itemId uint32) {
	// TODO: Remove item from order
}

func checkoutOrder(id uint32) {
	// TODO: Checkout order
}

type orderDetails struct {
	orderId   uint32
	userId    uint32
	paid      bool
	totalCost float32
	itemIds   []uint32
}
