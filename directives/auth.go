package directives

import (
	"context"
	"errors"
	"go-cms-gql/graph/middlewares"

	"github.com/99designs/gqlgen/graphql"
)

func CheckAdmin(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	if err := middlewares.CheckAdminRole(ctx); err != nil {
		return nil, errors.New("access denied")
	}

	return next(ctx)
}
