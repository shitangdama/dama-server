package mongodb

// import (
//     // "time"
// //     // "go.mongodb.org/mongo-driver/bson"
//     // "go.mongodb.org/mongo-driver/mongo"
//     // "go.mongodb.org/mongo-driver/mongo/options"
// )

// // MongoDB struct
// type MongoDB struct {
//     Client      *mongo.Client
//     Ctx         context.Context
//     ServiceAddr string
// }

// // NewDBClient new
// func NewDBClient(url string) *MongoDB {
//     ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
//     client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
//     if err != nil {
//         panic(err)
//     }
//     return &MongoDB{
//         Client: client,
//         Ctx:    ctx,
//     }
// }

// // func (mongoDB *MongoDB) Disconnect() error {
// // 	err = client.Disconnect(context.TODO())

// // 	if err != nil {
// // 		log.Fatal(err)
// // 		return err
// // 	}
// // 	return nil
// // }


// // // PingTest test
// // func (mongoDB *MongoDB) PingTest() {
// // 	err := mongoDB.Client.Ping(mongoDB.Ctx, readpref.Primary())
// // 	if err != nil {
// // 		panic(err)
// // 	}
// // }