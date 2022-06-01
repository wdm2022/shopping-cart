package integration

import (
	"context"
	"fmt"
	"os"
	"path"
	"testing"
	"time"

	mongoDriver "go.mongodb.org/mongo-driver/mongo"

	tc "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	ctx = context.Background()
)

func setupSuite(tb testing.TB) func(tb testing.TB, client mongoDriver.Client) {
	envVars := make(map[string]string)
	envVars["MONGO_INITDB_ROOT_USERNAME"] = "LoFiBeats"
	envVars["MONGO_INITDB_ROOT_PASSWORD"] = "LoFiBeats"

	pathCwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(pathCwd)

	req := tc.ContainerRequest{
		Image:        "mongo:4.4",
		ExposedPorts: []string{"27017/tcp"},
		Env:          envVars,
		WaitingFor:   wait.ForLog("Listening on"),
		Mounts: tc.Mounts(
			tc.BindMount(path.Join(pathCwd, "init.js"), "/docker-entrypoint-initdb.d/init.js"),
		),
	}
	mongoC, err := tc.GenericContainer(ctx, tc.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	time.Sleep(10 * time.Second)

	if err != nil {
		tb.Error(err)
	}

	port, err := mongoC.MappedPort(ctx, "27017")
	fmt.Println("Port NUMBER:", port)
	if err != nil {
		tb.Error(err)
	}

	return func(tb testing.TB, client mongoDriver.Client) {
		mongoC.Terminate(ctx)
	}
}
