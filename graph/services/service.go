package services

import "go-cms-gql/graph/model"

type UserService interface {
	Register(input model.NewUser) (*model.User, error)
	Login(input model.LoginInput) (string, error)
	GetUserInfo(userID string) (*model.User, error)
}

// TODO: implement content service
type ContentService interface {
}
