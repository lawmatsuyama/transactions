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

type UserDB struct {
	Client *mongo.Client
}

func NewUserDB(client *mongo.Client) UserDB {
	return UserDB{
		Client: client,
	}
}

func (db UserDB) GetByID(ctx context.Context, id string) (domain.User, error) {
	l := log.WithField("user_id", id)
	c := db.Client.Database("account").Collection("user")
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	opts := options.FindOne().SetMaxTime(time.Second * 10)
	var user domain.User
	err := c.FindOne(ctx, filter, opts).Decode(&user)

	if err == mongo.ErrNoDocuments {
		l.Debug("User not found")
		return domain.User{}, domain.ErrUserNotFound
	}

	if err != nil {
		l.WithError(err).Error("Failed to get user by ID")
		return domain.User{}, domain.ErrUnknow
	}

	return user, err
}
