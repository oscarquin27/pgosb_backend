package repository

import (
	"fdms/src/models"
	"fdms/src/services"
	"fdms/src/utils/results"

	"github.com/jackc/pgx/v5/pgxpool"
)

type StationRepository struct {
	*AbstractRepository[models.Station]
}

func NewStationService(db *pgxpool.Pool) services.StationService {

	abstractImplent := NewAbstractRepository[models.Station](db)

	return &StationRepository{
		&abstractImplent,
	}
}

const selectStationQuery = "SELECT * FROM hq.stations WHERE id = $1"

const selectAllStationQuery = "SELECT * FROM hq.stations"

const insertStationQuery = `INSERT INTO hq.stations (
    
    id,name, description, abbreviation, phones, region_id, 
    state_id, state, municipality_id, municipality, parish_id,
    parish, sector_id, sector, urb_id, urb, street, address
)
VALUES (
    @id, @name, @description, @abbreviation, @phones, @region_id,
    @state_id, @state, @municipality_id, @municipality, @parish_id, 
    @parish, @sector_id, @sector, @urb_id, @urb, @street, @address
) RETURNING id`

const updateStationQuery = `UPDATE hq.stations
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

const deleteStationQuery = `DELETE FROM hq.stations WHERE id = $1`

func (u *StationRepository) Get(id int64) *results.ResultWithValue[*models.Station] {
	r := u.AbstractRepository.Get(id, selectStationQuery)

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)

}
func (u *StationRepository) GetAll() ([]models.Station, *results.GeneralError) {

	var states []models.Station = make([]models.Station, 0)

	values, err := u.AbstractRepository.GetAll(selectAllStationQuery)

	if err != nil {
		return states, err
	}

	return values, nil
}

func (u *StationRepository) Create(state *models.Station) *results.ResultWithValue[*models.Station] {

	r := u.AbstractRepository.Create(*state, insertStationQuery, state.GetNameArgs(), state.SetId)

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *StationRepository) Update(state *models.Station) *results.ResultWithValue[*models.Station] {
	r := u.AbstractRepository.Update(*state, updateStationQuery, state.GetNameArgs())

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *StationRepository) Delete(id int64) *results.Result {

	return u.AbstractRepository.Delete(id, deleteStationQuery)
}
