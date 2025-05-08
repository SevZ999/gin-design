package controller

import "github.com/gin-gonic/gin"

type UserService interface {
}

type UserController struct {
	srv UserService
}

func NewUserController(srv UserService) *UserController {
	return &UserController{srv: srv}
}

func (ctrl *UserController) GetUser(c *gin.Context) {

}
