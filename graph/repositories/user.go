package repositories

import (
	"context"
	"errors"
	"go-cms-gql/database"
	"go-cms-gql/graph/model"
	"go-cms-gql/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserRepositoryImpl struct {
}

const userCollection = utils.USER_COLLECTION

func InitUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (ur *UserRepositoryImpl) Register(ctx context.Context, input model.NewUser) (*model.User, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("error occurred when creating password")
	}

	var password string = string(bs)

	var newUser model.User = model.User{
		Username:  input.Username,
		Email:     input.Email,
		Password:  password,
		Role:      utils.USER_ROLE,
		CreatedAt: time.Now(),
	}

	var collection *mongo.Collection = database.GetCollection(userCollection)

	var foundUser *model.User = &model.User{}
	userFilter := bson.M{"email": input.Email}

	err = collection.FindOne(ctx, userFilter).Decode(foundUser)

	if err == nil {
		return nil, errors.New("email already exists")
	} else if err != mongo.ErrNoDocuments {
		return nil, errors.New("error occurred when fetching document")
	}

	res, err := collection.InsertOne(ctx, newUser)

	if err != nil {
		return nil, errors.New("registration failed")
	}

	var user *model.User = &model.User{}
	var filter primitive.D = bson.D{{Key: "_id", Value: res.InsertedID}}

	var userRecord *mongo.SingleResult = collection.FindOne(ctx, filter)
	if err := userRecord.Decode(user); err != nil {
		return nil, errors.New("error occurred when fetching user")
	}

	return user, nil
}

func (ur *UserRepositoryImpl) GetUserByEmail(ctx context.Context, input model.LoginInput) (*model.User, error) {
	var collection *mongo.Collection = database.GetCollection(userCollection)

	var user *model.User = &model.User{}
	filter := bson.M{"email": input.Email}

	var res *mongo.SingleResult = collection.FindOne(ctx, filter)
	if err := res.Decode(user); err != nil {
		return nil, errors.New("invalid email")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	return user, nil
}

func (ur *UserRepositoryImpl) GetUserInfo(ctx context.Context, userID string) (*model.User, error) {
	uID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, errors.New("id is invalid")
	}

	var query primitive.D = bson.D{{Key: "_id", Value: uID}}
	var collection *mongo.Collection = database.GetCollection(userCollection)

	var userData *mongo.SingleResult = collection.FindOne(ctx, query)

	if userData.Err() != nil {
		return nil, errors.New("user not found")
	}

	var user *model.User = &model.User{}
	if err := userData.Decode(user); err != nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}
