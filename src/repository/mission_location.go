package repository

import (
	"context"
	"fdms/src/models"
	"fdms/src/services"
	"fdms/src/utils/results"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MissionLocationRepository struct {
	*AbstractRepository[models.MissionLocation]
}

func NewMissionLocationService(db *pgxpool.Pool) services.MissionLocationService {

	abstractImplent := NewAbstractRepository[models.MissionLocation](db)

	return &MissionLocationRepository{
		&abstractImplent,
	}
}

const selectMissionLocationQuery = "SELECT * FROM missions.locations WHERE id = $1"
const selectMissionLocationQueryTemplate = "SELECT * FROM missions.locations_template WHERE id = $1"

const selectMissionLocationQuerybyServiceId = "SELECT * FROM missions.locations WHERE mission_id = $1"
const selectMissionLocationQuerybyServiceIdTemplate = "SELECT * FROM missions.locations_template WHERE mission_id = $1"

const selectAllMissionLocationQuery = "SELECT * FROM missions.locations"
const selectAllMissionLocationQueryTemplate = "SELECT * FROM missions.locations_template"

const insertMissionLocationQuery = `INSERT INTO missions.locations (
    
    alias,state_id, state, municipality_id, municipality, parish_id,
    parish, sector_id, sector, urb_id, urb,  address , mission_id, street, beach
)
VALUES (
     @alias, 
    @state_id, @state, @municipality_id, @municipality, @parish_id, 
    @parish, @sector_id, @sector, @urb_id, @urb,  @address, @mission_id, @street, @beach
) RETURNING id`
const insertMissionLocationQueryTemplate = `INSERT INTO missions.locations_template (
    
alias,state_id, state, municipality_id, municipality, parish_id,
parish, sector_id, sector, urb_id, urb,  address , mission_id, street, beach
)
VALUES (
 @alias, 
@state_id, @state, @municipality_id, @municipality, @parish_id, 
@parish, @sector_id, @sector, @urb_id, @urb,  @address, @mission_id, @street, @beach
) RETURNING id`

const updateMissionLocationQuery = `UPDATE missions.locations
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
	street = @street,
	beach = @beach,
	mission_id = @mission_id
WHERE id = @id; `
const updateMissionLocationQueryTemplate = `UPDATE missions.locations_template
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
	street = @street,
	beach = @beach,
	mission_id = @mission_id
WHERE id = @id; `

const deleteMissionLocationQuery = `DELETE FROM missions.locations WHERE id = $1`
const deleteMissionLocationQueryTemplate = `DELETE FROM missions.locations_template WHERE id = $1`

func (u *MissionLocationRepository) Get(id int64, isTemplate bool) *results.ResultWithValue[*models.MissionLocation] {
	query := selectMissionLocationQuery

	if isTemplate {
		query = selectMissionLocationQueryTemplate
	}

	r := u.AbstractRepository.Get(id, query, false)

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)

}
func (u *MissionLocationRepository) GetAll(isTemplate bool, params ...string) ([]models.MissionLocation, *results.GeneralError) {

	var states []models.MissionLocation = make([]models.MissionLocation, 0)

	query := selectAllMissionLocationQuery

	if isTemplate {
		query = selectAllMissionLocationQueryTemplate
	}

	values, err := u.AbstractRepository.GetAll(query, false, params...)

	if err != nil {
		return states, err
	}

	return values, nil
}

func (u *MissionLocationRepository) Create(state *models.MissionLocation, isTemplate bool) *results.ResultWithValue[*models.MissionLocation] {

	query := insertMissionLocationQuery

	if isTemplate {
		query = insertMissionLocationQueryTemplate
	}

	r, id := u.AbstractRepository.Create(*state, query, state.GetNameArgs(), state.SetId, false)

	r.Value.Id = id

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *MissionLocationRepository) Update(state *models.MissionLocation, isTemplate bool) *results.ResultWithValue[*models.MissionLocation] {

	query := updateMissionLocationQuery

	if isTemplate {
		query = updateMissionLocationQueryTemplate
	}

	r := u.AbstractRepository.Update(*state, query, state.GetNameArgs(), false)

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *MissionLocationRepository) Delete(id int64, isTemplate bool) *results.Result {

	query := deleteMissionLocationQuery

	if isTemplate {
		query = deleteMissionLocationQueryTemplate
	}

	return u.AbstractRepository.Delete(id, query, false)
}

func (u *MissionLocationRepository) GetLocationsByServiceId(id int64, isTemplate bool) *results.ResultWithValue[[]models.MissionLocation] {

	defaultList := make([]models.MissionLocation, 0)

	query := selectMissionLocationQuerybyServiceId

	if isTemplate {
		query = selectMissionLocationQuerybyServiceIdTemplate
	}

	rest := results.NewResultWithValue("GetLocationByServiceId", false, defaultList, nil)

	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return rest.WithError(
			results.NewError("no se pudo adquirir conexion", err))
	}

	defer conn.Release()

	rows, err := conn.Query(ctx, query, id)

	if err != nil {
		return rest.WithError(
			results.NewError("no se pudo ejecutar el query", err))
	}

	defer rows.Close()

	registers, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.MissionLocation])

	if err != nil {
		if err == pgx.ErrNoRows {
			return rest.WithError(results.NewNotFoundError("no encontraron registros", err))
		}

		return rest.WithError(
			results.NewError("no se pudo ejecutar el query", err))
	}

	return rest.WithValue(registers).Success()
}
