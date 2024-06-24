package user_domain

import (
	entities "fdms/domain/entities/users"

	"github.com/jackc/pgx/v5/pgxpool"
)


type UserRepository interface {
	GetUser(id int64) (*entities.User, error)
	GetAll() ([]entities.User, error)
	Create(user *entities.UserCreateDto) (error)
	Update(user *entities.UserUpdateDto) (error)
	Delete(id int64) (error)
}

type UserImpl struct {
	db *pgxpool.Pool
}

func NewUserService(db *pgxpool.Pool) UserRepository {
	return &UserImpl{
		db : db,
	}
}