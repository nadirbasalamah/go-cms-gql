package graph

import (
	"context"
	"errors"
	"go-cms-gql/graph/middlewares"
	"go-cms-gql/graph/model"
	"go-cms-gql/utils"
)

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context) (*model.UserData, error) {
	user := middlewares.ForContext(ctx)
	if user == nil {
		return nil, errors.New("access denied")
	}

	return utils.ConvertToUserData(user), nil
}

// Categories is the resolver for the categories field.
func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	user := middlewares.ForContext(ctx)
	if user == nil {
		return nil, errors.New("access denied")
	}

	categories, err := r.categoryService.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return categories, nil
}

// Content is the resolver for the content field.
func (r *queryResolver) Category(ctx context.Context, id string) (*model.Category, error) {
	user := middlewares.ForContext(ctx)
	if user == nil {
		return nil, errors.New("access denied")
	}

	category, err := r.categoryService.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return category, nil
}

// Contents is the resolver for the contents field.
func (r *queryResolver) Contents(ctx context.Context, keyword *string) ([]*model.Content, error) {
	user := middlewares.ForContext(ctx)
	if user == nil {
		return nil, errors.New("access denied")
	}

	var actualKeyword string

	if keyword != nil {
		actualKeyword = *keyword
	}

	contents, err := r.contentService.GetAll(ctx, actualKeyword)

	if err != nil {
		return nil, err
	}

	return contents, nil
}

// Content is the resolver for the content field.
func (r *queryResolver) Content(ctx context.Context, id string) (*model.Content, error) {
	user := middlewares.ForContext(ctx)
	if user == nil {
		return nil, errors.New("access denied")
	}

	content, err := r.contentService.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return content, nil
}

// ContentsByCategory is the resolver for the contentsByCategory field.
func (r *queryResolver) ContentsByCategory(ctx context.Context, categoryID string) ([]*model.Content, error) {
	user := middlewares.ForContext(ctx)
	if user == nil {
		return nil, errors.New("access denied")
	}

	contents, err := r.contentService.GetByCategoryID(ctx, categoryID)

	if err != nil {
		return nil, err
	}

	return contents, nil
}

// Tags is the resolver for the tags field.
func (r *queryResolver) Tags(ctx context.Context, content string) ([]string, error) {
	return r.recommendationService.GetTags(ctx, content)
}

// GenerateContent is the resolver for the generateContent field.
func (r *queryResolver) GenerateContent(ctx context.Context, generateInput model.GenerateContent) (string, error) {
	return r.recommendationService.GenerateContent(ctx, generateInput)
}
