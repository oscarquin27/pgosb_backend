package unit_domain

import (
	entities "fdms/domain/entities/units"

	"github.com/jackc/pgx/v5/pgxpool"
)


type UnitRepository interface {
	GetUnit(id int64) (*entities.Unit, error)
	GetAll() ([]entities.Unit, error)
	Create(user *entities.Unit) (error)
	Update(user *entities.Unit) (error)
	Delete(id int64) (error)
}

type UnitImpl struct {
	db *pgxpool.Pool
}

func NewUnityService(db *pgxpool.Pool) UnitRepository {
	return &UnitImpl{
		db : db,
	}
}