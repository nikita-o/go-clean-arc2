package controllers

import (
	"clean-arc-2/internal/domains/user"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUseCase user.UseCase
}

func NewUserController(uu user.UseCase) *UserController {
	return &UserController{
		userUseCase: uu,
	}
}

func (u UserController) CreateUser(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u UserController) GetUserById(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u UserController) UpdateUser(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u UserController) DeleteUser(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
