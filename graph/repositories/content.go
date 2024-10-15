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
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ContentRepositoryImpl struct {
	categoryRepo CategoryRepository
}

const contentCollection = utils.CONTENT_COLLECTION

func InitContentRepository() ContentRepository {
	return &ContentRepositoryImpl{
		categoryRepo: InitCategoryRepository(),
	}
}

func (cr *ContentRepositoryImpl) GetAll() ([]*model.Content, error) {
	var query primitive.D = bson.D{{}}

	var findOptions *options.FindOptions = options.Find()
	findOptions.SetSort(bson.D{{Key: "createdAt", Value: -1}})

	cursor, err := database.GetCollection(contentCollection).Find(context.TODO(), query, findOptions)
	if err != nil {
		return nil, err
	}

	var contents []*model.Content = make([]*model.Content, 0)

	if err := cursor.All(context.TODO(), &contents); err != nil {
		return nil, err
	}

	return contents, nil
}
func (cr *ContentRepositoryImpl) GetByID(contentID string) (*model.Content, error) {
	cID, err := primitive.ObjectIDFromHex(contentID)
	if err != nil {
		return nil, errors.New("id is invalid")
	}

	var query primitive.D = bson.D{{Key: "_id", Value: cID}}
	var collection *mongo.Collection = database.GetCollection(contentCollection)

	var contentData *mongo.SingleResult = collection.FindOne(context.TODO(), query)

	if contentData.Err() != nil {
		return nil, errors.New("content not found")
	}

	var content *model.Content = &model.Content{}

	if err := contentData.Decode(content); err != nil {
		return nil, err
	}

	return content, nil
}

func (cr *ContentRepositoryImpl) Create(input model.NewContent, user model.User) (*model.Content, error) {
	category, err := cr.categoryRepo.GetByID(input.CategoryID)

	if err != nil {
		return nil, errors.New("create content failed, category not found")
	}

	var content model.Content = model.Content{
		Title:     input.Title,
		Content:   input.Content,
		Author:    utils.ConvertToUserData(&user),
		Category:  category,
		CreatedAt: time.Now(),
	}

	var collection *mongo.Collection = database.GetCollection(contentCollection)

	result, err := collection.InsertOne(context.TODO(), content)

	if err != nil {
		return nil, errors.New("create content failed")
	}

	var filter primitive.D = bson.D{{Key: "_id", Value: result.InsertedID}}
	var createdRecord *mongo.SingleResult = collection.FindOne(context.TODO(), filter)

	var createdContent *model.Content = &model.Content{}

	if err := createdRecord.Decode(createdContent); err != nil {
		return nil, err
	}

	return createdContent, nil
}

func (cr *ContentRepositoryImpl) Update(input model.EditContent, user model.User) (*model.Content, error) {
	category, err := cr.categoryRepo.GetByID(input.CategoryID)

	if err != nil {
		return nil, errors.New("create content failed, category not found")
	}

	cID, err := primitive.ObjectIDFromHex(input.ContentID)
	if err != nil {
		return nil, errors.New("id is invalid")
	}

	var query primitive.D = bson.D{
		{Key: "_id", Value: cID},
		{Key: "author._id", Value: user.ID},
	}
	var update primitive.D = bson.D{{
		Key: "$set",
		Value: bson.D{
			{Key: "title", Value: input.Title},
			{Key: "content", Value: input.Content},
			{Key: "category", Value: *category},
			{Key: "updatedAt", Value: time.Now()},
		},
	}}

	var collection *mongo.Collection = database.GetCollection(contentCollection)

	var updateResult *mongo.SingleResult = collection.FindOneAndUpdate(
		context.TODO(),
		query,
		update,
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	)

	if updateResult.Err() != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("content not found")
		}
		return nil, errors.New("update content failed")
	}

	var editedContent *model.Content = &model.Content{}

	if err := updateResult.Decode(editedContent); err != nil {
		return nil, err
	}

	return editedContent, nil

}

func (cr *ContentRepositoryImpl) Delete(input model.DeleteContent, user model.User) (bool, error) {
	cID, err := primitive.ObjectIDFromHex(input.ContentID)
	if err != nil {
		return false, err
	}

	var query primitive.D = bson.D{
		{Key: "_id", Value: cID},
		{Key: "author._id", Value: user.ID},
	}

	var collection *mongo.Collection = database.GetCollection(contentCollection)

	result, err := collection.DeleteOne(context.TODO(), query)
	var isFailed bool = err != nil || result.DeletedCount < 1

	if isFailed {
		return !isFailed, err
	}

	return true, nil
}
