package usecases

import (
	"context"

	"github.com/lawmatsuyama/transactions/domain"
	log "github.com/sirupsen/logrus"
)

type TransactionUseCase struct {
	TransactionRepository domain.TransactionRepository
	UserRepository        domain.UserRepository
	MessagePublisher      domain.MessagePublisher
	SessionControl        domain.SessionControlRepository
}

func NewTransactionUseCase(transactionRepository domain.TransactionRepository, userRepository domain.UserRepository, messagePublisher domain.MessagePublisher, sessionControl domain.SessionControlRepository) TransactionUseCase {
	return TransactionUseCase{
		TransactionRepository: transactionRepository,
		UserRepository:        userRepository,
		MessagePublisher:      messagePublisher,
		SessionControl:        sessionControl,
	}
}

func (useCase TransactionUseCase) Save(ctx context.Context, userID string, transactions domain.Transactions) ([]domain.TransactionValidateResult, error) {

	err := validateRequest(userID, transactions)
	if err != nil {
		return nil, err
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

	err = useCase.SessionControl.WithSession(ctx, func(sc context.Context) error {
		err := useCase.TransactionRepository.Save(sc, transactions)
		if err != nil {
			return err
		}

		err = useCase.MessagePublisher.Publish(ctx, domain.ExchangeTransaction, "", transactions, 9)
		return err
	})

	return nil, err

}

func validateRequest(userID string, transactions domain.Transactions) error {
	l := log.WithField("user_id", userID)
	if userID == "" {
		l.WithError(domain.ErrInvalidUser).Error("Invalid userID request")
		return domain.ErrInvalidUser
	}

	if len(transactions) == 0 {
		l.WithError(domain.ErrInvalidTransaction).Error("There is no transactions to process")
		return domain.ErrInvalidTransaction
	}

	return nil
}
