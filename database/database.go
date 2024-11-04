package database

import (
	"context"
	"fmt"
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
	var (
		DB_PROTOCOL = utils.GetValue("DB_PROTOCOL")
		DB_USER     = utils.GetValue("DB_USER")
		DB_PASSWORD = utils.GetValue("DB_PASSWORD")
		DB_HOST     = utils.GetValue("DB_HOST")
		DB_OPTIONS  = utils.GetValue("DB_OPTIONS")
	)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	uri := fmt.Sprintf(
		"%s://%s:%s@%s/%s",
		DB_PROTOCOL,
		DB_USER,
		DB_PASSWORD,
		DB_HOST,
		DB_OPTIONS,
	)

	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

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

func Disconnect(ctx context.Context) error {
	if err := DB.Client.Disconnect(ctx); err != nil {
		log.Fatalf("error when disconnecting the database: %v\n", err)
		return err
	}

	return nil
}

func GetCollection(name string) *mongo.Collection {
	return DB.Database.Collection(name)
}
