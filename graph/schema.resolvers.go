package graph

import (
	"context"
	"errors"
	"go-cms-gql/graph/middlewares"
	"go-cms-gql/graph/model"
	"go-cms-gql/graph/services"
)

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input model.NewUser) (*model.User, error) {
	user, err := r.userService.Register(input)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (string, error) {
	token, err := r.userService.Login(input)

	if err != nil {
		return "", err
	}

	return token, nil
}

// NewCategory is the resolver for the newCategory field.
func (r *mutationResolver) NewCategory(ctx context.Context, input model.NewCategory) (*model.Category, error) {
	user := middlewares.ForContext(ctx)
	if user == nil {
		return nil, errors.New("access denied")
	}

	category, err := r.categoryService.Create(input)

	if err != nil {
		return nil, err
	}

	return category, nil
}

// EditCategory is the resolver for the editCategory field.
func (r *mutationResolver) EditCategory(ctx context.Context, input model.EditCategory) (*model.Category, error) {
	user := middlewares.ForContext(ctx)
	if user == nil {
		return nil, errors.New("access denied")
	}

	category, err := r.categoryService.Update(input)

	if err != nil {
		return nil, err
	}

	return category, nil
}

// DeleteCategory is the resolver for the deleteCategory field.
func (r *mutationResolver) DeleteCategory(ctx context.Context, input model.DeleteCategory) (bool, error) {
	user := middlewares.ForContext(ctx)
	if user == nil {
		return false, errors.New("access denied")
	}

	result, err := r.categoryService.Delete(input)

	if err != nil {
		return result, err
	}

	return result, nil
}

// NewContent is the resolver for the newContent field.
func (r *mutationResolver) NewContent(ctx context.Context, input model.NewContent) (*model.Content, error) {
	user := middlewares.ForContext(ctx)
	if user == nil {
		return nil, errors.New("access denied")
	}

	content, err := r.contentService.Create(input, *user)

	if err != nil {
		return nil, err
	}

	return content, nil
}

// EditContent is the resolver for the editContent field.
func (r *mutationResolver) EditContent(ctx context.Context, input model.EditContent) (*model.Content, error) {
	user := middlewares.ForContext(ctx)
	if user == nil {
		return nil, errors.New("access denied")
	}

	content, err := r.contentService.Update(input, *user)

	if err != nil {
		return nil, err
	}

	return content, nil
}

// DeleteContent is the resolver for the deleteContent field.
func (r *mutationResolver) DeleteContent(ctx context.Context, input model.DeleteContent) (bool, error) {
	user := middlewares.ForContext(ctx)
	if user == nil {
		return false, errors.New("access denied")
	}

	result, err := r.contentService.Delete(input, *user)

	if err != nil {
		return result, err
	}

	return result, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	user := middlewares.ForContext(ctx)
	if user == nil {
		return nil, errors.New("access denied")
	}

	return user, nil
}

// Categories is the resolver for the categories field.
func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	user := middlewares.ForContext(ctx)
	if user == nil {
		return nil, errors.New("access denied")
	}

	categories, err := r.categoryService.GetAll()

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

	category, err := r.categoryService.GetByID(id)

	if err != nil {
		return nil, err
	}

	return category, nil
}

// Contents is the resolver for the contents field.
func (r *queryResolver) Contents(ctx context.Context) ([]*model.Content, error) {
	user := middlewares.ForContext(ctx)
	if user == nil {
		return nil, errors.New("access denied")
	}

	contents, err := r.contentService.GetAll()

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

	content, err := r.contentService.GetByID(id)

	if err != nil {
		return nil, err
	}

	return content, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver {
	r.userService = services.InitUserService()
	r.categoryService = services.InitCategoryService()
	r.contentService = services.InitContentService()

	return &mutationResolver{r}
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
