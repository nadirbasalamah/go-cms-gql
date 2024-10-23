package services

import (
	"context"
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

func (cs *ContentServiceImpl) GetAll(ctx context.Context) ([]*model.Content, error) {
	return cs.repository.GetAll(ctx)
}

func (cs *ContentServiceImpl) GetByID(ctx context.Context, contentID string) (*model.Content, error) {
	return cs.repository.GetByID(ctx, contentID)
}

func (cs *ContentServiceImpl) Create(ctx context.Context, input model.NewContent, user model.User) (*model.Content, error) {
	return cs.repository.Create(ctx, input, user)
}

func (cs *ContentServiceImpl) Update(ctx context.Context, input model.EditContent, user model.User) (*model.Content, error) {
	return cs.repository.Update(ctx, input, user)
}

func (cs *ContentServiceImpl) Delete(ctx context.Context, input model.DeleteContent, user model.User) (bool, error) {
	return cs.repository.Delete(ctx, input, user)
}
