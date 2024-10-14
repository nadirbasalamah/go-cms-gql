package graph

import "go-cms-gql/graph/services"

type Resolver struct {
	userService     services.UserService
	categoryService services.CategoryService
	contentService  services.ContentService
}

func InitResolver() *Resolver {
	return &Resolver{
		userService:     services.InitUserService(),
		categoryService: services.InitCategoryService(),
		contentService:  services.InitContentService(),
	}
}
