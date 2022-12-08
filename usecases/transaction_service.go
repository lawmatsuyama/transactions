package usecases

import (
	"context"

	"github.com/lawmatsuyama/transactions/domain"
	log "github.com/sirupsen/logrus"
)

type TransactionService struct {
	TransactionRepository domain.TransactionRepository
	UserRepository        domain.UserRepository
}

func (service TransactionService) Save(ctx context.Context, userID string, transactions domain.Transactions) ([]domain.TransactionValidateResult, error) {
	l := log.WithField("user_id", userID)
	if userID == "" {
		l.WithError(domain.ErrInvalidUser).Error("Invalid userID request")
		return nil, domain.ErrInvalidUser
	}

	user, err := service.UserRepository.GetByID(ctx, userID)
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

	err = service.TransactionRepository.Save(ctx, transactions)
	return nil, err
}
