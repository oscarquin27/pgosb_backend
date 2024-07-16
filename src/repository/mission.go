package repository

import (
	"context"
	"fdms/src/models"
	"fdms/src/services"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MissionRepository struct {
	db *pgxpool.Pool
}

// GetAll implements services.MissionService.
func (u *MissionRepository) GetAll() ([]models.Mission, error) {
	panic("unimplemented")
}

func NewMissionService(db *pgxpool.Pool) services.MissionService {
	return &MissionRepository{
		db: db,
	}
}

func (u *MissionRepository) Get(id int) (*models.Mission, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, `SELECT id, created_at, code where mission_id = $1;`, id)

	if err != nil {
		return nil, err
	}

	services, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Mission])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorMissionNotFound
		}

		return nil, err
	}

	return &services, nil
}

func (u *MissionRepository) Create(s *models.Mission) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return err
	}

	_, err = conn.Exec(ctx, `insert into missions.mission (code)
	values ($1)`, s.Code)

	if err != nil {
		return models.ErrorMissionNotCreated
	}

	return nil
}

func (u *MissionRepository) Update(s *models.Mission) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return err
	}

	_, err = conn.Exec(ctx, `UPDATE missions.mission
	SET code = $1
	WHERE id = $2`, s.Code, s.Id)

	if err != nil {
		return err
	}

	return nil
}

func (u *MissionRepository) Delete(id int) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return err
	}

	_, err = conn.Exec(ctx, "delete from missions.mission where id = $1", id)

	if err != nil {
		return models.ErrorMissionServiceNotDeleted
	}

	return nil
}
