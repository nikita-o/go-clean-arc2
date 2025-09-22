package use_case

import (
	"clean-arc-2/internal/domains/user"
	"clean-arc-2/internal/repositories"
)

type UseCases struct {
	UserUseCase user.UseCase
}

func NewUseCases(repositories repositories.Repositories) *UseCases {
	return &UseCases{
		UserUseCase: NewUserUseCase(repositories.UserRepository),
	}
}
