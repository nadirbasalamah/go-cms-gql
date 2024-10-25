package repositories

import (
	"context"
	"go-cms-gql/graph/model"
)

type UserRepository interface {
	Register(ctx context.Context, input model.NewUser) (*model.User, error)
	GetUserByEmail(ctx context.Context, input model.LoginInput) (*model.User, error)
	GetUserInfo(ctx context.Context, userID string) (*model.User, error)
}

type CategoryRepository interface {
	GetAll(ctx context.Context) ([]*model.Category, error)
	GetByID(ctx context.Context, categoryID string) (*model.Category, error)
	Create(ctx context.Context, input model.NewCategory) (*model.Category, error)
	Update(ctx context.Context, input model.EditCategory) (*model.Category, error)
	Delete(ctx context.Context, input model.DeleteCategory) (bool, error)
}

type ContentRepository interface {
	GetAll(ctx context.Context, keyword string) ([]*model.Content, error)
	GetByID(ctx context.Context, contentID string) (*model.Content, error)
	GetByCategoryID(ctx context.Context, categoryID string) ([]*model.Content, error)
	GetByUser(ctx context.Context) ([]*model.Content, error)
	Create(ctx context.Context, input model.NewContent, user model.User) (*model.Content, error)
	Update(ctx context.Context, input model.EditContent, user model.User) (*model.Content, error)
	Delete(ctx context.Context, input model.DeleteContent, user model.User) (bool, error)
}
