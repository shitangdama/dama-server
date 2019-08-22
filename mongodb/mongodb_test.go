package mongodb

import (
	"context"
	"fmt"
	"testing"
)

// You will be using this Trainer type later in the program
type Trainer struct {
	Name string
	Age  int
	City string
}

func TestNewDBClient(t *testing.T) {

	ash := Trainer{"Ash", 10, "Pallet Town"}
	// misty := Trainer{"Misty", 10, "Cerulean City"}
	// brock := Trainer{"Brock", 15, "Pewter City"}
	url := "mongodb://admin:kbr199sd5shi@localhost:27017"
	baseName := "data_test"
	mongoClient := NewDBClient(url, baseName)
	mongoClient.PingTest()
	database := mongoClient.GetDatabase()
	collection := database.Collection("collection")
	insertResult, err := collection.InsertOne(context.TODO(), ash)

	if err != nil {
		t.Log(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}
