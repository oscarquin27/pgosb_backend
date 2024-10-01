package repository

import (
	"context"
	"fdms/src/models"
	"fdms/src/services"
	"fdms/src/utils/results"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MissionAuthorityRepository struct {
	*AbstractRepository[models.MissionAuthority]
}

func NewMissionAuthorityService(db *pgxpool.Pool) services.MissionAuthorityService {

	abstractImplent := NewAbstractRepository[models.MissionAuthority](db)

	return &MissionAuthorityRepository{
		&abstractImplent,
	}
}

const selectMissionAuthorityQuery = "SELECT * FROM missions.authorities WHERE id = $1"

const selectMissionAuthorityQuerybyMissionId = "SELECT * FROM missions.authorities WHERE mission_id = $1"

const selectMissionAuthorityQuerybyMissionSummary = "SELECT * FROM missions.vw_mission_authority_summary WHERE mission_id = $1"

const selectAllMissionAuthorityQuery = "SELECT * FROM missions.authorities"

const insertMissionAuthorityQuery = `INSERT INTO missions.authorities (
    alias, mission_id, institution_id
)
VALUES (
     @alias, 
    @mission_id,
    @institution_id
) RETURNING id`

const updateMissionAuthorityQuery = `UPDATE missions.authorities
SET 
    alias = @alias, 
    
    institution_id = @institution_id
WHERE id = @id; `

const deleteMissionAuthorityQuery = `DELETE FROM missions.authorities WHERE id = $1`

func (u *MissionAuthorityRepository) Get(id int64) *results.ResultWithValue[*models.MissionAuthority] {
	r := u.AbstractRepository.Get(id, selectMissionAuthorityQuery)

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)

}
func (u *MissionAuthorityRepository) GetAll(params ...string) ([]models.MissionAuthority, *results.GeneralError) {

	var states []models.MissionAuthority = make([]models.MissionAuthority, 0)

	values, err := u.AbstractRepository.GetAll(selectAllMissionAuthorityQuery, params...)

	if err != nil {
		return states, err
	}

	return values, nil
}

func (u *MissionAuthorityRepository) Create(state *models.MissionAuthority) *results.ResultWithValue[*models.MissionAuthority] {

	r := u.AbstractRepository.Create(*state, insertMissionAuthorityQuery, state.GetNameArgs(), state.SetId)

	r.Value = *state

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *MissionAuthorityRepository) Update(state *models.MissionAuthority) *results.ResultWithValue[*models.MissionAuthority] {
	r := u.AbstractRepository.Update(*state, updateMissionAuthorityQuery, state.GetNameArgs())

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *MissionAuthorityRepository) Delete(id int64) *results.Result {

	return u.AbstractRepository.Delete(id, deleteMissionAuthorityQuery)
}

func (u *MissionAuthorityRepository) GetByMissionId(id int64) *results.ResultWithValue[[]models.MissionAuthority] {

	defaultList := make([]models.MissionAuthority, 0)

	rest := results.NewResultWithValue("GetByMissionId", false, defaultList, nil)

	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return rest.WithError(
			results.NewError("no se pudo adquirir conexion", err))
	}

	defer conn.Release()

	rows, err := conn.Query(ctx, selectMissionAuthorityQuerybyMissionId, id)

	if err != nil {
		return rest.WithError(
			results.NewError("no se pudo ejecutar el query", err))
	}

	defer rows.Close()

	registers, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.MissionAuthority])

	if err != nil {
		if err == pgx.ErrNoRows {
			return rest.WithError(results.NewNotFoundError("no encontraron registros", err))
		}

		return rest.WithError(
			results.NewError("no se pudo ejecutar el query", err))
	}

	return rest.WithValue(registers).Success()
}

func (u *MissionAuthorityRepository) GetSummaryByMissionId(id int64) *results.ResultWithValue[[]models.MissionAuthoritySummary] {

	defaultList := make([]models.MissionAuthoritySummary, 0)

	rest := results.NewResultWithValue("GetSummaryByMissionId", false, defaultList, nil)

	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return rest.WithError(
			results.NewError("no se pudo adquirir conexion", err))
	}

	defer conn.Release()

	rows, err := conn.Query(ctx, selectMissionAuthorityQuerybyMissionSummary, id)

	if err != nil {
		return rest.WithError(
			results.NewError("no se pudo ejecutar el query", err))
	}

	defer rows.Close()

	registers, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.MissionAuthoritySummary])

	if err != nil {

		if err == pgx.ErrNoRows {
			return rest.WithError(results.NewNotFoundError("no encontraron registros", err))
		}

		return rest.WithError(
			results.NewError("no se pudo ejecutar el query", err))
	}

	return rest.WithValue(registers).Success()

}
