package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoDB struct
type MongoDB struct {
	Client      *mongo.Client
	Ctx         context.Context
	ServiceAddr string
}

// NewDBClient new
func NewDBClient(url string) *MongoDB {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		panic(err)
	}
	return &MongoDB{
		ServiceAddr: url,
		Client:      client,
		Ctx:         ctx,
	}
}

// PingTest test
func (mongoDB *MongoDB) PingTest() {
	err := mongoDB.Client.Ping(mongoDB.Ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
}

// InsertOne insert
func (mongoDB *MongoDB) InsertOne(data bson.D) {
	mongoDB.Client.Database("test").Collection("trainers")
	_, err := mongoDB.Client.InsertOne(context.TODO(), data)
	if err != nil {
		log.Fatal(err)
	}

}

// err = client.Disconnect(context.TODO())

// if err != nil {
//     log.Fatal(err)
// }
// fmt.Println("Connection to MongoDB closed.")
