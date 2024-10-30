package graph

import (
	"go-cms-gql/graph/repositories"
	"go-cms-gql/graph/services"
	"go-cms-gql/utils"
)

type Resolver struct {
	userService           services.UserService
	categoryService       services.CategoryService
	contentService        services.ContentService
	recommendationService services.RecommendationService
}

func InitResolver() *Resolver {
	tokenGenFunc := utils.GenerateNewAccessToken

	return &Resolver{
		userService: services.InitUserService(
			repositories.InitUserRepository(),
			tokenGenFunc,
		),
		categoryService:       services.InitCategoryService(repositories.InitCategoryRepository()),
		contentService:        services.InitContentService(repositories.InitContentRepository()),
		recommendationService: services.InitRecommendationService(),
	}
}
