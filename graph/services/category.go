package services

import (
	"go-cms-gql/graph/model"
	"go-cms-gql/graph/repositories"
)

type CategoryServiceImpl struct {
	repository repositories.CategoryRepository
}

func InitCategoryService() CategoryService {
	return &CategoryServiceImpl{
		repository: repositories.InitCategoryRepository(),
	}
}

func (cs *CategoryServiceImpl) GetAll() ([]*model.Category, error) {
	return cs.repository.GetAll()
}

func (cs *CategoryServiceImpl) GetByID(categoryID string) (*model.Category, error) {
	return cs.repository.GetByID(categoryID)
}

func (cs *CategoryServiceImpl) Create(input model.NewCategory) (*model.Category, error) {
	return cs.repository.Create(input)
}

func (cs *CategoryServiceImpl) Update(input model.EditCategory) (*model.Category, error) {
	return cs.repository.Update(input)
}

func (cs *CategoryServiceImpl) Delete(input model.DeleteCategory) (bool, error) {
	return cs.repository.Delete(input)
}
