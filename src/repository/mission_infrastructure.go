package repository

import (
	"context"
	"fdms/src/mikro"
	"fdms/src/models"
	"fdms/src/services"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MissionInfrastructureRepository struct {
	db *pgxpool.Pool
}

func NewMissionInfrastructureService(db *pgxpool.Pool) services.MissionInfrastructureService {
	return &MissionInfrastructureRepository{
		db: db,
	}
}

// GetAll implements services.MissionInfrastructureService.
func (u *MissionInfrastructureRepository) GetAll() ([]models.MissionInfrastructure, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return nil, err
	}

	defer conn.Release()

	rows, err := conn.Query(ctx, `
	SELECT *
	FROM missions.infrastructure`)

	if err != nil {
		return nil, err
	}

	infra, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.MissionInfrastructure])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorMissionInfrastructureNotFound
		}

		return nil, err
	}

	return infra, nil
}

// GetByServiceId implements services.MissionInfrastructureService.
func (u *MissionInfrastructureRepository) GetMissionId(id int) ([]models.MissionInfrastructure, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return nil, err
	}

	defer conn.Release()

	rows, err := conn.Query(ctx, `
	SELECT *
	FROM missions.infrastructure
	WHERE mission_id = $1`, id)

	if err != nil {
		return nil, err
	}

	infra, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.MissionInfrastructure])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorMissionInfrastructureNotFound
		}

		return nil, err
	}

	return infra, nil
}

func (u *MissionInfrastructureRepository) Get(id int) (*models.MissionInfrastructure, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return nil, err
	}

	defer conn.Release()

	rows, err := conn.Query(ctx, `
	SELECT *
	FROM missions.infrastructure
	WHERE id = $1`, id)

	if err != nil {
		return nil, err
	}

	infra, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.MissionInfrastructure])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorMissionInfrastructureNotFound
		}

		return nil, err
	}

	return &infra, nil
}

func (u *MissionInfrastructureRepository) Create(infra *models.MissionInfrastructure) error {
	m := mikro.NewMkModel(u.db)

	rows, err := m.Model(infra).Omit("id").Insert("missions.infrastructure")

	if err != nil {
		return err
	}

	if rows > 0 {
		return nil
	}

	return models.ErrorMissionInfrastructureNotCreated
}

func (u *MissionInfrastructureRepository) Update(infra *models.MissionInfrastructure) error {
	m := mikro.NewMkModel(u.db)

	rows, err := m.Model(infra).Omit("id").Omit("mission_id").Where("id", "=", infra.Id).Update("missions.infrastructure")

	if err != nil {
		return err
	}

	if rows > 0 {
		return nil
	}

	return models.ErrorMissionInfrastructureNotUpdated
}

func (u *MissionInfrastructureRepository) Delete(id int) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return err
	}

	defer conn.Release()

	rows, err := conn.Exec(ctx, "delete from missions.infrastructure where id = $1", id)

	if err != nil {
		return err
	}

	numberOfRows := rows.RowsAffected()

	if numberOfRows == 1 {
		return nil
	}

	return models.ErrorVehicleNotDeleted
}
