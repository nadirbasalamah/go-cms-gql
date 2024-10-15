// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type Category struct {
	ID        string     `json:"id" bson:"_id,omitempty"`
	Title     string     `json:"title" bson:"title"`
	CreatedAt time.Time  `json:"createdAt" bson:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" bson:"updatedAt"`
}

type Content struct {
	ID        string     `json:"id" bson:"_id,omitempty"`
	Title     string     `json:"title" bson:"title"`
	Content   string     `json:"content" bson:"content"`
	Author    *UserData  `json:"author,omitempty" bson:"author"`
	Category  *Category  `json:"category,omitempty" bson:"category"`
	CreatedAt time.Time  `json:"createdAt" bson:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" bson:"updatedAt"`
}

type DeleteCategory struct {
	CategoryID string `json:"categoryId" bson:"categoryId"`
}

type DeleteContent struct {
	ContentID string `json:"contentId" bson:"contentId"`
}

type EditCategory struct {
	CategoryID string `json:"categoryId" bson:"categoryId"`
	Title      string `json:"title" bson:"title"`
}

type EditContent struct {
	ContentID  string `json:"contentId" bson:"contentId"`
	Title      string `json:"title" bson:"title"`
	Content    string `json:"content" bson:"content"`
	CategoryID string `json:"categoryId" bson:"categoryId"`
}

type LoginInput struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type Mutation struct {
}

type NewCategory struct {
	Title string `json:"title" bson:"title"`
}

type NewContent struct {
	Title      string `json:"title" bson:"title"`
	Content    string `json:"content" bson:"content"`
	CategoryID string `json:"categoryId" bson:"categoryId"`
}

type NewUser struct {
	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type Query struct {
}

type User struct {
	ID        string     `json:"id" bson:"_id,omitempty"`
	Username  string     `json:"username" bson:"username"`
	Email     string     `json:"email" bson:"email"`
	Password  string     `json:"password" bson:"password"`
	CreatedAt time.Time  `json:"createdAt" bson:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" bson:"updatedAt"`
}

type UserData struct {
	ID        string     `json:"id" bson:"_id,omitempty"`
	Username  string     `json:"username" bson:"username"`
	Email     string     `json:"email" bson:"email"`
	CreatedAt time.Time  `json:"createdAt" bson:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" bson:"updatedAt"`
}
