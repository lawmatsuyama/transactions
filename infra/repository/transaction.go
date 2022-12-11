package repository

import (
	"context"
	"time"

	"github.com/lawmatsuyama/transactions/domain"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	batchSize int32 = 20
)

type TransactionDB struct {
	Client *mongo.Client
}

func NewTransactionDB(client *mongo.Client) TransactionDB {
	return TransactionDB{
		Client: client,
	}
}

func (db TransactionDB) Save(ctx context.Context, transactions domain.Transactions) error {
	c := db.Client.Database("account").Collection("transaction")
	models := bulkInsertModel(transactions)
	_, err := c.BulkWrite(ctx, models)
	if err != nil {
		log.WithField("transactions", transactions).WithError(err).Error("Failed to save transactions")
		return domain.ErrUnknow
	}

	return nil
}

func (db TransactionDB) Get(ctx context.Context, filterTrs domain.TransactionFilter) (trs domain.Transactions, err error) {
	l := log.WithField("filter", filterTrs)
	c := db.Client.Database("account").Collection("transaction")
	filter := bson.D{}
	filter = filterSimple(filter, "_id", filterTrs.ID, isZeroComparable[string])
	filter = filterSimple(filter, "user_id", filterTrs.UserID, isZeroComparable[string])
	filter = filterSimple(filter, "description", filterTrs.Description, isZeroComparable[string])
	filter = filterRange(filter, "created_at", filterTrs.DateFrom, filterTrs.DateTo, isZeroTime)
	filter = filterSimple(filter, "origin", filterTrs.Origin, isZeroComparable[domain.OriginChannel])
	filter = filterSimple(filter, "operation_type", filterTrs.OperationType, isZeroComparable[domain.OperationType])
	filter = filterRange(filter, "amount", filterTrs.AmountGreater, filterTrs.AmountLess, isZeroComparable[float64])

	sort := bson.D{bsonE("created_at", 1), bsonE("_id", 1)}
	opts := options.Find().
		SetSort(sort).
		SetBatchSize(batchSize).
		SetMaxTime(time.Second * 20).
		SetSkip(filterTrs.Paging.TransactionsSkip()).
		SetLimit(filterTrs.Paging.LimitTransactionsByPage())
	cur, err := c.Find(ctx, filter, opts)

	if err == mongo.ErrNoDocuments {
		err = domain.ErrTransactionsNotFound
		return
	}

	if err != nil {
		l.WithError(err).Error("Failed to get transactions")
		err = domain.ErrUnknow
		return
	}

	trs = domain.Transactions{}
	err = cur.All(ctx, &trs)
	if err != nil {
		l.WithError(err).Error("Failed to iterate over transactions")
		err = domain.ErrUnknow
		return
	}

	return trs, err
}
