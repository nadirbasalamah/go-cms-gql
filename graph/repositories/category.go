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

type CategoryRepositoryImpl struct {
}

const categoryCollection = utils.CATEGORY_COLLECTION

func InitCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (cr *CategoryRepositoryImpl) GetAll() ([]*model.Category, error) {
	var query primitive.D = bson.D{{}}

	var findOptions *options.FindOptions = options.Find()
	findOptions.SetSort(bson.D{{Key: "createdAt", Value: -1}})

	cursor, err := database.GetCollection(categoryCollection).Find(context.TODO(), query, findOptions)
	if err != nil {
		return nil, err
	}

	var categories []*model.Category = make([]*model.Category, 0)

	if err := cursor.All(context.TODO(), &categories); err != nil {
		return nil, err
	}

	return categories, nil
}

func (cr *CategoryRepositoryImpl) GetByID(categoryID string) (*model.Category, error) {
	cID, err := primitive.ObjectIDFromHex(categoryID)
	if err != nil {
		return nil, errors.New("id is invalid")
	}

	var query primitive.D = bson.D{{Key: "_id", Value: cID}}
	var collection *mongo.Collection = database.GetCollection(categoryCollection)

	var categoryData *mongo.SingleResult = collection.FindOne(context.TODO(), query)

	if categoryData.Err() != nil {
		return nil, errors.New("category not found")
	}

	var category *model.Category = &model.Category{}

	if err := categoryData.Decode(category); err != nil {
		return nil, err
	}

	return category, nil
}

func (cr *CategoryRepositoryImpl) Create(input model.NewCategory) (*model.Category, error) {
	var category model.Category = model.Category{
		Title:     input.Title,
		CreatedAt: time.Now(),
	}

	var collection *mongo.Collection = database.GetCollection(categoryCollection)

	result, err := collection.InsertOne(context.TODO(), category)

	if err != nil {
		return nil, errors.New("create category failed")
	}

	var filter primitive.D = bson.D{{Key: "_id", Value: result.InsertedID}}
	var createdRecord *mongo.SingleResult = collection.FindOne(context.TODO(), filter)

	var createdCategory *model.Category = &model.Category{}

	if err := createdRecord.Decode(createdCategory); err != nil {
		return nil, err
	}

	return createdCategory, nil
}

func (cr *CategoryRepositoryImpl) Update(input model.EditCategory) (*model.Category, error) {
	cID, err := primitive.ObjectIDFromHex(input.CategoryID)
	if err != nil {
		return nil, errors.New("id is invalid")
	}

	var query primitive.D = bson.D{
		{Key: "_id", Value: cID},
	}
	var update primitive.D = bson.D{{
		Key: "$set",
		Value: bson.D{
			{Key: "title", Value: input.Title},
			{Key: "updatedAt", Value: time.Now()},
		},
	}}

	var collection *mongo.Collection = database.GetCollection(categoryCollection)

	var updateResult *mongo.SingleResult = collection.FindOneAndUpdate(
		context.TODO(),
		query,
		update,
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	)

	if updateResult.Err() != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("category not found")
		}
		return nil, errors.New("update category failed")
	}

	var editedCategory *model.Category = &model.Category{}

	if err := updateResult.Decode(editedCategory); err != nil {
		return nil, err
	}

	return editedCategory, nil
}

func (cr *CategoryRepositoryImpl) Delete(input model.DeleteCategory) (bool, error) {
	cID, err := primitive.ObjectIDFromHex(input.CategoryID)
	if err != nil {
		return false, err
	}

	var query primitive.D = bson.D{
		{Key: "_id", Value: cID},
	}

	var collection *mongo.Collection = database.GetCollection(categoryCollection)

	result, err := collection.DeleteOne(context.TODO(), query)
	var isFailed bool = err != nil || result.DeletedCount < 1

	if isFailed {
		return !isFailed, err
	}

	return true, nil
}
