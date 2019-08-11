package mongodb

import (
	"testing"
)

func TestNewDBClient(t *testing.T) {
	url := "mongodb://localhost:27017"
	mongoClient := NewDBClient(url)
	mongoClient.PingTest()
}
