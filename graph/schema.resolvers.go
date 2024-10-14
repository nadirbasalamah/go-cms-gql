package graph

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver {
	// r.userService = services.InitUserService()
	// r.categoryService = services.InitCategoryService()
	// r.contentService = services.InitContentService()

	return &mutationResolver{r}
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
