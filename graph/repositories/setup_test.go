package repositories_test

import (
	"context"
	"go-cms-gql/graph/repositories/mocks"
	"go-cms-gql/graph/services"
	"os"
	"testing"
)

var (
	userRepository mocks.UserRepository
	userService    services.UserService

	categoryRepository mocks.CategoryRepository
	categoryService    services.CategoryService

	contentRepository mocks.ContentRepository
	contentService    services.ContentService
	ctx               context.Context
)

func TestMain(m *testing.M) {
	userService = services.InitUserService(&userRepository)
	categoryService = services.InitCategoryService(&categoryRepository)
	contentService = services.InitContentService(&contentRepository)

	ctx = context.TODO()

	os.Exit(m.Run())
}
