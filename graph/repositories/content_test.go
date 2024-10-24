package repositories_test

import (
	"errors"
	"go-cms-gql/graph/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	t.Run("GetAll | Valid", func(t *testing.T) {
		contentRepository.On("GetAll", ctx, "").Return([]*model.Content{}, nil).Once()

		result, err := contentService.GetAll(ctx, "")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetAll | Invalid", func(t *testing.T) {
		contentRepository.On("GetAll", ctx, "").Return([]*model.Content{}, errors.New("error")).Once()

		result, err := contentService.GetAll(ctx, "")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("GetByID | Valid", func(t *testing.T) {
		contentRepository.On("GetByID", ctx, "1").Return(&model.Content{}, nil).Once()

		result, err := contentService.GetByID(ctx, "1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetByID | Invalid", func(t *testing.T) {
		contentRepository.On("GetByID", ctx, "0").Return(&model.Content{}, errors.New("whoops")).Once()

		result, err := contentService.GetByID(ctx, "0")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetByCategoryID(t *testing.T) {
	t.Run("GetByCategoryID | Valid", func(t *testing.T) {
		contentRepository.On("GetByCategoryID", ctx, "1").Return([]*model.Content{}, nil).Once()

		result, err := contentService.GetByCategoryID(ctx, "1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetByCategoryID | Invalid", func(t *testing.T) {
		contentRepository.On("GetByCategoryID", ctx, "0").Return([]*model.Content{}, errors.New("whoops")).Once()

		result, err := contentService.GetByCategoryID(ctx, "0")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetByUser(t *testing.T) {
	t.Run("GetByUser | Valid", func(t *testing.T) {
		contentRepository.On("GetByUser", ctx, model.User{}).Return([]*model.Content{}, nil).Once()

		result, err := contentService.GetByUser(ctx, model.User{})

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetByUser | Invalid", func(t *testing.T) {
		contentRepository.On("GetByUser", ctx, model.User{}).Return([]*model.Content{}, errors.New("whoops")).Once()

		result, err := contentService.GetByUser(ctx, model.User{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Create | Valid", func(t *testing.T) {
		contentRepository.On("Create", ctx, model.NewContent{}, model.User{}).Return(&model.Content{}, nil).Once()

		result, err := contentService.Create(ctx, model.NewContent{}, model.User{})

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Create | Invalid", func(t *testing.T) {
		contentRepository.On("Create", ctx, model.NewContent{}, model.User{}).Return(&model.Content{}, errors.New("whoops")).Once()

		result, err := contentService.Create(ctx, model.NewContent{}, model.User{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		contentRepository.On("Update", ctx, model.EditContent{}, model.User{}).Return(&model.Content{}, nil).Once()

		result, err := contentService.Update(ctx, model.EditContent{}, model.User{})

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Update | Invalid", func(t *testing.T) {
		contentRepository.On("Update", ctx, model.EditContent{}, model.User{}).Return(&model.Content{}, errors.New("whoops")).Once()

		result, err := contentService.Update(ctx, model.EditContent{}, model.User{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		contentRepository.On("Delete", ctx, model.DeleteContent{}, model.User{}).Return(true, nil).Once()

		res, err := contentService.Delete(ctx, model.DeleteContent{}, model.User{})

		assert.True(t, res)
		assert.Nil(t, err)
	})

	t.Run("Delete | Invalid", func(t *testing.T) {
		contentRepository.On("Delete", ctx, model.DeleteContent{}, model.User{}).Return(false, errors.New("whoops")).Once()

		res, err := contentService.Delete(ctx, model.DeleteContent{}, model.User{})

		assert.False(t, res)
		assert.NotNil(t, err)
	})
}
