package user

import "ecommerce/user"

type UserHandler struct {
	repo user.UserServe
}

func NewUserHandler(repo user.UserServe) *UserHandler {
	return &UserHandler{repo: repo}
}
