package unity_domain

import (
	entities "fdms/domain/entities/unities"

	"github.com/jackc/pgx/v5/pgxpool"
)


type UnityRepository interface {
	GetUnity(id int64) (*entities.Unity, error)
	GetAll() ([]entities.Unity, error)
	Create(user *entities.Unity) (error)
	Update(user *entities.Unity) (error)
	Delete(id int64) (error)
}

type UnityImpl struct {
	db *pgxpool.Pool
}

func NewUnityService(db *pgxpool.Pool) UnityRepository {
	return &UnityImpl{
		db : db,
	}
}