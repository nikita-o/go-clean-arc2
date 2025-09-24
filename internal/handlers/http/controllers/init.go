package controllers

import (
	use_case "clean-arc-2/internal/use-case"
)

type Controllers struct {
	UserController *UserController
}

type DataInitControllers struct {
	UseCases *use_case.UseCases
}

func NewControllers(data DataInitControllers) *Controllers {
	return &Controllers{
		UserController: NewUserController(data.UseCases.UserUseCase),
	}
}
