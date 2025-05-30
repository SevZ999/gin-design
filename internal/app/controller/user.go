package controller

import "loan-admin/internal/app/service"

type UserController struct {
	srv *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		srv: userService,
	}
}
