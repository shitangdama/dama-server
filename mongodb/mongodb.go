package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoDB struct 只负责生成mongo的client
type MongoDB struct {
	Client      *mongo.Client
	Ctx         context.Context
	ServiceAddr string
	DBName      string
}

// NewDBClient new
func NewDBClient(url string, name string) *MongoDB {
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		panic(err)
	}
	return &MongoDB{
		ServiceAddr: url,
		Client:      client,
		Ctx:         ctx,
		DBName:      name,
	}
}

// PingTest test
func (mongoDB *MongoDB) PingTest() {
	err := mongoDB.Client.Ping(mongoDB.Ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
}

// GetDatabase returns a handle for a given database.
func (mongoDB *MongoDB) GetDatabase() *mongo.Database {
	return mongoDB.Client.Database(mongoDB.DBName)
}
