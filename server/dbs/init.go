package dbs

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Uri            string = "mongodb://localhost:27017"
	DbName         string = "VueToDo"
	collectionName string = "VueToDo"
	Collection     *mongo.Collection
)

func init() {
	var clientOptions = options.Client().ApplyURI(Uri)
	var connect, connectionErr = mongo.Connect(context.TODO(), clientOptions)
	if connectionErr != nil {
		fmt.Println("Couldnt't connect to DB")
		log.Panicln(connectionErr)
	}
	Collection = connect.Database(DbName).Collection(collectionName)
}
