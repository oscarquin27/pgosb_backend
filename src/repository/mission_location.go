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

const selectMissionLocationQuery = "SELECT * FROM mission.locations WHERE id = $1"

const selectMissionLocationQuerybyServiceId = "SELECT * FROM mission.locations WHERE mission_id = $1"

const selectAllMissionLocationQuery = "SELECT * FROM mission.locations"

const insertMissionLocationQuery = `INSERT INTO mission.locations (
    
    alias,state_id, state, municipality_id, municipality, parish_id,
    parish, sector_id, sector, urb_id, urb,  address
)
VALUES (
     @alias, 
    @state_id, @state, @municipality_id, @municipality, @parish_id, 
    @parish, @sector_id, @sector, @urb_id, @urb,  @address
) RETURNING id`

const updateMissionLocationQuery = `UPDATE mission.locations
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
    address = @address
WHERE id = @id; `

const deleteMissionLocationQuery = `DELETE FROM mission.locations WHERE id = $1`

func (u *MissionLocationRepository) Get(id int64) *results.ResultWithValue[*models.MissionLocation] {
	r := u.AbstractRepository.Get(id, selectMissionLocationQuery)

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)

}
func (u *MissionLocationRepository) GetAll() ([]models.MissionLocation, *results.GeneralError) {

	var states []models.MissionLocation = make([]models.MissionLocation, 0)

	values, err := u.AbstractRepository.GetAll(selectAllMissionLocationQuery)

	if err != nil {
		return states, err
	}

	return values, nil
}

func (u *MissionLocationRepository) Create(state *models.MissionLocation) *results.ResultWithValue[*models.MissionLocation] {

	r := u.AbstractRepository.Create(*state, insertMissionLocationQuery, state.GetNameArgs(), state.SetId)

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *MissionLocationRepository) Update(state *models.MissionLocation) *results.ResultWithValue[*models.MissionLocation] {
	r := u.AbstractRepository.Update(*state, updateMissionLocationQuery, state.GetNameArgs())

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *MissionLocationRepository) Delete(id int64) *results.Result {

	return u.AbstractRepository.Delete(id, deleteMissionLocationQuery)
}

func (u *MissionLocationRepository) GetLocationsByServiceId(id int64) *results.ResultWithValue[[]models.MissionLocation] {

	defaultList := make([]models.MissionLocation, 0)

	rest := results.NewResultWithValue("GetLocationByServiceId", false, defaultList, nil)

	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return rest.WithError(
			results.NewError("no se pudo adquirir conexion", err))
	}

	defer conn.Release()

	rows, err := conn.Query(ctx, selectMissionLocationQuerybyServiceId)

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
