package unity_domain

import (
	"context"
	entities "fdms/domain/entities/unities"

	"github.com/jackc/pgx/v5"
)


func (u *UnityImpl) GetUnity(id int64) (*entities.Unity, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return nil,err
	}

	rows, err := conn.Query(ctx, `SELECT 
	id, 
	plate, 
	zone, 
	station, 
	unity_type, 
	make, 
	drivers, 
	unity_condition, 
	vehicle_serial, 
	motor_serial, 
	capacity, 
	details, 
	fuel_type, 
	water_capacity, 
	observations
	FROM vehicles.unity
 	where id = $1;`, id)

	unity, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[entities.Unity])
	
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, entities.ErrorUnityNotFound
		}

		return nil, err
	}

	return &unity,nil
}

func (u *UnityImpl) GetAll() ([]entities.Unity, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return nil,err
	}

	rows, err := conn.Query(ctx, `SELECT 
	id, 
	plate, 
	zone, 
	station, 
	unity_type, 
	make, 
	drivers, 
	unity_condition, 
	vehicle_serial, 
	motor_serial, 
	capacity, 
	details, 
	fuel_type, 
	water_capacity, 
	observations
	FROM vehicles.unity`)

	unity, err := pgx.CollectRows(rows, pgx.RowToStructByName[entities.Unity])
	
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, entities.ErrorUnityNotFound
		}

		return nil, err
	}

	return unity,nil
}

func (u *UnityImpl) Create(unity *entities.Unity) (error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, `INSERT INTO vehicles.unity
(plate, zone, station, unity_type, make, drivers, unity_condition, vehicle_serial, motor_serial, capacity, details, fuel_type, water_capacity, observations)
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14);
`, 
	unity.Plate,
	unity.Zone,
	unity.Station,
	unity.Unity_type,
	unity.Make,
	unity.Drivers,
	unity.Unity_condition,
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

	return entities.ErrorUnityNotCreated
}

func (u *UnityImpl) Update(unity *entities.Unity) (error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx,`
		UPDATE vehicles.unity
		SET plate=$1, 
			zone=$2, 
			station=$3, 
			unity_type=$4, 
			make=$5, 
			drivers=$6, 
			unity_condition=$7, 
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
		unity.Unity_type,
		unity.Make,
		unity.Drivers,
		unity.Unity_condition,
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

	return entities.ErrorUnityNotUpdated
}

func (u *UnityImpl) Delete(id int64) (error) {
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

	return entities.ErrorUnityNotDeleted
}
