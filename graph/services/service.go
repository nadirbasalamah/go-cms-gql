package services

import (
	"context"
	"go-cms-gql/graph/model"
)

type UserService interface {
	Register(ctx context.Context, input model.NewUser) (*model.User, error)
	Login(ctx context.Context, input model.LoginInput) (string, error)
	GetUserInfo(ctx context.Context, userID string) (*model.User, error)
}

type CategoryService interface {
	GetAll(ctx context.Context) ([]*model.Category, error)
	GetByID(ctx context.Context, categoryID string) (*model.Category, error)
	Create(ctx context.Context, input model.NewCategory) (*model.Category, error)
	Update(ctx context.Context, input model.EditCategory) (*model.Category, error)
	Delete(ctx context.Context, input model.DeleteCategory) (bool, error)
}

type ContentService interface {
	GetAll(ctx context.Context, keyword string) ([]*model.Content, error)
	GetByID(ctx context.Context, contentID string) (*model.Content, error)
	GetByCategoryID(ctx context.Context, categoryID string) ([]*model.Content, error)
	Create(ctx context.Context, input model.NewContent, user model.User) (*model.Content, error)
	Update(ctx context.Context, input model.EditContent, user model.User) (*model.Content, error)
	Delete(ctx context.Context, input model.DeleteContent, user model.User) (bool, error)
}

type RecommendationService interface {
	GenerateContent(ctx context.Context, generateInput model.GenerateContent) (string, error)
	GetTags(ctx context.Context, input model.GetTag) ([]string, error)
}
