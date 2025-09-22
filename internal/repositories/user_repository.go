package repositories

import (
	"clean-arc-2/internal/domains/user"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	database *pgxpool.Pool
}

func NewUserRepository(database *pgxpool.Pool) user.Repository {
	return &UserRepository{
		database: database,
	}
}

func (r UserRepository) CreateUser(user user.User) (user.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r UserRepository) GetUserById(id int64) (user.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r UserRepository) UpdateUser() (user.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r UserRepository) DeleteUser(id int64) error {
	//TODO implement me
	panic("implement me")
}
