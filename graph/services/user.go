package services

import (
	"context"
	"go-cms-gql/graph/model"
	"go-cms-gql/graph/repositories"
)

type UserServiceImpl struct {
	repository     repositories.UserRepository
	tokenGenerator func(userId string) (string, error)
}

func InitUserService(userRepository repositories.UserRepository, tokenGenFunc func(userId string) (string, error)) UserService {
	return &UserServiceImpl{
		repository:     userRepository,
		tokenGenerator: tokenGenFunc,
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

	token, err := us.tokenGenerator(user.ID)

	if err != nil {
		return "", err
	}

	return token, nil

}

func (us *UserServiceImpl) GetUserInfo(ctx context.Context, userID string) (*model.User, error) {
	return us.repository.GetUserInfo(ctx, userID)
}
