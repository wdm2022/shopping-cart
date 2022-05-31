package unit

import (
	"context"
	"fmt"
	"log"
	mongoPayment "shopping-cart/pkg/payment/mongo"
	"strings"
	"testing"

	mongoDriver "go.mongodb.org/mongo-driver/mongo"

	mongoUtils "shopping-cart/pkg/utils/mongo"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	tc "github.com/testcontainers/testcontainers-go"
)

var (
	ctx         = context.Background()
	mongoConfig = mongoUtils.Config{
		Host:     "localhost",
		Port:     27017,
		Username: "root",
		Password: "LoFiBeats",
		Database: "payment",
	}
)

func setupSuite(tb testing.TB) func(tb testing.TB, client mongoDriver.Client) {
	composeFilePaths := []string{"../../dev/docker-compose.yml"}
	identifier := strings.ToLower(uuid.New().String())
	compose := tc.NewLocalDockerCompose(composeFilePaths, identifier)
	execError := compose.WithCommand([]string{"up", "-d"}).Invoke()
	err := execError.Error
	if err != nil {
		log.Fatal(fmt.Errorf("Could not run compose file: %v - %v", composeFilePaths, err))
		return nil
	}

	return func(tb testing.TB, client mongoDriver.Client) {
		execError := compose.Down()
		err := execError.Error
		if err != nil {
			log.Fatal(fmt.Errorf("Could not run compose file: %v - %v", composeFilePaths, err))
		}
		_ = client.Disconnect(ctx)
	}
}

func TestCreateUser(t *testing.T) {
	teardownSuite := setupSuite(t)
	mongoClient := mongoPayment.Connect(&mongoConfig)
	defer teardownSuite(t, mongoClient)

	paymentConn := mongoPayment.Init(&mongoClient)

	user, err := paymentConn.CreateUser()
	assert.True(t, err == nil)
	assert.True(t, len(user) > 0)
}
