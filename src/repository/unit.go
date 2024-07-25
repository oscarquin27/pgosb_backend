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

func (u *UnitRepository) GetUnitTypes() ([]string, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, `SELECT 
	distinct unit_type as unit_type from locations.unit`)

	if err != nil {
		return nil, err
	}

	units, err := pgx.CollectRows(rows, pgx.RowToStructByName[string])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorUnitNotFound
		}

		return nil, err
	}

	return units, nil
}

func (u *UnitRepository) Get(id int64) (*models.Unit, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, ` SELECT 
	id, 
	plate, 
	station, 
	unit_type, 
	make, 
	unit_condition, 
	vehicle_serial, 
               motor_serial, 
			   capacity, 
			   details, 
			   fuel_type, 
			   water_capacity, 
			   observations,
               hurt_capacity, 
			   doors, 
			   performance, 
			   load_capacity,
			    model, 
				alias, 
				color,
               year, 
			   purpose, 
			   init_kilometer
        FROM vehicles.unit
        WHERE id = $1`, id)

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

	rows, err := conn.Query(ctx, ` SELECT id, plate, station, unit_type, make, unit_condition, vehicle_serial, 
               motor_serial, capacity, details, fuel_type, water_capacity, observations,
               hurt_capacity, doors, performance, load_capacity, model, alias, color,
               year, purpose, init_kilometer
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

func (u *UnitRepository) Create(unit *models.Unit) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, `
	INSERT INTO vehicles.unit (
		plate, station, unit_type, make, unit_condition, vehicle_serial, motor_serial, 
		capacity, details, fuel_type, water_capacity, observations, hurt_capacity, 
		doors, performance, load_capacity, model, alias, color, year, purpose, init_kilometer
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22)`,
		unit.Plate, unit.Station, unit.Unit_type, unit.Make, unit.Unit_condition, unit.Vehicle_serial,
		unit.Motor_serial, unit.Capacity, unit.Details, unit.Fuel_type, unit.Water_capacity,
		unit.Observations, unit.Hurt_capacity, unit.Doors, unit.Performance, unit.Load_capacity,
		unit.Model, unit.Alias, unit.Color, unit.Year, unit.Purpose, unit.Init_kilometer,
	)

	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return models.ErrorUnitNotCreated
}

func (u *UnitRepository) Update(unit *models.Unit) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, `
	UPDATE vehicles.unit
	SET plate = $1, station = $2, unit_type = $3, make = $4, unit_condition = $5,
		vehicle_serial = $6, motor_serial = $7, capacity = $8, details = $9, 
		fuel_type = $10, water_capacity = $11, observations = $12, hurt_capacity = $13,
		doors = $14, performance = $15, load_capacity = $16, model = $17, alias = $18,
		color = $19, year = $20, purpose = $21, init_kilometer = $22
	WHERE id = $23`,
		unit.Plate, unit.Station, unit.Unit_type, unit.Make, unit.Unit_condition, unit.Vehicle_serial,
		unit.Motor_serial, unit.Capacity, unit.Details, unit.Fuel_type, unit.Water_capacity,
		unit.Observations, unit.Hurt_capacity, unit.Doors, unit.Performance, unit.Load_capacity,
		unit.Model, unit.Alias, unit.Color, unit.Year, unit.Purpose, unit.Init_kilometer, unit.Id,
	)

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

	rows, err := conn.Exec(ctx, "delete from vehicles.unit where id = $1", id)

	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return models.ErrorUnitNotDeleted
}
