package containerhelper

import (
	"context"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
)

// CreateContainer should be used to create containers in integration test. Ex.: create container for mongodb
func CreateContainer(ctx context.Context, req testcontainers.ContainerRequest, exposedPort string) (container testcontainers.Container, port int, err error) {
	container, err = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return
	}

	natPort, err := container.MappedPort(ctx, nat.Port(exposedPort))
	if err != nil {
		return
	}

	return container, natPort.Int(), err
}
