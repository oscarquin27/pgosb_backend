package repository

import (
	"context"
	"fdms/src/mikro"
	"fdms/src/models"
	"fdms/src/services"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MissionVehicleRepository struct {
	db *pgxpool.Pool
}

func NewMissionVehicleService(db *pgxpool.Pool) services.MissionVehicleService {
	return &MissionVehicleRepository{
		db: db,
	}
}

// GetAll implements services.MissionVehicleService.
func (u *MissionVehicleRepository) GetAll() ([]models.MissionVehicle, error) {

	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, `
	SELECT *
	FROM missions.vehicles`)

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

// GetByServiceId implements services.MissionVehicleService.
func (u *MissionVehicleRepository) GetMissionId(id int) ([]models.MissionVehicle, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, `
	SELECT *
	FROM missions.vehicles
 	WHERE mission_id = $1`, id)

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

func (u *MissionVehicleRepository) Get(id int) (*models.MissionVehicle, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, `
	SELECT *
	FROM missions.vehicles
 	where id = $1`, id)

	if err != nil {
		return nil, err
	}

	vehicle, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.MissionVehicle])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorMissionVehicleNotFound
		}

		return nil, err
	}

	return &vehicle, nil
}

func (u *MissionVehicleRepository) Create(vehicle *models.MissionVehicle) error {
	m := mikro.NewMkModel(u.db)

	rows, err := m.Model(vehicle).Omit("id").Insert("missions.vehicles")

	if err != nil {
		return err
	}

	if rows > 0 {
		return nil
	}

	return models.ErrorMissionVehicleNotCreated
}

func (u *MissionVehicleRepository) Update(vehicle *models.MissionVehicle) error {
	m := mikro.NewMkModel(u.db)

	rows, err := m.Model(vehicle).Omit("id").
		Omit("mission_id").Where("id", "=", vehicle.Id).Update("missions.vehicles")

	if err != nil {
		return err
	}

	if rows > 0 {
		return nil
	}

	return models.ErrorMissionVehicleNotUpdated
}

func (u *MissionVehicleRepository) Delete(id int) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return err
	}

	defer conn.Release()

	rows, err := conn.Exec(ctx, "delete from missions.vehicles where id = $1", id)

	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return models.ErrorVehicleNotDeleted
}
