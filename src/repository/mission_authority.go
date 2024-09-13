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

const selectMissionAuthorityQuery = "SELECT * FROM missions.locations WHERE id = $1"

const selectMissionAuthorityQuerybyServiceId = "SELECT * FROM missions.locations WHERE mission_id = $1"

const selectAllMissionAuthorityQuery = "SELECT * FROM missions.locations"

const insertMissionAuthorityQuery = `INSERT INTO missions.locations (
    
    alias,state_id, state, municipality_id, municipality, parish_id,
    parish, sector_id, sector, urb_id, urb,  address , mission_id
)
VALUES (
     @alias, 
    @state_id, @state, @municipality_id, @municipality, @parish_id, 
    @parish, @sector_id, @sector, @urb_id, @urb,  @address, @mission_id
) RETURNING id`

const updateMissionAuthorityQuery = `UPDATE missions.locations
SET 
    alias = @alias, 
    
    state_id = @state_id,
    state = @state,
    municipality_id = @municipality_id,
    municipality = @municipality,
    parish_id = @parish_id,
    parish = @parish,
    sector_id = @sector_id,
    sector = @sector,
    urb_id = @urb_id,
    urb = @urb,
    address = @address,
	mission_id = @mission_id
WHERE id = @id; `

const deleteMissionAuthorityQuery = `DELETE FROM missions.locations WHERE id = $1`

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

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *MissionAuthorityRepository) Update(state *models.MissionAuthority) *results.ResultWithValue[*models.MissionAuthority] {
	r := u.AbstractRepository.Update(*state, updateMissionAuthorityQuery, state.GetNameArgs())

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *MissionAuthorityRepository) Delete(id int64) *results.Result {

	return u.AbstractRepository.Delete(id, deleteMissionAuthorityQuery)
}

func (u *MissionAuthorityRepository) GetByServiceId(id int64) *results.ResultWithValue[[]models.MissionAuthority] {

	defaultList := make([]models.MissionAuthority, 0)

	rest := results.NewResultWithValue("GetLocationByServiceId", false, defaultList, nil)

	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return rest.WithError(
			results.NewError("no se pudo adquirir conexion", err))
	}

	defer conn.Release()

	rows, err := conn.Query(ctx, selectMissionAuthorityQuerybyServiceId, id)

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
