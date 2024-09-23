package repository

import (
	"context"
	"fdms/src/models"
	"fdms/src/services"
	"fdms/src/utils/results"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MissionAuthorityPersonRepository struct {
	*AbstractRepository[models.MissionAuthorityPerson]
}

func NewMissionAuthorityPersonService(db *pgxpool.Pool) services.MissionAuthorityPersonService {

	abstractImplent := NewAbstractRepository[models.MissionAuthorityPerson](db)

	return &MissionAuthorityPersonRepository{
		&abstractImplent,
	}
}

const selectMissionAuthorityPersonQuery = "SELECT * FROM missions.authorities_person WHERE id = $1"

const selectMissionAuthorityPersonQuerybyAuthorityId = "SELECT * FROM missions.authorities_person WHERE authority_id = $1"

const selectAllMissionAuthorityPersonQuery = "SELECT * FROM missions.authorities_person"

const insertMissionAuthorityPersonQuery = `INSERT INTO missions.authorities_person ( 
	mission_id,
	authority_id,
    name,

    last_name,
    legal_id,
    identification_number,
    phone,
    gender,
    observations
    )
VALUES (
	@mission_id,
	@authority_id, 

	@name,
	@last_name,
	@legal_id,
	@identification_number,
	@phone,
	@gender,
	@observations

) RETURNING id`

const updateMissionAuthorityPersonQuery = `UPDATE missions.authorities_person
SET 
	name = @name,
	last_name = @last_name,
	legal_id = @legal_id,
	identification_number = @identification_number,
	phone = @phone,
	gender = @gender,
	observations = @observations
WHERE id = @id; `

const deleteMissionAuthorityPersonQuery = `DELETE FROM missions.authorities_person WHERE id = $1`

func (u *MissionAuthorityPersonRepository) Get(id int64) *results.ResultWithValue[*models.MissionAuthorityPerson] {
	r := u.AbstractRepository.Get(id, selectMissionAuthorityPersonQuery)

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)

}
func (u *MissionAuthorityPersonRepository) GetAll(params ...string) ([]models.MissionAuthorityPerson, *results.GeneralError) {

	var states []models.MissionAuthorityPerson = make([]models.MissionAuthorityPerson, 0)

	values, err := u.AbstractRepository.GetAll(selectAllMissionAuthorityPersonQuery, params...)

	if err != nil {
		return states, err
	}

	return values, nil
}

func (u *MissionAuthorityPersonRepository) Create(state *models.MissionAuthorityPerson) *results.ResultWithValue[*models.MissionAuthorityPerson] {

	r := u.AbstractRepository.Create(*state, insertMissionAuthorityPersonQuery, state.GetNameArgs(), state.SetId)

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *MissionAuthorityPersonRepository) Update(state *models.MissionAuthorityPerson) *results.ResultWithValue[*models.MissionAuthorityPerson] {
	r := u.AbstractRepository.Update(*state, updateMissionAuthorityPersonQuery, state.GetNameArgs())

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *MissionAuthorityPersonRepository) Delete(id int64) *results.Result {

	return u.AbstractRepository.Delete(id, deleteMissionAuthorityPersonQuery)
}

func (u *MissionAuthorityPersonRepository) GetByAuthorityId(id int64) *results.ResultWithValue[[]models.MissionAuthorityPerson] {

	defaultList := make([]models.MissionAuthorityPerson, 0)

	rest := results.NewResultWithValue("GetByMissionId", false, defaultList, nil)

	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return rest.WithError(
			results.NewError("no se pudo adquirir conexion", err))
	}

	defer conn.Release()

	rows, err := conn.Query(ctx, selectMissionAuthorityPersonQuerybyAuthorityId, id)

	if err != nil {
		return rest.WithError(
			results.NewError("no se pudo ejecutar el query", err))
	}

	defer rows.Close()

	registers, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.MissionAuthorityPerson])

	if err != nil {
		if err == pgx.ErrNoRows {
			return rest.WithError(results.NewNotFoundError("no encontraron registros", err))
		}

		return rest.WithError(
			results.NewError("no se pudo ejecutar el query", err))
	}

	return rest.WithValue(registers).Success()
}
