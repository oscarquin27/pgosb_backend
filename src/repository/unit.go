package repository

import (
	"context"
	"fdms/src/models"
	"fdms/src/services"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UnitRepository struct {
	db *pgxpool.Pool
}

func NewUnityService(db *pgxpool.Pool) services.UnitService {
	return &UnitRepository{
		db: db,
	}
}

func (u *UnitRepository) Get(id int64) (*models.Unit, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, `SELECT 
	id, 
	plate, 
	zone, 
	station, 
	unit_type, 
	make, 
	drivers, 
	unit_condition, 
	vehicle_serial, 
	motor_serial, 
	capacity, 
	details, 
	fuel_type, 
	water_capacity, 
	observations
	FROM vehicles.unit
 	where id = $1;`, id)

	if err != nil {
		return nil, err
	}

	unity, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Unit])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorUnitNotFound
		}

		return nil, err
	}

	return &unity, nil
}

func (u *UnitRepository) GetAll() ([]models.Unit, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, `SELECT 
	id, 
	plate, 
	zone, 
	station, 
	unit_type, 
	make, 
	drivers, 
	unit_condition, 
	vehicle_serial, 
	motor_serial, 
	capacity, 
	details, 
	fuel_type, 
	water_capacity, 
	observations
	FROM vehicles.unit`)

	if err != nil {
		return nil, err
	}

	unity, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Unit])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorUnitNotFound
		}

		return nil, err
	}

	return unity, nil
}

func (u *UnitRepository) Create(unity *models.Unit) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, `INSERT INTO vehicles.unit
(plate, zone, station, unit_type, make, drivers, unit_condition, vehicle_serial, motor_serial, capacity, details, fuel_type, water_capacity, observations)
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14);
`,
		unity.Plate,
		unity.Zone,
		unity.Station,
		unity.Unit_type,
		unity.Make,
		unity.Drivers,
		unity.Unit_condition,
		unity.Vehicle_serial,
		unity.Motor_serial,
		unity.Capacity,
		unity.Details,
		unity.Fuel_type,
		unity.Water_capacity,
		unity.Observations)

	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return models.ErrorUnitNotCreated
}

func (u *UnitRepository) Update(unity *models.Unit) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, `
		UPDATE vehicles.unit
		SET plate=$1, 
			zone=$2, 
			station=$3, 
			unit_type=$4, 
			make=$5, 
			drivers=$6, 
			unit_condition=$7, 
			vehicle_serial=$8, 
			motor_serial=$9, 
			capacity=$10, 
			details=$11, 
			fuel_type=$12, 
			water_capacity=$13, 
			observations=$14
		WHERE id=$15;
		`,
		unity.Plate,
		unity.Zone,
		unity.Station,
		unity.Unit_type,
		unity.Make,
		unity.Drivers,
		unity.Unit_condition,
		unity.Vehicle_serial,
		unity.Motor_serial,
		unity.Capacity,
		unity.Details,
		unity.Fuel_type,
		unity.Water_capacity,
		unity.Observations,
		unity.Id)

	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return models.ErrorUnitNotUpdated
}

func (u *UnitRepository) Delete(id int64) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, "delete from vehicles.unity where id = $1", id)

	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return models.ErrorUnitNotDeleted
}
