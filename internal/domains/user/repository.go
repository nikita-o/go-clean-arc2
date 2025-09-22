package user

type Repository interface {
	CreateUser(user User) (User, error)
	GetUserById(id int64) (User, error)
	UpdateUser() (User, error)
	DeleteUser(id int64) error
}
