package services

import (
	"context"
	"go-cms-gql/graph/model"
	"go-cms-gql/graph/repositories"
)

type ContentServiceImpl struct {
	repository repositories.ContentRepository
}

func InitContentService(contentRepository repositories.ContentRepository) ContentService {
	return &ContentServiceImpl{
		repository: contentRepository,
	}
}

func (cs *ContentServiceImpl) GetAll(ctx context.Context, keyword string) ([]*model.Content, error) {
	return cs.repository.GetAll(ctx, keyword)
}

func (cs *ContentServiceImpl) GetByID(ctx context.Context, contentID string) (*model.Content, error) {
	return cs.repository.GetByID(ctx, contentID)
}

func (cs *ContentServiceImpl) GetByCategoryID(ctx context.Context, categoryID string) ([]*model.Content, error) {
	return cs.repository.GetByCategoryID(ctx, categoryID)
}

func (cs *ContentServiceImpl) GetByUser(ctx context.Context, user model.User) ([]*model.Content, error) {
	return cs.repository.GetByUser(ctx, user)
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
