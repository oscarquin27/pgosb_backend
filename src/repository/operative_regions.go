package repository

import (
	"fdms/src/models"
	"fdms/src/services"
	"fdms/src/utils/results"

	"github.com/jackc/pgx/v5/pgxpool"
)

type OperativeRegionsRepository struct {
	*AbstractRepository[models.OperativeRegions]
}

func NewOPerativeRegionsService(db *pgxpool.Pool) services.OperativeRegionService {

	abstractImplent := NewAbstractRepository[models.OperativeRegions](db)

	return &OperativeRegionsRepository{
		&abstractImplent,
	}
}

const selectOperativeRegion = "SELECT id, description, abbreviation, phone, coverage FROM hq.operative_regions WHERE id = $1"

const selectAllOperativeRegions = "SELECT id, description, abbreviation, phone, coverage FROM hq.operative_regions"

const insertOperativeRegion = `INSERT INTO hq.operative_regions (description, abbreviation, phone, coverage) VALUES (@description, @abbreviation, @phone, @coverage) `

const updateOperativeRegion = `UPDATE hq.operative_regions
	SET description=@description, abbreviation=@abbreviation, phone=@phone, coverage=@coverage
	WHERE id = @id `

const deleteOperativeRegion = `DELETE FROM hq.operative_regions WHERE id = @id`

func (u *OperativeRegionsRepository) Get(id int64) *results.ResultWithValue[*models.OperativeRegions] {

	r := u.AbstractRepository.Get(id, selectOperativeRegion)

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *OperativeRegionsRepository) GetAll(params ...string) ([]models.OperativeRegions, *results.GeneralError) {

	var operativeRegions []models.OperativeRegions = make([]models.OperativeRegions, 0)

	values, err := u.AbstractRepository.GetAll(selectAllOperativeRegions, params...)

	if err != nil {
		return operativeRegions, err
	}

	return values, nil
}

func (u *OperativeRegionsRepository) Create(operativeRegions *models.OperativeRegions) *results.ResultWithValue[*models.OperativeRegions] {

	r := u.AbstractRepository.Create(*operativeRegions, insertOperativeRegion, operativeRegions.GetNameArgs(), operativeRegions.SetId)

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *OperativeRegionsRepository) Update(state *models.OperativeRegions) *results.ResultWithValue[*models.OperativeRegions] {

	r := u.AbstractRepository.Update(*state, updateUrbanizationQuery, state.GetNameArgs())

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *OperativeRegionsRepository) Delete(id int64) *results.Result {

	return u.AbstractRepository.Delete(id, deleteUrbanizationQuery)
}
