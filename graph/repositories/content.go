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

func (cr *ContentRepositoryImpl) GetAll(ctx context.Context, keyword string) ([]*model.Content, error) {
	var query primitive.D = bson.D{{}}

	if keyword != "" {
		query = bson.D{
			{Key: "$or", Value: bson.A{
				bson.D{{Key: "title", Value: primitive.Regex{Pattern: keyword, Options: "i"}}},
				bson.D{{Key: "description", Value: primitive.Regex{Pattern: keyword, Options: "i"}}},
			}},
		}
	}

	var findOptions *options.FindOptions = options.Find()
	findOptions.SetSort(bson.D{{Key: "createdAt", Value: -1}})

	cursor, err := database.GetCollection(contentCollection).Find(ctx, query, findOptions)
	if err != nil {
		return nil, errors.New("error occurred when fetching contents")
	}

	var contents []*model.Content = make([]*model.Content, 0)

	if err := cursor.All(ctx, &contents); err != nil {
		return nil, errors.New("error occurred when fetching contents")
	}

	return contents, nil
}

func (cr *ContentRepositoryImpl) GetByID(ctx context.Context, contentID string) (*model.Content, error) {
	cID, err := primitive.ObjectIDFromHex(contentID)
	if err != nil {
		return nil, errors.New("id is invalid")
	}

	var query primitive.D = bson.D{{Key: "_id", Value: cID}}
	var collection *mongo.Collection = database.GetCollection(contentCollection)

	var contentData *mongo.SingleResult = collection.FindOne(ctx, query)

	if contentData.Err() != nil {
		return nil, errors.New("content not found")
	}

	var content *model.Content = &model.Content{}

	if err := contentData.Decode(content); err != nil {
		return nil, errors.New("error occurred when fetching content")
	}

	return content, nil
}

func (cr *ContentRepositoryImpl) GetByCategoryID(ctx context.Context, categoryID string) ([]*model.Content, error) {
	_, err := primitive.ObjectIDFromHex(categoryID)
	if err != nil {
		return nil, errors.New("id is invalid")
	}

	var query primitive.D = bson.D{{Key: "category._id", Value: categoryID}}

	var findOptions *options.FindOptions = options.Find()
	findOptions.SetSort(bson.D{{Key: "createdAt", Value: -1}})

	cursor, err := database.GetCollection(contentCollection).Find(ctx, query, findOptions)
	if err != nil {
		return nil, errors.New("error occurred when fetching contents")
	}

	var contents []*model.Content = make([]*model.Content, 0)

	if err := cursor.All(ctx, &contents); err != nil {
		return nil, errors.New("error occurred when fetching contents")
	}

	return contents, nil
}

func (cr *ContentRepositoryImpl) GetByUser(ctx context.Context) ([]*model.Content, error) {
	user, err := utils.GetAuthenticatedUser(ctx)

	if err != nil {
		return nil, err
	}

	var query primitive.D = bson.D{{Key: "author._id", Value: &user.ID}}

	var findOptions *options.FindOptions = options.Find()
	findOptions.SetSort(bson.D{{Key: "createdAt", Value: -1}})

	cursor, err := database.GetCollection(contentCollection).Find(ctx, query, findOptions)
	if err != nil {
		return nil, errors.New("error occurred when fetching contents")
	}

	var contents []*model.Content = make([]*model.Content, 0)

	if err := cursor.All(ctx, &contents); err != nil {
		return nil, errors.New("error occurred when fetching contents")
	}

	return contents, nil
}

func (cr *ContentRepositoryImpl) Create(ctx context.Context, input model.NewContent) (*model.Content, error) {
	category, err := cr.categoryRepo.GetByID(ctx, input.CategoryID)

	if err != nil {
		return nil, errors.New("create content failed, category not found")
	}

	user, err := utils.GetAuthenticatedUser(ctx)

	if err != nil {
		return nil, err
	}

	var content model.Content = model.Content{
		Title:     input.Title,
		Content:   input.Content,
		Author:    utils.ConvertToUserData(user),
		Category:  category,
		CreatedAt: time.Now(),
	}

	var collection *mongo.Collection = database.GetCollection(contentCollection)

	result, err := collection.InsertOne(ctx, content)

	if err != nil {
		return nil, errors.New("create content failed")
	}

	var filter primitive.D = bson.D{{Key: "_id", Value: result.InsertedID}}
	var createdRecord *mongo.SingleResult = collection.FindOne(ctx, filter)

	var createdContent *model.Content = &model.Content{}

	if err := createdRecord.Decode(createdContent); err != nil {
		return nil, errors.New("error occurred when fetching created content")
	}

	return createdContent, nil
}

func (cr *ContentRepositoryImpl) Update(ctx context.Context, input model.EditContent) (*model.Content, error) {
	category, err := cr.categoryRepo.GetByID(ctx, input.CategoryID)

	if err != nil {
		return nil, errors.New("create content failed, category not found")
	}

	cID, err := primitive.ObjectIDFromHex(input.ContentID)
	if err != nil {
		return nil, errors.New("id is invalid")
	}

	user, err := utils.GetAuthenticatedUser(ctx)

	if err != nil {
		return nil, err
	}

	var query primitive.D = bson.D{
		{Key: "_id", Value: cID},
		{Key: "author._id", Value: &user.ID},
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
		ctx,
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
		return nil, errors.New("error occurred when fetching updated content")
	}

	return editedContent, nil

}

func (cr *ContentRepositoryImpl) Delete(ctx context.Context, input model.DeleteContent) (bool, error) {
	cID, err := primitive.ObjectIDFromHex(input.ContentID)
	if err != nil {
		return false, errors.New("id is invalid")
	}

	user, err := utils.GetAuthenticatedUser(ctx)

	if err != nil {
		return false, err
	}

	var query primitive.D = bson.D{
		{Key: "_id", Value: cID},
		{Key: "author._id", Value: &user.ID},
	}

	var collection *mongo.Collection = database.GetCollection(contentCollection)

	result, err := collection.DeleteOne(ctx, query)
	var isFailed bool = err != nil || result.DeletedCount < 1

	if isFailed {
		return !isFailed, errors.New("delete content failed")
	}

	return true, nil
}
