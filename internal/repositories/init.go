package repositories

import (
	"clean-arc-2/internal/domains/user"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repositories struct {
	UserRepository user.Repository
}

type DataInitRepositories struct {
	Database *pgxpool.Pool
}

func NewRepositories(data DataInitRepositories) *Repositories {
	return &Repositories{
		UserRepository: NewUserRepository(data.Database),
	}
}
