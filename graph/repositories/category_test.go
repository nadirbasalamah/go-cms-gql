package repositories_test

import (
	"errors"
	"go-cms-gql/graph/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllCategories(t *testing.T) {
	t.Run("GetAll | Valid", func(t *testing.T) {
		categoryRepository.On("GetAll", ctx).Return([]*model.Category{}, nil).Once()

		result, err := categoryService.GetAll(ctx)

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetAll | Invalid", func(t *testing.T) {
		categoryRepository.On("GetAll", ctx).Return([]*model.Category{}, errors.New("error")).Once()

		result, err := categoryService.GetAll(ctx)

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetCategoryByID(t *testing.T) {
	t.Run("GetByID | Valid", func(t *testing.T) {
		categoryRepository.On("GetByID", ctx, "1").Return(&model.Category{}, nil).Once()

		result, err := categoryService.GetByID(ctx, "1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetByID | Invalid", func(t *testing.T) {
		categoryRepository.On("GetByID", ctx, "0").Return(&model.Category{}, errors.New("whoops")).Once()

		result, err := categoryService.GetByID(ctx, "0")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestCreateCategory(t *testing.T) {
	t.Run("Create | Valid", func(t *testing.T) {
		categoryRepository.On("Create", ctx, model.NewCategory{}).Return(&model.Category{}, nil).Once()

		result, err := categoryService.Create(ctx, model.NewCategory{})

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Create | Invalid", func(t *testing.T) {
		categoryRepository.On("Create", ctx, model.NewCategory{}).Return(&model.Category{}, errors.New("whoops")).Once()

		result, err := categoryService.Create(ctx, model.NewCategory{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestUpdateCategory(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		categoryRepository.On("Update", ctx, model.EditCategory{}).Return(&model.Category{}, nil).Once()

		result, err := categoryService.Update(ctx, model.EditCategory{})

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Update | Invalid", func(t *testing.T) {
		categoryRepository.On("Update", ctx, model.EditCategory{}).Return(&model.Category{}, errors.New("whoops")).Once()

		result, err := categoryService.Update(ctx, model.EditCategory{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestDeleteCategory(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		categoryRepository.On("Delete", ctx, model.DeleteCategory{}).Return(true, nil).Once()

		res, err := categoryService.Delete(ctx, model.DeleteCategory{})

		assert.True(t, res)
		assert.Nil(t, err)
	})

	t.Run("Delete | Invalid", func(t *testing.T) {
		categoryRepository.On("Delete", ctx, model.DeleteCategory{}).Return(false, errors.New("whoops")).Once()

		res, err := categoryService.Delete(ctx, model.DeleteCategory{})

		assert.False(t, res)
		assert.NotNil(t, err)
	})
}
