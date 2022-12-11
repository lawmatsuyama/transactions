package usecases

import (
	"context"

	"github.com/lawmatsuyama/transactions/domain"
	log "github.com/sirupsen/logrus"
)

// TransactionUseCase implements interface domain.TransactionUseCase
type TransactionUseCase struct {
	TransactionRepository domain.TransactionRepository
	UserRepository        domain.UserRepository
	MessagePublisher      domain.MessagePublisher
	SessionControl        domain.SessionControlRepository
}

// NewTransactionUseCase returns a new TransactionUseCase
func NewTransactionUseCase(transactionRepository domain.TransactionRepository, userRepository domain.UserRepository, messagePublisher domain.MessagePublisher, sessionControl domain.SessionControlRepository) TransactionUseCase {
	return TransactionUseCase{
		TransactionRepository: transactionRepository,
		UserRepository:        userRepository,
		MessagePublisher:      messagePublisher,
		SessionControl:        sessionControl,
	}
}

// Save it will save transactions according to the given userID. Transactions and user must be valid, otherwise it will rollback and return error.
func (useCase TransactionUseCase) Save(ctx context.Context, userID string, transactions domain.Transactions) ([]domain.TransactionSaveResult, error) {
	l := log.WithField("user_id", userID)
	if userID == "" {
		l.WithError(domain.ErrInvalidUser).Error("Invalid userID request")
		return nil, domain.ErrInvalidUser
	}

	user, err := useCase.UserRepository.GetByID(ctx, userID)
	if err != nil {
		l.WithError(err).Error("Failed to get user")
		return nil, err
	}

	if err = user.IsValid(); err != nil {
		l.WithError(err).Error("User is invalid")
		return nil, err
	}

	trsResult, err := transactions.ValidateTransactionsToSave()
	if err != nil {
		l.WithError(err).Error("There are some transactions invalid")
		return trsResult, err
	}

	err = useCase.SessionControl.WithSession(ctx, func(sc context.Context) error {
		err := useCase.TransactionRepository.Save(sc, transactions)
		if err != nil {
			l.WithError(err).Error("Failed to save transactions")
			return err
		}

		err = useCase.MessagePublisher.Publish(ctx, domain.ExchangeTransactionSaved, "", transactions, 9)
		if err != nil {
			l.WithError(err).Error("Failed to publish transactions in exchange")
		}
		return err
	})

	return nil, err

}

// Get it will return transactions given filter.
func (useCase TransactionUseCase) Get(ctx context.Context, filter domain.TransactionFilter) (domain.TransactionsPaging, error) {
	l := log.WithField("filter", filter)
	err := filter.Validate()
	if err != nil {
		l.WithError(err).Error("Filter is invalid")
		return domain.TransactionsPaging{}, err
	}

	trs, err := useCase.TransactionRepository.Get(ctx, filter)
	if err != nil {
		return domain.TransactionsPaging{}, err
	}

	trsPage := domain.TransactionsPaging{Transactions: trs, Paging: filter.Paging}
	if err = trsPage.IsValid(); err != nil {
		return domain.TransactionsPaging{}, err
	}

	trsPage.SetNextPaging()
	return trsPage, nil
}
