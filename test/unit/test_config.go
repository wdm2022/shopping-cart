package unit

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"

	mongoDriver "go.mongodb.org/mongo-driver/mongo"

	mongoUtils "shopping-cart/pkg/utils/mongo"

	"github.com/google/uuid"
	tc "github.com/testcontainers/testcontainers-go"
)

var (
	ctx                    = context.Background()
	mongoTestConfigPayment = mongoUtils.Config{
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
