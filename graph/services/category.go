package services

import (
	"context"
	"go-cms-gql/graph/model"
	"go-cms-gql/graph/repositories"
)

type CategoryServiceImpl struct {
	repository repositories.CategoryRepository
}

func InitCategoryService(categoryRepository repositories.CategoryRepository) CategoryService {
	return &CategoryServiceImpl{
		repository: categoryRepository,
	}
}

func (cs *CategoryServiceImpl) GetAll(ctx context.Context) ([]*model.Category, error) {
	return cs.repository.GetAll(ctx)
}

func (cs *CategoryServiceImpl) GetByID(ctx context.Context, categoryID string) (*model.Category, error) {
	return cs.repository.GetByID(ctx, categoryID)
}

func (cs *CategoryServiceImpl) Create(ctx context.Context, input model.NewCategory) (*model.Category, error) {
	return cs.repository.Create(ctx, input)
}

func (cs *CategoryServiceImpl) Update(ctx context.Context, input model.EditCategory) (*model.Category, error) {
	return cs.repository.Update(ctx, input)
}

func (cs *CategoryServiceImpl) Delete(ctx context.Context, input model.DeleteCategory) (bool, error) {
	return cs.repository.Delete(ctx, input)
}
