package services

import (
	"context"
	"go-cms-gql/graph/model"
	"go-cms-gql/graph/repositories"
	"go-cms-gql/utils"
)

type UserServiceImpl struct {
	repository repositories.UserRepository
}

func InitUserService(userRepository repositories.UserRepository) UserService {
	return &UserServiceImpl{
		repository: userRepository,
	}
}

func (us *UserServiceImpl) Register(ctx context.Context, input model.NewUser) (*model.User, error) {
	return us.repository.Register(ctx, input)
}

func (us *UserServiceImpl) Login(ctx context.Context, input model.LoginInput) (string, error) {
	user, err := us.repository.GetUserByEmail(ctx, input)

	if err != nil {
		return "", err
	}

	token, err := utils.GenerateNewAccessToken(user.ID)

	if err != nil {
		return "", err
	}

	return token, nil

}

func (us *UserServiceImpl) GetUserInfo(ctx context.Context, userID string) (*model.User, error) {
	return us.repository.GetUserInfo(ctx, userID)
}
