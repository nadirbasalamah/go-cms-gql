package graph

import (
	"go-cms-gql/graph/repositories"
	"go-cms-gql/graph/services"
)

type Resolver struct {
	userService           services.UserService
	categoryService       services.CategoryService
	contentService        services.ContentService
	recommendationService services.RecommendationService
}

func InitResolver() *Resolver {
	return &Resolver{
		userService:           services.InitUserService(repositories.InitUserRepository()),
		categoryService:       services.InitCategoryService(repositories.InitCategoryRepository()),
		contentService:        services.InitContentService(repositories.InitContentRepository()),
		recommendationService: services.InitRecommendationService(),
	}
}
