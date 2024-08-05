package repository

import (
	"context"
	"fdms/src/mikro"
	"fdms/src/models"
	"fdms/src/services"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MissionServiceRepository struct {
	db *pgxpool.Pool
}

func NewMissionServiceService(db *pgxpool.Pool) services.MissionServiceService {
	return &MissionServiceRepository{
		db: db,
	}
}

func (u *MissionServiceRepository) Get(id int) (*models.MissionService, error) {
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
	FROM missions.services where id = $1;`, id)

	if err != nil {
		return nil, err
	}

	service, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.MissionService])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorMissionServiceNotFound
		}

		return nil, err
	}

	return &service, nil
}

func (u *MissionServiceRepository) GetAll() ([]models.MissionService, error) {
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
	FROM missions.services order by id desc;`)

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

func (u *MissionServiceRepository) GetByMissionId(id int) ([]models.MissionService, error) {
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

func (u *MissionServiceRepository) Create(s *models.MissionService) (*models.MissionService, error) {
	m := mikro.NewMkModel(u.db)

	err := m.Model(s).Omit("id").
		Returning().InsertReturning("missions.services")

	if err != nil {
		return nil, err
	}

	if s.Id.Int32 >= 0 {
		return s, nil
	}

	return nil, models.ErrorMissionServiceNotCreated
}

func (u *MissionServiceRepository) Update(s *models.MissionService) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, `UPDATE missions.services
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

	if rows.RowsAffected() == 0 {
		return models.ErrorMissionServiceNotUpdated
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
