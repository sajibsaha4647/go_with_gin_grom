package user

import (
	"ecommerce/config"
	"ecommerce/user"
)

type UserHandler struct {
	repo user.UserServe
	cfg *config.Config
}

func NewUserHandler(repo user.UserServe,cfg *config.Config) *UserHandler {
	return &UserHandler{repo: repo,cfg:cfg}
}
