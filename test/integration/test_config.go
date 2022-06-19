package integration

import (
	"context"
	"os"
	"path"
	"strconv"
	"strings"
	"testing"

	mongoUtils "shopping-cart/pkg/utils/mongo"

	mongoDriver "go.mongodb.org/mongo-driver/mongo"

	"github.com/docker/go-connections/nat"

	tc "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	ctx    = context.Background()
	mongoC tc.Container
)

const mongodbPort = 27017

func setupSuite(tb testing.TB) (mongoUtils.Config, func(tb testing.TB, client mongoDriver.Client)) {
	if testing.Short() {
		tb.Skip("skipping test in short mode.")
	}
	envVars := make(map[string]string)
	envVars["MONGO_INITDB_ROOT_USERNAME"] = "LoFiBeats"
	envVars["MONGO_INITDB_ROOT_PASSWORD"] = "LoFiBeats"

	pathCwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	exposedPort := strconv.Itoa(mongodbPort) + "/tcp"
	req := tc.ContainerRequest{
		Image:        "mongo:4.4",
		ExposedPorts: []string{exposedPort},
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

	if err != nil {
		tb.Error(err)
	}

	portStr := nat.Port(strconv.Itoa(mongodbPort))
	port, err := mongoC.MappedPort(ctx, portStr)

	if err != nil {
		tb.Error(err)
	}
	s := strings.Split(string(port), "/")
	portNum, err := nat.ParsePort(s[0])

	if err != nil {
		tb.Error(err)
	}

	host := "localhost:" + strconv.Itoa(portNum)
	connectionConfig := mongoUtils.Config{
		Hosts:    []string{host},
		Username: "root",
		Password: "LoFiBeats",
		Database: "",
	}

	return connectionConfig, func(tb testing.TB, client mongoDriver.Client) {
		mongoC.Terminate(ctx)
	}
}
