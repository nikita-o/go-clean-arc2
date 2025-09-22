package controllers

import use_case "clean-arc-2/internal/use-case"

type Controllers struct {
	UserController *UserController
}

func NewControllers(useCases use_case.UseCases) *Controllers {
	return &Controllers{
		UserController: NewUserController(useCases.UserUseCase),
	}
}
