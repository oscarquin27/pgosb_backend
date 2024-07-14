package repository

import (
	"context"
	"fdms/src/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MissionVehicleRepository struct {
	db *pgxpool.Pool
}

func NewMissionVehicleService(db *pgxpool.Pool) *MissionVehicleRepository {
	return &MissionVehicleRepository{
		db: db,
	}
}

func (u *MissionVehicleRepository) Get(id int) ([]models.MissionVehicle, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, `SELECT 
	id, 
	service_id, 
	vehicle_condition, 
	make, 
	model, 
	year, 
	plate, 
	color, 
	vehicle_type, 
	motor_serial, 
	vehicle_verified
	FROM missions.vehicles
 	where service_id = $1;`, id)

	if err != nil {
		return nil, err
	}

	vehicle, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.MissionVehicle])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorMissionVehicleNotFound
		}

		return nil, err
	}

	return vehicle, nil
}

func (u *MissionVehicleRepository) Create(vehicle *models.MissionVehicle) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, `insert into missions.vehicles (service_id, 
	vehicle_condition, 
	make, 
	model, 
	year, 
	plate, 
	color, 
	vehicle_type, 
	motor_serial, 
	vehicle_verified)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);
`, vehicle.ServiceId, 
   vehicle.Make, 
   vehicle.Model, 
   vehicle.Year, 
   vehicle.Plate, 
   vehicle.Color, 
   vehicle.VehicleType, 
   vehicle.MotorSerial, 
   vehicle.VehicleVerified)

	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return models.ErrorMissionVehicleNotCreated
}

func (u *MissionVehicleRepository) Update(vehicle *models.MissionVehicle) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, `
		UPDATE missions.vehicles
		SET service_id = $1, 
	    vehicle_condition = $2, 
	    make = $3, 
	    model = $4, 
	    year = $5, 
	    plate = $6, 
	    color = $7, 
	    vehicle_type = $8, 
	    motor_serial = $9, 
	    vehicle_verified = $10
		where id = $11;
		`,
	vehicle.ServiceId,
	vehicle.VehicleCondition,
    vehicle.Make, 
    vehicle.Model, 
    vehicle.Year, 
    vehicle.Plate, 
    vehicle.Color, 
    vehicle.VehicleType, 
    vehicle.MotorSerial, 
    vehicle.VehicleVerified,
	vehicle.Id)

	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return models.ErrorVehicleNotUpdated
}

func (u *MissionVehicleRepository) Delete(id int64) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, "delete from missions.vehicles where id = $1", id)

	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return models.ErrorVehicleNotDeleted
}
