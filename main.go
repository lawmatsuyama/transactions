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
	repository.CloseDB(ctx)
	apimanager.ShutdownAPI()
	log.WithField("service", serviceName).Info("Service finished")
}

func waitSignal() {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	s := <-sigc
	log.WithField("service", serviceName).Infof("Signal received [%v] canceling everything", s)
}
