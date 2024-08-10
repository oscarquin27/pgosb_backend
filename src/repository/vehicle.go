package repository

import (
	"context"
	"fdms/src/models"
	"fdms/src/services"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type VehicleRepository struct {
	db *pgxpool.Pool
}

func NewVehicleService(db *pgxpool.Pool) services.VehicleService {
	return &VehicleRepository{
		db: db,
	}
}

func (u *VehicleRepository) Get(id int64) (*models.Vehicle, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, `SELECT 
	id, 
	make, 
	model, 
	year, 
	drive, 
	cylinders, 
	engine_displacement, 
	fuel_type, 
	transmission, 
	vehicle_size_class, 
	base_model
	FROM vehicles.vehicle
 	where id = $1;`, id)

	if err != nil {
		return nil, err
	}

	vehicle, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Vehicle])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorVehicleNotFound
		}

		return nil, err
	}

	return &vehicle, nil
}

func (u *VehicleRepository) GetVehicleModels(make string) ([]models.Vehicle, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, `SELECT 
	id, 
	make, 
	model, 
	year, 
	drive, 
	cylinders, 
	engine_displacement, 
	fuel_type, 
	transmission, 
	vehicle_size_class, 
	base_model
	FROM vehicles.vehicle
 	where lower(make) = lower($1)`, make)

	if err != nil {
		return nil, err
	}

	vehicle, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Vehicle])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorVehicleNotFound
		}

		return nil, err
	}

	return vehicle, nil
}

func (u *VehicleRepository) GetVehicleTypes() ([]string, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, `SELECT 
	distinct make as make from vehicles.vehicle;`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tp string
	var tps []string

	for rows.Next() {
		err := rows.Scan(&tp)

		if err != nil {
			return nil, err
		}

		tps = append(tps, tp)
	}

	return tps, nil
}

func (u *VehicleRepository) GetAll() ([]models.Vehicle, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, `SELECT 
	id, 
	make, 
	model, 
	year, 
	drive, 
	cylinders, 
	engine_displacement, 
	fuel_type, 
	transmission, 
	vehicle_size_class, 
	base_model
	FROM vehicles.vehicle`)

	if err != nil {
		return nil, err
	}

	vehicle, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Vehicle])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorVehicleNotFound
		}

		return nil, err
	}

	return vehicle, nil
}

func (u *VehicleRepository) Create(vehicle *models.Vehicle) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, `INSERT INTO vehicles.vehicle
(make, model, year, drive, cylinders, engine_displacement, fuel_type, transmission, vehicle_size_class, base_model)
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);
`,
		vehicle.Make,
		vehicle.Model,
		vehicle.Year,
		vehicle.Drive,
		vehicle.Cylinders,
		vehicle.Engine_displacement,
		vehicle.Fuel_type,
		vehicle.Transmission,
		vehicle.Vehicle_size_class,
		vehicle.Base_model)

	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return models.ErrorVehicleNotCreated
}

func (u *VehicleRepository) Update(vehicle *models.Vehicle) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, `
		UPDATE vehicles.vehicle
		SET make=$1, 
		model=$2, 
		year=$3, 
		drive=$4, 
		cylinders=$5, 
		engine_displacement=$6, 
		fuel_type=$7, 
		transmission=$8, 
		vehicle_size_class=$9, 
		base_model=$10
		where id=$11;
		`,
		vehicle.Make,
		vehicle.Model,
		vehicle.Year,
		vehicle.Drive,
		vehicle.Cylinders,
		vehicle.Engine_displacement,
		vehicle.Fuel_type,
		vehicle.Transmission,
		vehicle.Vehicle_size_class,
		vehicle.Base_model,
		vehicle.Id)

	if err != nil {
		return err
	}

	if rows.RowsAffected() > 1 {
		return fmt.Errorf("se ha actualizado mas de un vehiculo")
	}

	if rows.RowsAffected() == 0 {
		return models.ErrorVehicleNotUpdated
	}

	return nil
}

func (u *VehicleRepository) Delete(id int64) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, "delete from vehicles.vehicle where id = $1", id)

	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return models.ErrorVehicleNotDeleted
}
