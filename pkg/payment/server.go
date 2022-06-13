package payment

import (
	"context"
	"fmt"
	"log"
	"net"
	paymentApi "shopping-cart/api/proto/payment"
	mongo2 "shopping-cart/pkg/payment/mongo"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

type paymentServer struct {
	paymentApi.PaymentServer
	paymentConn *mongo2.PaymentConnection
}

// **************** Interface methods *********************

func (o paymentServer) Ping(ctx context.Context, in *paymentApi.PingRequest) (*paymentApi.PingResponse, error) {
	fmt.Println("Received ping")
	return &paymentApi.PingResponse{Message: "payment"}, nil
}

func (o paymentServer) Pay(ctx context.Context, in *paymentApi.PayRequest) (*paymentApi.PayResponse, error) {
	fmt.Println("Received a find payment request for user: ", in.UserId, ", order: ", in.OrderId, ", for an amount of: ", in.Amount)
	//if in.TxId is 0, we assume that this does not belong to a saga
	isPaid, err := o.paymentConn.PayOrder(in.TxId, in.UserId, in.OrderId, in.Amount)

	if err != nil {
		return &paymentApi.PayResponse{Success: isPaid}, err
	}

	return &paymentApi.PayResponse{Success: isPaid}, nil
}

func (o paymentServer) Cancel(ctx context.Context, in *paymentApi.CancelRequest) (*paymentApi.EmptyMessage, error) {
	fmt.Println("Received a cancel payment request for user: ", in.UserId, ", order: ", in.OrderId)

	err := o.paymentConn.CancelOrder(in.UserId, in.OrderId)

	if err != nil {
		return nil, err
	}

	return &paymentApi.EmptyMessage{}, nil
}

func (o paymentServer) Status(ctx context.Context, in *paymentApi.StatusRequest) (*paymentApi.StatusResponse, error) {
	fmt.Println("Received a status payment request for user: ", in.UserId, ", order: ", in.OrderId)

	isPaid, err := o.paymentConn.StatusPayment(in.UserId, in.OrderId)

	if err != nil {
		return nil, err
	}

	return &paymentApi.StatusResponse{Paid: isPaid}, nil
}

func (o paymentServer) AddFunds(ctx context.Context, in *paymentApi.AddFundsRequest) (*paymentApi.AddFundsResponse, error) {
	fmt.Println("Received a add funds to user: ", in.UserId, ", amount: ", in.Amount)

	err := o.paymentConn.AddFunds(in.UserId, in.Amount)
	if err != nil {
		return nil, err
	}

	return &paymentApi.AddFundsResponse{Success: true}, nil
}

func (o paymentServer) CreateUser(ctx context.Context, in *paymentApi.EmptyMessage) (*paymentApi.CreateUserResponse, error) {
	fmt.Println("Received a create user request")

	user, err := o.paymentConn.CreateUser()
	if err != nil {
		return nil, err
	}

	return &paymentApi.CreateUserResponse{UserId: user}, nil
}

func (o paymentServer) FindUser(ctx context.Context, in *paymentApi.FindUserRequest) (*paymentApi.FindUserResponse, error) {
	fmt.Println("Received a find user request: ", in.UserId)

	user, err := o.paymentConn.FindUser(in.UserId)
	if err != nil {
		return nil, err
	}

	return &paymentApi.FindUserResponse{UserId: in.UserId, Credits: user.Credit}, nil
}

func (o paymentServer) Rollback(ctx context.Context, in *paymentApi.RollbackRequest) (*paymentApi.EmptyMessage, error) {
	fmt.Println("Received a rollback request: ", in.TxId)

	// TODO: Rollback transactions with the given transaction ID # Rahim
	//user, err := o.paymentConn.FindUser(in.UserId)
	//if err != nil {
	//	return nil, err
	//}

	return &paymentApi.EmptyMessage{}, nil
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
// func PayOrder(conn *mongo2.PaymentConnection, id string, id2 string, amount float32) (bool, error) {
// 	// TODO
// 	return true, nil
// }

// func CancelOrder(conn *mongo2.PaymentConnection, id string, id2 string) error {
// 	// TODO
// 	return nil
// }

// func StatusPayment(conn *mongo2.PaymentConnection, id string, id2 string) (bool, error) {
// 	// TODO
// 	return true, nil
// }

//func AddFundsToUser(conn *mongo2.PaymentConnection, id string, amount float32) (bool, error) {
//	// TODO
//	return true, nil
//}
//
//func CreateUser(conn *mongo2.PaymentConnection) (string, error) {
//	// TODO
//	return "Brownie", nil
//}
//
//func FindUser(conn *mongo2.PaymentConnection, id string) (float32, error) {
//	// TODO
//	return 0.0, nil
//}
