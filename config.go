package main

import (
	"context"

	"github.com/joho/godotenv"
	joonix "github.com/joonix/log"
	"github.com/lawmatsuyama/transactions/infra/apimanager"
	"github.com/lawmatsuyama/transactions/infra/messagebroker"
	"github.com/lawmatsuyama/transactions/infra/repository"
	"github.com/lawmatsuyama/transactions/usecases"
	log "github.com/sirupsen/logrus"
)

// LoadEnv load all enviroment variables from file if it is not already loaded
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.WithError(err).Error("couldnt load .env")
	}
}

// LoggerSetup setup log format
func LoggerSetup() {
	log.SetFormatter(joonix.NewFormatter(joonix.PrettyPrintFormat, joonix.DefaultFormat))
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(false)
}

func StartDependencies(ctxWithCancel context.Context) {
	brokeSetuper := messagebroker.NewSetuper()
	messagebroker.Start(ctxWithCancel, brokeSetuper)
	dbCli := repository.Start(context.Background())

	transactionRepository := repository.NewTransactionDB(dbCli)
	userRepository := repository.NewUserDB(dbCli)
	sessionControlRepository := repository.NewSessionControlDB(dbCli)

	publisher := messagebroker.NewMessagePublisher()

	transactionUseCase := usecases.NewTransactionUseCase(transactionRepository, userRepository, publisher, sessionControlRepository)

	transactionAPI := apimanager.NewTransactionAPI(transactionUseCase)
	handler := apimanager.NewHandler(transactionAPI)
	apimanager.StartAPI(ctxWithCancel, handler, "8080", "transactions")

}
