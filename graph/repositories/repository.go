package repositories

import "go-cms-gql/graph/model"

type UserRepository interface {
	Register(input model.NewUser) (*model.User, error)
	GetUserByEmail(input model.LoginInput) (*model.User, error)
	GetUserInfo(userID string) (*model.User, error)
}

type CategoryRepository interface {
	GetAll() ([]*model.Category, error)
	GetByID(categoryID string) (*model.Category, error)
	Create(input model.NewCategory) (*model.Category, error)
	Update(input model.EditCategory) (*model.Category, error)
	Delete(input model.DeleteCategory) (bool, error)
}

type ContentRepository interface {
	GetAll() ([]*model.Content, error)
	GetByID(contentID string) (*model.Content, error)
	Create(input model.NewContent, user model.User) (*model.Content, error)
	Update(input model.EditContent, user model.User) (*model.Content, error)
	Delete(input model.DeleteContent, user model.User) (bool, error)
}
