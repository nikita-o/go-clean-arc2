package user

type UseCase interface {
	CreateUser(user User) (User, error)
	GetUserById(id int64) (User, error)
	UpdateUser() (User, error)
	DeleteUser(id int64) error
}
