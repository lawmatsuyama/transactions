package usecases

import (
	"context"

	"github.com/lawmatsuyama/transactions/domain"
	log "github.com/sirupsen/logrus"
)

const (
	exchangeTransaction = "EXC.TRANSACTION.SAVED"
)

type TransactionUseCase struct {
	TransactionRepository domain.TransactionRepository
	UserRepository        domain.UserRepository
	MessagePublisher      domain.MessagePublisher
}

func NewTransactionUseCase(transactionRepository domain.TransactionRepository, userRepository domain.UserRepository, messagePublisher domain.MessagePublisher) TransactionUseCase {
	return TransactionUseCase{
		TransactionRepository: transactionRepository,
		UserRepository:        userRepository,
		MessagePublisher:      messagePublisher,
	}
}

func (useCase TransactionUseCase) Save(ctx context.Context, userID string, transactions domain.Transactions) ([]domain.TransactionValidateResult, error) {
	l := log.WithField("user_id", userID)
	if userID == "" {
		l.WithError(domain.ErrInvalidUser).Error("Invalid userID request")
		return nil, domain.ErrInvalidUser
	}

	user, err := useCase.UserRepository.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if err = user.IsValid(); err != nil {
		return nil, err
	}

	trsResult, err := transactions.ValidateTransactions()
	if err != nil {
		return trsResult, err
	}

	err = useCase.TransactionRepository.Save(ctx, transactions)
	if err != nil {
		return nil, err
	}

	err = useCase.MessagePublisher.Publish(ctx, exchangeTransaction, "", transactions, 9)
	if err != nil {
		return nil, err
	}

	return nil, err
}
