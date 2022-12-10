package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/lawmatsuyama/transactions/infra/apimanager"
	"github.com/lawmatsuyama/transactions/infra/messagebroker"
	"github.com/lawmatsuyama/transactions/infra/repository"
	log "github.com/sirupsen/logrus"
)

// @title Transactions
// @version 1.0
// @description Save and list transactions made by user.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost
// @BasePath /v1

var (
	serviceName = "transactions"
)

func init() {
	LoadEnv()
	LoggerSetup()
}

func main() {
	ctx, cancel := start()
	defer shutdown(ctx, cancel)
	waitSignal()
}

func start() (ctx context.Context, cancel context.CancelFunc) {
	log.WithField("service", serviceName).Info("Starting service")
	ctxWithCancel, cancel := context.WithCancel(context.Background())
	StartDependencies(ctxWithCancel)
	log.WithField("service", serviceName).Info("Service is ready")
	return ctxWithCancel, cancel
}

func shutdown(ctx context.Context, cancel context.CancelFunc) {
	cancel()
	messagebroker.Shutdown()
	err := repository.CloseDB(ctx)
	if err != nil {
		log.Warn("Failed to close database")
	}
	apimanager.ShutdownAPI()
	log.WithField("service", serviceName).Info("Service finished")
}

func waitSignal() {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	s := <-sigc
	log.WithField("service", serviceName).Infof("Signal received [%v] canceling everything", s)
}
