package tools

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MongoDBName = "apirest"
)

// MongoConn ...
var MongoConn = MongoClient()


// MongoClient ...
func MongoClient() *mongo.Client {

	opt := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), opt)

	if err != nil {
		log.Fatal(err.Error())
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("conection successfully")
	return client
}


// CheckConection ...
func CheckConection() bool {
	err := MongoConn.Ping(context.TODO(), nil)

	if err != nil {

		return false
	}

	return true
}
