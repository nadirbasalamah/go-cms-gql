package utils

import "go-cms-gql/graph/model"

func ConvertToUserData(user *model.User) *model.UserData {
	return &model.UserData{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
