package repository

import (
	"fdms/src/models"
	"fdms/src/services"
	"fdms/src/utils/results"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UrbanizationRepository struct {
	*AbstractRepository[models.Urbanization]
}

func NewUrbanizationService(db *pgxpool.Pool) services.UrbanizationService {

	abstractImplent := NewAbstractRepository[models.Urbanization](db)

	return &UrbanizationRepository{
		&abstractImplent,
	}
}

const selectUrbanizationQuery = "SELECT id, sector_id, name  FROM locations.urbanization WHERE id = $1"

const selectAllUrbanizationQuery = "SELECT id, sector_id, name FROM locations.urbanization"

const insertUrbanizationQuery = `INSERT INTO locations.urbanization(sector_id, name)
	VALUES (@sector_id, @name)`

const updateUrbanizationQuery = `UPDATE locations.urbanization
	SET sector_id=@sector_id,  name=@name
	WHERE id = @id `

const deleteUrbanizationQuery = `DELETE FROM locations.urbanization WHERE id = $1`

func (u *UrbanizationRepository) Get(id int64) *results.ResultWithValue[*models.Urbanization] {

	r := u.AbstractRepository.Get(id, selectUrbanizationQuery)

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *UrbanizationRepository) GetAll(params ...string) ([]models.Urbanization, *results.GeneralError) {

	var municipalities []models.Urbanization = make([]models.Urbanization, 0)

	values, err := u.AbstractRepository.GetAll(selectAllUrbanizationQuery, params...)

	if err != nil {
		return municipalities, err
	}

	return values, nil
}

func (u *UrbanizationRepository) Create(state *models.Urbanization) *results.ResultWithValue[*models.Urbanization] {

	r := u.AbstractRepository.Create(*state, insertUrbanizationQuery, state.GetNameArgs(), state.SetId)

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *UrbanizationRepository) Update(state *models.Urbanization) *results.ResultWithValue[*models.Urbanization] {

	r := u.AbstractRepository.Update(*state, updateUrbanizationQuery, state.GetNameArgs())

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *UrbanizationRepository) Delete(id int64) *results.Result {

	return u.AbstractRepository.Delete(id, deleteUrbanizationQuery)
}
