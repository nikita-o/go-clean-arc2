package repositories

import (
	"clean-arc-2/internal/domains/user"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repositories struct {
	UserRepository user.Repository
}

func NewRepositories(database *pgxpool.Pool) *Repositories {
	return &Repositories{
		UserRepository: NewUserRepository(database),
	}
}
