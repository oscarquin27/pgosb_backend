package vehicle_domain

import (
	"context"
	entities "fdms/domain/entities/vehicles"

	"github.com/jackc/pgx/v5"
)


func (u *VehicleImpl) GetVehicle(id int64) (*entities.Vehicle, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return nil,err
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

	vehicle, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[entities.Vehicle])
	
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, entities.ErrorVehicleNotFound
		}

		return nil, err
	}

	return &vehicle,nil
}

func (u *VehicleImpl) GetAll() ([]entities.Vehicle, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return nil,err
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

	vehicle, err := pgx.CollectRows(rows, pgx.RowToStructByName[entities.Vehicle])
	
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, entities.ErrorVehicleNotFound
		}

		return nil, err
	}

	return vehicle,nil
}

func (u *VehicleImpl) Create(vehicle *entities.Vehicle) (error) {
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

	return entities.ErrorVehicleNotCreated
}

func (u *VehicleImpl) Update(vehicle *entities.Vehicle) (error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx,`
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

	if rows.RowsAffected() > 0 {
		return nil
	}

	return entities.ErrorVehicleNotUpdated
}

func (u *VehicleImpl) Delete(id int64) (error) {
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

	return entities.ErrorVehicleNotDeleted
}
