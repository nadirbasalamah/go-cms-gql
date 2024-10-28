package graph

import (
	"context"
	"errors"
	"go-cms-gql/graph/model"
	"go-cms-gql/utils"
)

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context) (*model.UserData, error) {
	user := utils.ForContext(ctx)
	if user == nil {
		return nil, errors.New("access denied")
	}

	return utils.ConvertToUserData(user), nil
}

// Categories is the resolver for the categories field.
func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	return r.categoryService.GetAll(ctx)
}

// Content is the resolver for the content field.
func (r *queryResolver) Category(ctx context.Context, id string) (*model.Category, error) {
	return r.categoryService.GetByID(ctx, id)
}

// Contents is the resolver for the contents field.
func (r *queryResolver) Contents(ctx context.Context, keyword *string) ([]*model.Content, error) {
	var actualKeyword string

	if keyword != nil {
		actualKeyword = *keyword
	}

	return r.contentService.GetAll(ctx, actualKeyword)
}

// Content is the resolver for the content field.
func (r *queryResolver) Content(ctx context.Context, id string) (*model.Content, error) {
	return r.contentService.GetByID(ctx, id)
}

// ContentsByCategory is the resolver for the contentsByCategory field.
func (r *queryResolver) ContentsByCategory(ctx context.Context, categoryID string) ([]*model.Content, error) {
	return r.contentService.GetByCategoryID(ctx, categoryID)
}

// ContentsByUser is the resolver for the contentsByUser field.
func (r *queryResolver) ContentsByUser(ctx context.Context) ([]*model.Content, error) {
	return r.contentService.GetByUser(ctx)
}

// Tags is the resolver for the tags field.
func (r *queryResolver) Tags(ctx context.Context, input model.GetTag) ([]string, error) {
	return r.recommendationService.GetTags(ctx, input)
}

// GenerateContent is the resolver for the generateContent field.
func (r *queryResolver) GenerateContent(ctx context.Context, generateInput model.GenerateContent) (string, error) {
	return r.recommendationService.GenerateContent(ctx, generateInput)
}
