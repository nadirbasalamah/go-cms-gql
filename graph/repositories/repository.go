package repositories

import "go-cms-gql/graph/model"

type UserRepository interface {
	Register(input model.NewUser) (*model.User, error)
	GetUserByEmail(input model.LoginInput) (*model.User, error)
	GetUserInfo(userID string) (*model.User, error)
}

//TODO: implement content repository
type ContentRepository interface {
}
