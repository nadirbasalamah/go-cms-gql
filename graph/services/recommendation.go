package services

import (
	"context"
	"go-cms-gql/graph/model"
	"go-cms-gql/utils"
)

type RecommendationServiceImpl struct {
}

func InitRecommendationService() RecommendationService {
	return &RecommendationServiceImpl{}
}

func (rs *RecommendationServiceImpl) GenerateContent(ctx context.Context, generateInput model.GenerateContent) (string, error) {
	return utils.GenerateContent(ctx, generateInput)
}

func (rs *RecommendationServiceImpl) GetTags(ctx context.Context, content string) ([]string, error) {
	return utils.GetTags(ctx, content)
}
