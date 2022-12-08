package repository

import (
	"context"

	"github.com/lawmatsuyama/transactions/domain"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionDB struct {
	Client *mongo.Client
}

func (db TransactionDB) Save(ctx context.Context, transactions []domain.Transaction) error {
	c := db.Client.Database("account").Collection("transaction")
	models := bulkInsertModel(transactions)
	_, err := c.BulkWrite(ctx, models)
	if err != nil {
		log.WithField("transactions", transactions).WithError(err).Error("failed to save transactions")
	}

	return err
}
