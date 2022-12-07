package repository_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/lawmatsuyama/transactions/infra/containerhelper"
	"github.com/lawmatsuyama/transactions/infra/repository"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gotest.tools/assert"
)

func TestStartMongodb(t *testing.T) {
	ctx := context.Background()
	container, port, err := startMongodb(ctx)
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		err := container.Terminate(ctx)
		if err != nil {
			t.Fatal(err)
		}
	}()

	os.Setenv("MONGODB_URI", fmt.Sprintf("mongodb://usercore:usercore@localhost:%d/?authSource=admin", port))
	repository.Start(ctx)
	cli, err := repository.GetClientDB(ctx)
	if err != nil {
		t.Fatal(err)
	}

	err = cli.Ping(ctx, nil)
	assert.NilError(t, err, "ping error should be nil")

	err = repository.CloseDB(ctx)
	assert.NilError(t, err, "close db error should be nil")
}

func startMongodb(ctx context.Context) (container testcontainers.Container, port int, err error) {
	req := testcontainers.ContainerRequest{
		Image: "mongo:4.0.19",
		Env: map[string]string{
			"MONGO_INITDB_DATABASE":      "account",
			"MONGO_INITDB_ROOT_USERNAME": "usercore",
			"MONGO_INITDB_ROOT_PASSWORD": "usercore",
		},
		ExposedPorts: []string{"27017/tcp"},
		WaitingFor:   wait.ForLog("waiting for connections on port"),
	}

	return containerhelper.CreateContainer(ctx, req, "27017")
}
