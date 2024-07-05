package user_application

import (
	user_domain "fdms/domain/entities/users"
	"fdms/repository/user_repository/postgres_respository"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserService interface {
	user_domain.UserRepository
}

type UserServiceImpl struct {
	*postgres_respository.UserImpl
}

func NewService(db *pgxpool.Pool) UserService {
	return &UserServiceImpl{
		postgres_respository.NewUserImplDb(db),
	}
}
