package repository

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBClient represents database client
type DBClient struct {
	Client  *mongo.Client
	Once    sync.Once
	Options []*options.ClientOptions
}

var clientDB DBClient

var defaultPoolConnection *options.ClientOptions = options.Client().
	SetMinPoolSize(4).
	SetMaxConnIdleTime(time.Minute * 2).
	SetMaxPoolSize(50)

// Start initialize a connection with database. The opts argument should use to configure a custom pool connection with database.
func Start(ctx context.Context, opts ...*options.ClientOptions) *mongo.Client {
	clientDB.Options = opts
	cli, err := GetClientDB(ctx)
	if err != nil {
		panic(fmt.Sprintf("Cannot connect to mongodb: %v", err))
	}

	err = cli.Ping(ctx, nil)
	if err != nil {
		panic(fmt.Sprintf("Cannot ping to mongodb: %v", err))
	}

	return cli
}

// GetClientDB returns the DB client connected
func GetClientDB(ctx context.Context) (*mongo.Client, error) {
	var err error
	clientDB.Once.Do(func() {
		err = setClientDB(ctx)
	})

	return clientDB.Client, err
}

// CloseDB it will close database connection
func CloseDB(ctx context.Context) (err error) {
	if clientDB.Client == nil {
		return nil
	}
	err = clientDB.Client.Disconnect(ctx)
	return
}

func setClientDB(ctx context.Context) error {
	uri := dbURI()
	optURI := options.Client().ApplyURI(uri).SetDirect(true)

	var err error
	if len(clientDB.Options) == 0 {
		clientDB.Client, err = mongo.Connect(ctx, optURI, defaultPoolConnection)
		log.Info("Connect mongodb using default pool connection configurations")
		return err
	}

	opts := append(clientDB.Options, optURI)
	clientDB.Client, err = mongo.Connect(ctx, opts...)
	log.Info("Connect mongodb using custom pool connection configurations")
	return err
}

func dbURI() (uri string) {
	return os.Getenv("MONGODB_URI")
}
