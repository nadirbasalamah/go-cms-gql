package main

import (
	"context"
	"go-cms-gql/database"
	"go-cms-gql/graph/model"
	"go-cms-gql/utils"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

const userCollection = utils.USER_COLLECTION

func main() {
	connectToDB()

	ctx := context.TODO()
	input := model.NewUser{
		Username: utils.GetValue("ADMIN_NAME"),
		Email:    utils.GetValue("ADMIN_EMAIL"),
		Password: utils.GetValue("ADMIN_PASSWORD"),
	}

	generateAdmin(ctx, input)
}

func connectToDB() {
	err := database.Connect(utils.GetValue("DATABASE_NAME"))
	if err != nil {
		log.Fatalf("Cannot connect to the database: %v\n", err)
	}
}

func generateAdmin(ctx context.Context, input model.NewUser) {
	//TODO: add request validation

	bs, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("error occurred when creating password: %v\n", err)
	}

	var password string = string(bs)

	var newUser model.User = model.User{
		Username:  input.Username,
		Email:     input.Email,
		Password:  password,
		Role:      utils.ADMIN_ROLE,
		CreatedAt: time.Now(),
	}

	var collection *mongo.Collection = database.GetCollection(userCollection)

	var foundUser *model.User = &model.User{}
	userFilter := bson.M{"role": utils.ADMIN_ROLE}

	err = collection.FindOne(ctx, userFilter).Decode(foundUser)

	if err == nil {
		log.Fatal("admin already exists")
	} else if err != mongo.ErrNoDocuments {
		log.Fatalf("error occurred when fetching document: %v\n", err)
	}

	res, err := collection.InsertOne(ctx, newUser)

	if err != nil || res.InsertedID == nil {
		log.Fatalf("admin creation failed: %v\n", err)
	}
}
