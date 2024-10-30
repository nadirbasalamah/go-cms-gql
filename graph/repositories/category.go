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

type CategoryRepositoryImpl struct{}

const categoryCollection = utils.CATEGORY_COLLECTION

func InitCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (cr *CategoryRepositoryImpl) GetAll(ctx context.Context) ([]*model.Category, error) {
	var query primitive.D = bson.D{{}}

	var findOptions *options.FindOptions = options.Find()
	findOptions.SetSort(bson.D{{Key: "createdAt", Value: -1}})

	cursor, err := database.GetCollection(categoryCollection).Find(ctx, query, findOptions)
	if err != nil {
		return nil, errors.New("error occurred when fetching categories")
	}

	var categories []*model.Category = make([]*model.Category, 0)

	if err := cursor.All(ctx, &categories); err != nil {
		return nil, errors.New("error occurred when fetching categories")
	}

	return categories, nil
}

func (cr *CategoryRepositoryImpl) GetByID(ctx context.Context, categoryID string) (*model.Category, error) {
	cID, err := primitive.ObjectIDFromHex(categoryID)
	if err != nil {
		return nil, errors.New("id is invalid")
	}

	var query primitive.D = bson.D{{Key: "_id", Value: cID}}
	var collection *mongo.Collection = database.GetCollection(categoryCollection)

	var categoryData *mongo.SingleResult = collection.FindOne(ctx, query)

	if categoryData.Err() != nil {
		return nil, errors.New("category not found")
	}

	var category *model.Category = &model.Category{}

	if err := categoryData.Decode(category); err != nil {
		return nil, errors.New("error occurred when fetching category")
	}

	return category, nil
}

func (cr *CategoryRepositoryImpl) Create(ctx context.Context, input model.NewCategory) (*model.Category, error) {
	var category model.Category = model.Category{
		Title:     input.Title,
		CreatedAt: time.Now(),
	}

	var collection *mongo.Collection = database.GetCollection(categoryCollection)

	var foundCategory *model.Category = &model.Category{}
	categoryFilter := bson.M{"title": input.Title}

	err := collection.FindOne(ctx, categoryFilter).Decode(foundCategory)

	if err == nil {
		return nil, errors.New("category already exists")
	} else if err != mongo.ErrNoDocuments {
		return nil, errors.New("error occurred when fetching document")
	}

	result, err := collection.InsertOne(ctx, category)

	if err != nil {
		return nil, errors.New("create category failed")
	}

	var filter primitive.D = bson.D{{Key: "_id", Value: result.InsertedID}}
	var createdRecord *mongo.SingleResult = collection.FindOne(ctx, filter)

	var createdCategory *model.Category = &model.Category{}

	if err := createdRecord.Decode(createdCategory); err != nil {
		return nil, errors.New("error occurred when fetching created category")
	}

	return createdCategory, nil
}

func (cr *CategoryRepositoryImpl) Update(ctx context.Context, input model.EditCategory) (*model.Category, error) {
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
		ctx,
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
		return nil, errors.New("error occurred when fetching updated category")
	}

	return editedCategory, nil
}

func (cr *CategoryRepositoryImpl) Delete(ctx context.Context, input model.DeleteCategory) (bool, error) {
	cID, err := primitive.ObjectIDFromHex(input.CategoryID)
	if err != nil {
		return false, errors.New("id is invalid")
	}

	var query primitive.D = bson.D{
		{Key: "_id", Value: cID},
	}

	isUsed, err := isCategoryInUse(ctx, input.CategoryID)

	if err != nil {
		return false, err
	}

	if isUsed {
		return false, errors.New("delete category failed, category is used")
	}

	var collection *mongo.Collection = database.GetCollection(categoryCollection)

	result, err := collection.DeleteOne(ctx, query)
	var isFailed bool = err != nil || result.DeletedCount < 1

	if isFailed {
		return !isFailed, errors.New("delete category failed, category not found")
	}

	return true, nil
}

func isCategoryInUse(ctx context.Context, categoryID string) (bool, error) {
	var query primitive.D = bson.D{{Key: "category._id", Value: categoryID}}
	var collection *mongo.Collection = database.GetCollection(contentCollection)

	count, err := collection.CountDocuments(ctx, query)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
