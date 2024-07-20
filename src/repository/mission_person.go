package repository

import (
	"context"
	"fdms/src/models"
	"fdms/src/services"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MissionPersonRepository struct {
	db   *pgxpool.Pool
}

func NewMissionPersonService(db *pgxpool.Pool) services.MissionPersonService {
	return &MissionPersonRepository{
		db:   db,
	}
}

func (u *MissionPersonRepository) Get(id int) ([]models.MissionPerson, error) {
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


func (u *MissionPersonRepository) Create(p *models.MissionPerson) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return err
	}

	_, err = conn.Exec(ctx, `insert into missions.person (service_id, 
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
	person_condition)
	values (
	$1,
	$2,
	$3,
	$4,
	$5,
	$6,
	$7,
	$8,
	$9,
	$10,
	$11,
	$12,
	$13,
	$14,
	$15,
	$16,
	$17,
	$18)`, p.ServiceId, p.UnitId, p.InfrastructureId, p.VehicleId, p.FirstName, p.LastName, p.Age, p.Gender, p.LegalId, p.Phone, p.Employment, p.State, p.Municipality, p.Parish, p.Address, p.Pathology, p.Observations, p.Condition)

	if err != nil {
		return err
	}

	return nil
}

func (u *MissionPersonRepository) Update(p *models.MissionPerson) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, `UPDATE missions.person
	SET service_id = $1, 
	    unit_id = $2, 
	    infrastructure_id = $3, 
	    vehicle_id = $4, 
	    first_name = $5, 
	    last_name = $6, 
	    age = $7, 
	    gender = $8, 
	    legal_id = $9, 
	    phone = $10, 
	    employment = $11,
	    state = $12,
	    municipality = $13, 
	    parish = $14, 
	    address = $15, 
	    pathology = $16, 
	    observations = $17, 
	    person_condition = $18
		WHERE id = $19`, p.ServiceId, p.UnitId, p.InfrastructureId, p.VehicleId, p.FirstName, p.LastName, p.Age, p.Gender, p.LegalId, p.Phone, p.Employment, p.State, p.Municipality, p.Parish, p.Address, p.Pathology, p.Observations, p.Condition, p.Id)

	if err != nil {
		return err
	}

	if rows.RowsAffected() == 0 {
		return models.ErrorMissionPersonNotUpdated
	}
	return nil
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
