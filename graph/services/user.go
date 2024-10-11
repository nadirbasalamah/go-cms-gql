package services

import (
	"go-cms-gql/graph/model"
	"go-cms-gql/graph/repositories"
	"go-cms-gql/utils"
)

type UserServiceImpl struct {
	repository repositories.UserRepository
}

func InitUserService() UserService {
	return &UserServiceImpl{
		repository: repositories.InitUserRepository(),
	}
}

func (us *UserServiceImpl) Register(input model.NewUser) (*model.User, error) {
	return us.repository.Register(input)
}

func (us *UserServiceImpl) Login(input model.LoginInput) (string, error) {
	user, err := us.repository.GetUserByEmail(input)

	if err != nil {
		return "", err
	}

	token, err := utils.GenerateNewAccessToken(user.ID)

	if err != nil {
		return "", err
	}

	return token, nil

}

func (us *UserServiceImpl) GetUserInfo(userID string) (*model.User, error) {
	return us.repository.GetUserInfo(userID)
}
