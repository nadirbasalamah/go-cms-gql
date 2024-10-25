package directives

import (
	"context"
	"errors"
	"go-cms-gql/utils"

	"github.com/99designs/gqlgen/graphql"
)

func CheckAdmin(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	if err := utils.CheckAdminRole(ctx); err != nil {
		return nil, errors.New("access denied")
	}

	return next(ctx)
}

func GetAuthenticatedUser(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	_, err := utils.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, errors.New("access denied")
	}

	return next(ctx)
}
