package repository

import (
	"context"
	"fdms/src/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MissionServiceRepository struct {
	db   *pgxpool.Pool
}

func NewMissionServiceService(db *pgxpool.Pool) *MissionServiceRepository {
	return &MissionServiceRepository{
		db:   db,
	}
}


func (u *MissionServiceRepository) Get(id int) ([]models.MissionService, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, `SELECT id, mission_id, 
	antares_id, 
	units, 
	bombers, 
	summary, 
	description
	FROM missions.services where mission_id = $1;`, id)

	if err != nil {
		return nil, err
	}

	services, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.MissionService])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorMissionServiceNotFound
		}

		return nil, err
	}

	return services, nil
}


func (u *MissionServiceRepository) Create(s *models.MissionService) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return err
	}

	_, err = conn.Exec(ctx, `insert into missions.person (mission_id, 
	antares_id, 
	units, 
	bombers, 
	summary, 
	description)
	values ($1, $2, $3, $4, $5, $6, $7)`, s.MissionId, s.AntaresId, s.Units, s.Bombers, s.Summary, s.Description)

	if err != nil {
		return models.ErrorMissionServiceNotCreated
	}

	return nil
}

func (u *MissionServiceRepository) Update(s *models.MissionService) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return err
	}

	_, err = conn.Exec(ctx, `UPDATE missions.services
	SET mission_id = $1, 
	antares_id = $2, 
	units = $3, 
	bombers = $4, 
	summary = $5, 
	description = $6
	WHERE id = $7`, s.MissionId, s.AntaresId, s.Units, s.Bombers, s.Summary, s.Description, s.Id)

	if err != nil {
		return err
	}


	return nil
}

func (u *MissionServiceRepository) Delete(id int) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return err
	}

	_, err = conn.Exec(ctx, "delete from missions.services where id = $1", id)

	if err != nil {
		return models.ErrorMissionServiceNotDeleted
	}


	return nil
}
