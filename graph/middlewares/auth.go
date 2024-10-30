package middlewares

import (
	"context"
	"net/http"

	"go-cms-gql/graph/model"
	"go-cms-gql/graph/repositories"
	"go-cms-gql/graph/services"
	"go-cms-gql/utils"
)

func NewMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var header string = r.Header.Get("Authorization")

			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			tokenData, err := utils.CheckToken(r)

			if err != nil {
				http.Error(w, "invalid token", http.StatusForbidden)
				return
			}

			var userService services.UserService = services.InitUserService(repositories.InitUserRepository(), utils.GenerateNewAccessToken)

			userData, err := userService.GetUserInfo(r.Context(), tokenData.UserId)

			if err != nil {
				http.Error(w, "user not found", http.StatusForbidden)
				return
			}

			var user model.User = *userData

			ctx := context.WithValue(r.Context(), utils.UserCtxKey, &user)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
