package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Uri            string = "mongodb://localhost:27017/"
	DbName         string = "vueToDo"
	collectionName string = "vueToDo"
	Collection     *mongo.Collection
)

func init() {
	var clientOptions = options.Client().ApplyURI(Uri)

	var connection, connectionErr = mongo.Connect(context.TODO(), clientOptions)

	if connectionErr != nil {
		log.Panicln("error during connection")
		log.Panic(connectionErr)
	}

	Collection = connection.Database(DbName).Collection(collectionName)
	log.Println("conneced to the Db")
}
