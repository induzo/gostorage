package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Config is a conf for the mongo database
type Config struct {
	Host           string
	Port           string
	User           string
	Password       string
	ConnectTimeout time.Duration
}

// NewCli connects to db and return a connection pool
func NewCli(mongoConf Config) (*mongo.Client, error) {
	uri := "mongodb://" +
		mongoConf.User + ":" +
		mongoConf.Password + "@" +
		mongoConf.Host + ":" +
		mongoConf.Port

	ctx, cancel := context.WithTimeout(
		context.Background(),
		mongoConf.ConnectTimeout,
	)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("NewCli: mongo.NewClient %v", err)
	}

	errP := client.Ping(ctx, nil)
	if errP != nil {
		return nil, fmt.Errorf("NewCli: client.Ping() %v", errP)
	}

	return client, nil
}
