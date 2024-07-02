package user_domain

import (
	entities "fdms/domain/entities/users"

	"github.com/jackc/pgx/v5/pgxpool"
)


type UserRepository interface {
	GetUser(id int64) (*entities.User, error)
	GetAll() ([]entities.User, error)
	Create(user *entities.User) (error)
	Update(user *entities.User) (error)
	Delete(id int64) (error)
	MapToDto(user *entities.User) (entities.UserDto)
	MapFromDto(user *entities.UserDto) (entities.User)
}

type UserImpl struct {
	db *pgxpool.Pool
}

func NewUserService(db *pgxpool.Pool) UserRepository {
	return &UserImpl{
		db : db,
	}
}