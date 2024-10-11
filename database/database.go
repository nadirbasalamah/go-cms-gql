package database

import (
	"context"
	"go-cms-gql/utils"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client   *mongo.Client
	Database *mongo.Database
}

var DB MongoInstance

func Connect(dbName string) error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	opts := options.Client().ApplyURI(utils.GetValue("MONGO_URI")).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	// Send a ping to confirm a successful connection
	var result bson.M

	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		panic(err)
	}

	db := client.Database(dbName)

	DB = MongoInstance{
		Client:   client,
		Database: db,
	}

	log.Println("connected to the database")

	return nil
}

func GetCollection(name string) *mongo.Collection {
	return DB.Database.Collection(name)
}
