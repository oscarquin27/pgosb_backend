package repository

import (
	"fdms/src/models"
	"fdms/src/services"
	"fdms/src/utils/results"

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

const selectMissionLocationQuery = "SELECT * FROM hq.stations WHERE id = $1"

const selectAllMissionLocationQuery = "SELECT * FROM hq.stations"

const insertMissionLocationQuery = `INSERT INTO hq.stations (
    
    id,name, description, abbreviation, phones, region_id, 
    state_id, state, municipality_id, municipality, parish_id,
    parish, sector_id, sector, urb_id, urb, street, address
)
VALUES (
    @id, @name, @description, @abbreviation, @phones, @region_id,
    @state_id, @state, @municipality_id, @municipality, @parish_id, 
    @parish, @sector_id, @sector, @urb_id, @urb, @street, @address
) RETURNING id`

const updateMissionLocationQuery = `UPDATE hq.stations
SET 
    name = @name, 
    description = @description, 
    abbreviation = @abbreviation,
    phones = @phones,
    region_id = @region_id,
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
    street = @street,
    address = @address
WHERE id = @id; `

const deleteMissionLocationQuery = `DELETE FROM hq.stations WHERE id = $1`

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

	return rest
}
