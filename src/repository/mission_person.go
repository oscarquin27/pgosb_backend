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

// GetAll implements services.MissionPersonService.
func (u *MissionPersonRepository) GetAll() ([]models.MissionPerson, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, `SELECT id, service_id, 
	unit_id, 
	infrastructure_id, 
	vehicle_id, 
	first_name, 
	last_name, 
	age, 
	gender, 
	legal_id, 
	phone, 
	employment,
	state,
	municipality, 
	parish, 
	address, 
	pathology, 
	observations, 
	condition
FROM missions.person where;`)

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

// GetByServiceId implements services.MissionPersonService.
func (u *MissionPersonRepository) GetByServiceId(id int) ([]models.MissionPerson, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, `SELECT id, service_id, 
	unit_id, 
	infrastructure_id, 
	vehicle_id, 
	first_name, 
	last_name, 
	age, 
	gender, 
	legal_id, 
	phone, 
	employment,
	state,
	municipality, 
	parish, 
	address, 
	pathology, 
	observations, 
	condition
FROM missions.person where service_id = $1;`, id)

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

	rows, err := conn.Query(ctx, `SELECT id, service_id, 
	unit_id, 
	infrastructure_id, 
	vehicle_id, 
	first_name, 
	last_name, 
	age, 
	gender, 
	legal_id, 
	phone, 
	employment,
	state,
	municipality, 
	parish, 
	address, 
	pathology, 
	observations, 
	condition
FROM missions.person where id = $1;`, id)

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

	rows, err := m.Model(p).Omit("id").Where("id", "=", p.Id).Update("missions.person")

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

	defer conn.Release()

	if err != nil {
		return err
	}

	_, err = conn.Exec(ctx, "delete from missions.person where id = $1", id)

	if err != nil {
		return models.ErrorMissionPersonNotDeleted
	}

	return nil
}
