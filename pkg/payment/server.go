package payment

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"log"
	"net"
	paymentApi "shopping-cart/api/proto/payment"
	mongo2 "shopping-cart/pkg/order/mongo"
)

type paymentServer struct {
	paymentApi.PaymentServer
	paymentConn *mongo2.OrdersConnection
}

// **************** Interface methods *********************

func (o paymentServer) Ping(ctx context.Context, in *paymentApi.PingRequest) (*paymentApi.PingResponse, error) {
	fmt.Println("Received ping")
	return &paymentApi.PingResponse{Message: "order"}, nil
}

func (o paymentServer) Pay(ctx context.Context, in *paymentApi.PayRequest) (*paymentApi.PayResponse, error) {
	fmt.Println("Received a find payment request for user: ", in.UserId, ", order: ", in.OrderId, ", for an amount of: ", in.Amount)

	var paymentSuccessful, err = PayOrder(o.paymentConn, in.UserId, in.OrderId, in.Amount)

	if err != nil {
		return nil, err
	}

	return &paymentApi.PayResponse{Success: paymentSuccessful}, nil
}

func (o paymentServer) Cancel(ctx context.Context, in *paymentApi.CancelRequest) (*paymentApi.EmptyMessage, error) {
	fmt.Println("Received a cancel payment request for user: ", in.UserId, ", order: ", in.OrderId)

	var err = CancelOrder(o.paymentConn, in.UserId, in.OrderId)

	if err != nil {
		return nil, err
	}

	return &paymentApi.EmptyMessage{}, nil
}

func (o paymentServer) Status(ctx context.Context, in *paymentApi.StatusRequest) (*paymentApi.StatusResponse, error) {
	fmt.Println("Received a status payment request for user: ", in.UserId, ", order: ", in.OrderId)

	var isPaid, err = StatusPayment(o.paymentConn, in.UserId, in.OrderId)

	if err != nil {
		return nil, err
	}

	return &paymentApi.StatusResponse{Paid: isPaid}, nil
}

func (o paymentServer) AddFunds(ctx context.Context, in *paymentApi.AddFundsRequest) (*paymentApi.AddFundsResponse, error) {
	fmt.Println("Received a add funds to user: ", in.UserId, ", amount: ", in.Amount)

	var isDone, err = AddFundsToUser(o.paymentConn, in.UserId, in.Amount)

	if err != nil {
		return nil, err
	}

	return &paymentApi.AddFundsResponse{Success: isDone}, nil
}

func (o paymentServer) CreateUser(ctx context.Context, in *paymentApi.EmptyMessage) (*paymentApi.CreateUserResponse, error) {
	fmt.Println("Received a create user request")

	var userId, err = CreateUser(o.paymentConn)

	if err != nil {
		return nil, err
	}

	return &paymentApi.CreateUserResponse{UserId: userId}, nil
}

func (o paymentServer) FindUser(ctx context.Context, in *paymentApi.FindUserRequest) (*paymentApi.FindUserResponse, error) {
	fmt.Println("Received a find user request: ", in.UserId)

	var credits, err = FindUser(o.paymentConn, in.UserId)

	if err != nil {
		return nil, err
	}

	return &paymentApi.FindUserResponse{UserId: in.UserId, Credits: credits}, nil
}

func RunGrpcServer(client *mongo.Client, port *int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		return err
	}

	paymentConn := mongo2.Init(client)

	server := grpc.NewServer()
	paymentApi.RegisterPaymentServer(server, &paymentServer{paymentConn: paymentConn})

	log.Printf("server listening at %v", lis.Addr())
	return server.Serve(lis)
}

// *********************** Server methods **********************
func PayOrder(conn *mongo2.OrdersConnection, id []byte, id2 []byte, amount float32) (bool, error) {
	// TODO
	return true, nil
}

func CancelOrder(conn *mongo2.OrdersConnection, id []byte, id2 []byte) error {
	// TODO
	return nil
}

func StatusPayment(conn *mongo2.OrdersConnection, id []byte, id2 []byte) (bool, error) {
	// TODO
	return true, nil
}

func AddFundsToUser(conn *mongo2.OrdersConnection, id []byte, amount float32) (bool, error) {
	// TODO
	return true, nil
}

func CreateUser(conn *mongo2.OrdersConnection) ([]byte, error) {
	// TODO
	return []byte{}, nil
}

func FindUser(conn *mongo2.OrdersConnection, id []byte) (float32, error) {
	// TODO
	return 0.0, nil
}
