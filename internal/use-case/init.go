package use_case

import (
	"clean-arc-2/internal/domains/user"
	"clean-arc-2/internal/repositories"
)

type UseCases struct {
	UserUseCase user.UseCase
}

type DataInitUseCases struct {
	Repositories *repositories.Repositories
}

func NewUseCases(data DataInitUseCases) *UseCases {
	return &UseCases{
		UserUseCase: NewUserUseCase(data.Repositories.UserRepository),
	}
}
