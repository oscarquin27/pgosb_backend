package vehicle_domain

import (
	entities "fdms/domain/entities/vehicles"

	"github.com/jackc/pgx/v5/pgxpool"
)


type VehicleRepository interface {
	GetVehicle(id int64) (*entities.Vehicle, error)
	GetAll() ([]entities.Vehicle, error)
	Create(user *entities.Vehicle) (error)
	Update(user *entities.Vehicle) (error)
	Delete(id int64) (error)
}

type VehicleImpl struct {
	db *pgxpool.Pool
}

func NewVehicleService(db *pgxpool.Pool) VehicleRepository {
	return &VehicleImpl{
		db : db,
	}
}