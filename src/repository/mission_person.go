package repository

import (
	"context"
	"fdms/src/mikro"
	"fdms/src/models"
	"fdms/src/services"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MissionPersonRepository struct {
	db *pgxpool.Pool
}

func NewMissionPersonService(db *pgxpool.Pool) services.MissionPersonService {
	return &MissionPersonRepository{
		db: db,
	}
}

func (u *MissionPersonRepository) GetMissionId(id int) ([]models.MissionPerson, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return nil, err
	}

	defer conn.Release()

	rows, err := conn.Query(ctx, `
	SELECT *
	FROM missions.person WHERE mission_id = $1;`, id)

	if err != nil {
		return nil, err
	}

	person, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.MissionPerson])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorMissionPersonNotFound
		}

		return nil, err
	}

	return person, nil
}

func (u *MissionPersonRepository) Get(id int) (*models.MissionPerson, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, `
	SELECT *
    FROM missions.person WHERE id = $1;`, id)

	if err != nil {
		return nil, err
	}

	person, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.MissionPerson])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorMissionPersonNotFound
		}

		return nil, err
	}

	return &person, nil
}

func (u *MissionPersonRepository) GetAll() ([]models.MissionPerson, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, `
	SELECT *
	FROM missions.person`)

	if err != nil {
		return nil, err
	}

	person, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.MissionPerson])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorMissionPersonNotFound
		}

		return nil, err
	}

	return person, nil
}

func (u *MissionPersonRepository) Create(p *models.MissionPerson) error {
	m := mikro.NewMkModel(u.db)

	rows, err := m.Model(p).Omit("id").Insert("missions.person")

	if err != nil {
		return err
	}

	if rows > 0 {
		return nil
	}

	return models.ErrorMissionPersonNotCreated
}

func (u *MissionPersonRepository) Update(p *models.MissionPerson) error {
	m := mikro.NewMkModel(u.db)

	rows, err := m.Model(p).Omit("id").Omit("mission_id").Where("id", "=", p.Id).Update("missions.person")

	if err != nil {
		return err
	}

	if rows > 0 {
		return nil
	}

	return models.ErrorMissionPersonNotUpdated
}

func (u *MissionPersonRepository) Delete(id int) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return err
	}

	defer conn.Release()

	_, err = conn.Exec(ctx, "delete from missions.person where id = $1", id)

	if err != nil {
		return models.ErrorMissionPersonNotDeleted
	}

	return nil
}
