package services

import "go-cms-gql/graph/model"

type UserService interface {
	Register(input model.NewUser) (*model.User, error)
	Login(input model.LoginInput) (string, error)
	GetUserInfo(userID string) (*model.User, error)
}

type CategoryService interface {
	GetAll() ([]*model.Category, error)
	GetByID(categoryID string) (*model.Category, error)
	Create(input model.NewCategory) (*model.Category, error)
	Update(input model.EditCategory) (*model.Category, error)
	Delete(input model.DeleteCategory) (bool, error)
}

// TODO: implement content service
type ContentService interface {
}
