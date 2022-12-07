package messagebroker_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/docker/go-connections/nat"
	"github.com/lawmatsuyama/transactions/domain"
	"github.com/lawmatsuyama/transactions/infra/messagebroker"
	"github.com/streadway/amqp"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gotest.tools/assert"
)

type MessageTest struct {
	Message string `json:"message"`
}

func TestPublisherConsumerRabbitmq(t *testing.T) {
	ctxBase := context.Background()
	ctx, cancel := context.WithCancel(ctxBase)
	container, err := startRabbitmq(ctx)
	if err != nil {
		t.Error(err)
	}

	port, err := container.MappedPort(ctx, nat.Port("5672"))
	if err != nil {
		t.Error(err)
	}

	defer func() {
		cancel()
		if err := container.Terminate(ctxBase); err != nil {
			t.Fatalf("failed to terminate container: %s", err.Error())
		}
	}()

	os.Setenv("MESSAGE_BROKER_URL", fmt.Sprintf("amqp://guest:guest@localhost:%d?heartbeat=30&connection_timeout=120", port.Int()))
	var got MessageTest
	domain.AddTaskCount()
	setup := domain.BrokerSetup(func() error {
		_, err := messagebroker.CreateQueue("queue_test", true, nil)
		if err != nil {
			return err
		}

		testConsume := func(m amqp.Delivery) {
			defer domain.DoneTask()
			err := json.Unmarshal(m.Body, &got)
			if err != nil {
				t.Error(err)
			}
		}

		messagebroker.Consume(ctx, "queue_test", "service_test", testConsume)
		return nil
	})

	messagebroker.Start(ctx, setup)

	err = messagebroker.Publish(ctx, "", "queue_test", MessageTest{Message: "ok"}, 9)
	if err != nil {
		t.Error(err)
	}

	domain.WaitUntilAllTasksDone()
	assert.Equal(t, got, MessageTest{Message: "ok"}, "message test should be equal")
}

func startRabbitmq(ctx context.Context) (testcontainers.Container, error) {
	req := testcontainers.ContainerRequest{
		Image: "rabbitmq:3-management-alpine",
		Env: map[string]string{
			"RABBITMQ_DEFAULT_USER": "guest",
			"RABBITMQ_DEFAULT_PASS": "guest",
		},
		ExposedPorts: []string{"5672/tcp"},
		WaitingFor:   wait.ForLog("Server startup complete"),
	}

	return testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
}
