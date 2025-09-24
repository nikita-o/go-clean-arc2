package use_case

import (
	"clean-arc-2/internal/domains/user"
)

type UserUseCase struct {
	userRepository user.Repository
}

func NewUserUseCase(ur user.Repository) user.UseCase {
	return &UserUseCase{userRepository: ur}
}

func (u UserUseCase) CreateUser(user user.User) (user.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserUseCase) GetUserById(id int64) (user.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserUseCase) UpdateUser() (user.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserUseCase) DeleteUser(id int64) error {
	//TODO implement me
	panic("implement me")
}
