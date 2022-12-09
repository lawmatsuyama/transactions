package repository

import (
	"context"
	"fmt"

	"github.com/lawmatsuyama/transactions/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type SessionControlDB struct {
	Client *mongo.Client
}

func NewSessionControlDB(client *mongo.Client) SessionControlDB {
	return SessionControlDB{
		Client: client,
	}
}

func (db SessionControlDB) WithSession(ctx context.Context, f domain.FuncDBSession) error {
	var session mongo.Session
	var err error
	if session, err = db.Client.StartSession(); err != nil {
		return err
	}

	if err = session.StartTransaction(); err != nil {
		return err
	}

	defer session.EndSession(ctx)

	err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		err = f(sc)
		if err != nil {
			errAbort := session.AbortTransaction(sc)
			if errAbort != nil {
				return fmt.Errorf("origina error: %w error abort transaction: %w", err, errAbort)
			}
			return err
		}

		if err = session.CommitTransaction(sc); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
