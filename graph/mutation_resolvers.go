package graph

import (
	"context"
	"errors"
	"go-cms-gql/graph/model"
	"go-cms-gql/utils"
)

// User Resolvers

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input model.NewUser) (*model.UserData, error) {
	user, err := r.userService.Register(ctx, input)

	if err != nil {
		return nil, err
	}

	return utils.ConvertToUserData(user), nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (string, error) {
	return r.userService.Login(ctx, input)
}

// Category Resolvers

// NewCategory is the resolver for the newCategory field.
func (r *mutationResolver) NewCategory(ctx context.Context, input model.NewCategory) (*model.Category, error) {
	return r.categoryService.Create(ctx, input)
}

// EditCategory is the resolver for the editCategory field.
func (r *mutationResolver) EditCategory(ctx context.Context, input model.EditCategory) (*model.Category, error) {
	return r.categoryService.Update(ctx, input)
}

// DeleteCategory is the resolver for the deleteCategory field.
func (r *mutationResolver) DeleteCategory(ctx context.Context, input model.DeleteCategory) (bool, error) {
	return r.categoryService.Delete(ctx, input)
}

// Content Resolvers

// NewContent is the resolver for the newContent field.
func (r *mutationResolver) NewContent(ctx context.Context, input model.NewContent) (*model.Content, error) {
	user, err := utils.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, errors.New("access denied")
	}

	return r.contentService.Create(ctx, input, *user)
}

// EditContent is the resolver for the editContent field.
func (r *mutationResolver) EditContent(ctx context.Context, input model.EditContent) (*model.Content, error) {
	user, err := utils.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, errors.New("access denied")
	}

	return r.contentService.Update(ctx, input, *user)
}

// DeleteContent is the resolver for the deleteContent field.
func (r *mutationResolver) DeleteContent(ctx context.Context, input model.DeleteContent) (bool, error) {
	user, err := utils.GetAuthenticatedUser(ctx)
	if err != nil {
		return false, errors.New("access denied")
	}

	return r.contentService.Delete(ctx, input, *user)
}
