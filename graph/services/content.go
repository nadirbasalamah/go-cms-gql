package services

import (
	"go-cms-gql/graph/model"
	"go-cms-gql/graph/repositories"
)

type ContentServiceImpl struct {
	repository repositories.ContentRepository
}

func InitContentService() ContentService {
	return &ContentServiceImpl{
		repository: repositories.InitContentRepository(),
	}
}

func (cs *ContentServiceImpl) GetAll() ([]*model.Content, error) {
	return cs.repository.GetAll()
}

func (cs *ContentServiceImpl) GetByID(contentID string) (*model.Content, error) {
	return cs.repository.GetByID(contentID)
}

func (cs *ContentServiceImpl) Create(input model.NewContent, user model.User) (*model.Content, error) {
	return cs.repository.Create(input, user)
}

func (cs *ContentServiceImpl) Update(input model.EditContent, user model.User) (*model.Content, error) {
	return cs.repository.Update(input, user)
}

func (cs *ContentServiceImpl) Delete(input model.DeleteContent, user model.User) (bool, error) {
	return cs.repository.Delete(input, user)
}
