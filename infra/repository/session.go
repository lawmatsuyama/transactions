package repository

import (
	"context"
	"fmt"

	"github.com/lawmatsuyama/transactions/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

// SessionControlDB implements interface domain.SessionControlRepository
type SessionControlDB struct {
	Client *mongo.Client
}

// NewSessionControlDB returns a new SessionControlDB
func NewSessionControlDB(client *mongo.Client) SessionControlDB {
	return SessionControlDB{
		Client: client,
	}
}

// WithSession start a new session on mongo DB and execute the given FuncDBSession inside the session.
// if FuncDBSession returns no errors then it will commit the transactions, otherwise it will rollback the transactions.
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
				return fmt.Errorf("original error: %v error abort transaction: %v", err, errAbort)
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
