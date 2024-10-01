package repository

import (
	"fdms/src/models"
	"fdms/src/services"
	"fdms/src/utils/results"

	"github.com/jackc/pgx/v5/pgxpool"
)

type HealthcareCenterRepository struct {
	*AbstractRepository[models.HealthcareCenter]
}

func NewHealthcareCenterService(db *pgxpool.Pool) services.HealthcareCenterService {

	abstractImplent := NewAbstractRepository[models.HealthcareCenter](db)

	return &HealthcareCenterRepository{
		&abstractImplent,
	}
}

const selectHealthcareCenterQuery = "SELECT * FROM hq.centers WHERE id = $1 ORDER BY id ASC"

const selectAllHealthcareCenterQuery = "SELECT * FROM hq.centers ORDER BY id ASC"

const insertHealthcareCenterQuery = `INSERT INTO hq.centers (
    
    id,name, description, abbreviation, phones, region_id, 
    state_id, state, municipality_id, municipality, parish_id,
    parish, sector_id, sector, urb_id, urb, street, address
)
VALUES (
    @id, @name, @description, @abbreviation, @phones, @region_id,
    @state_id, @state, @municipality_id, @municipality, @parish_id, 
    @parish, @sector_id, @sector, @urb_id, @urb, @street, @address
) RETURNING id`

const updateHealthcareCenterQuery = `UPDATE hq.centers
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

const deleteHealthcareCenterQuery = `DELETE FROM hq.centers WHERE id = $1`

func (u *HealthcareCenterRepository) Get(id int64) *results.ResultWithValue[*models.HealthcareCenter] {
	r := u.AbstractRepository.Get(id, selectHealthcareCenterQuery)

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)

}
func (u *HealthcareCenterRepository) GetAll(params ...string) ([]models.HealthcareCenter, *results.GeneralError) {

	var states []models.HealthcareCenter = make([]models.HealthcareCenter, 0)

	values, err := u.AbstractRepository.GetAll(selectAllHealthcareCenterQuery, params...)

	if err != nil {
		return states, err
	}

	return values, nil
}

func (u *HealthcareCenterRepository) Create(state *models.HealthcareCenter) *results.ResultWithValue[*models.HealthcareCenter] {

	r := u.AbstractRepository.Create(*state, insertHealthcareCenterQuery, state.GetNameArgs(), state.SetId)

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *HealthcareCenterRepository) Update(state *models.HealthcareCenter) *results.ResultWithValue[*models.HealthcareCenter] {
	r := u.AbstractRepository.Update(*state, updateHealthcareCenterQuery, state.GetNameArgs())

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *HealthcareCenterRepository) Delete(id int64) *results.Result {

	return u.AbstractRepository.Delete(id, deleteHealthcareCenterQuery)
}
