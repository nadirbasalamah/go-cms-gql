package graph

import (
	"context"
	"errors"
	"fmt"
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

// NewContent is the resolver for the newContent field.
func (r *mutationResolver) NewContent(ctx context.Context, input model.NewContent) (*model.Content, error) {
	panic(fmt.Errorf("not implemented: NewContent - newContent"))
}

// EditContent is the resolver for the editContent field.
func (r *mutationResolver) EditContent(ctx context.Context, input model.EditContent) (*model.Content, error) {
	panic(fmt.Errorf("not implemented: EditContent - editContent"))
}

// DeleteContent is the resolver for the deleteContent field.
func (r *mutationResolver) DeleteContent(ctx context.Context, input model.DeleteContent) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteContent - deleteContent"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	user := middlewares.ForContext(ctx)
	if user == nil {
		return nil, errors.New("access denied")
	}

	return user, nil
}

// Contents is the resolver for the contents field.
func (r *queryResolver) Contents(ctx context.Context) ([]*model.Content, error) {
	panic(fmt.Errorf("not implemented: Contents - contents"))
}

// Content is the resolver for the content field.
func (r *queryResolver) Content(ctx context.Context, id string) (*model.Content, error) {
	panic(fmt.Errorf("not implemented: Content - content"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver {
	r.userService = services.InitUserService()
	return &mutationResolver{r}
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
