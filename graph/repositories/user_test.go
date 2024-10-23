package repositories_test

import (
	"errors"
	"go-cms-gql/graph/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	t.Run("Register | Valid", func(t *testing.T) {
		userRepository.On("Register", ctx, model.NewUser{}).Return(&model.User{}, nil).Once()

		result, err := userService.Register(ctx, model.NewUser{})

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Register | Invalid", func(t *testing.T) {
		userRepository.On("Register", ctx, model.NewUser{}).Return(&model.User{}, errors.New("something went wrong")).Once()

		result, err := userService.Register(ctx, model.NewUser{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetUserByEmail(t *testing.T) {
	t.Run("GetUserByEmail | Valid", func(t *testing.T) {
		userRepository.On("GetUserByEmail", ctx, model.LoginInput{}).Return(&model.User{}, nil).Once()

		result, err := userService.Login(ctx, model.LoginInput{})

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetUserByEmail | Invalid", func(t *testing.T) {
		userRepository.On("GetUserByEmail", ctx, model.LoginInput{}).Return(&model.User{}, errors.New("something went wrong")).Once()

		result, err := userService.Login(ctx, model.LoginInput{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetUserInfo(t *testing.T) {
	t.Run("GetUserInfo | Valid", func(t *testing.T) {
		userRepository.On("GetUserInfo", ctx, "1").Return(&model.User{}, nil).Once()

		result, err := userService.GetUserInfo(ctx, "1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetUserInfo | Invalid", func(t *testing.T) {
		userRepository.On("GetUserInfo", ctx, "0").Return(&model.User{}, errors.New("something went wrong")).Once()

		result, err := userService.GetUserInfo(ctx, "0")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}
