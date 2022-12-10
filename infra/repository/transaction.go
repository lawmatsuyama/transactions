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
	limitDocuments int32 = 20
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

func (db TransactionDB) Get(ctx context.Context, filterTrs domain.TransactionFilter) (trsPage domain.TransactionsPaging, err error) {
	c := db.Client.Database("account").Collection("transaction")
	filter := bson.D{}
	filter = filterSimple(filter, "_id", filterTrs.ID, isZeroComparable[string])
	filter = filterSimple(filter, "user_id", filterTrs.UserID, isZeroComparable[string])
	filter = filterSimple(filter, "description", filterTrs.Description, isZeroComparable[string])
	filter = filterRange(filter, "created_at", filterTrs.DateFrom, filterTrs.DateTo, isZeroTime)
	filter = filterSimple(filter, "origin", filterTrs.Origin, isZeroComparable[domain.OriginChannel])
	filter = filterSimple(filter, "operation_type", filterTrs.OperationType, isZeroComparable[domain.OperationType])
	filter = filterRange(filter, "amount", filterTrs.AmountGreater, filterTrs.AmountLess, isZeroComparable[float64])
	page := filterTrs.Page()

	sort := bson.D{bsonE("created_at", 1), bsonE("_id", 1)}
	opts := options.Find().SetSort(sort).SetBatchSize(limitDocuments).SetMaxTime(time.Second * 20).SetSkip(page).SetLimit(int64(limitDocuments))
	cur, err := c.Find(ctx, filter, opts)

	if err == mongo.ErrNoDocuments {
		err = domain.ErrTransactionsNotFound
		return
	}

	if err != nil {
		err = domain.ErrUnknow
		return
	}

	trs := []*domain.Transaction{}
	err = cur.All(ctx, &trs)
	if err != nil {
		return
	}

	if len(trs) == 0 {
		err = domain.ErrTransactionsNotFound
		return
	}

	trsPage = domain.TransactionsPaging{Transactions: trs}
	if len(trs) >= int(limitDocuments) {
		trsPage.Paging = &domain.Paging{
			Page: page + int64(len(trs)),
		}
	}

	return trsPage, err
}
