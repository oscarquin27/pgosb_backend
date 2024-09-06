package repository

import (
	"fdms/src/models"
	"fdms/src/services"
	"fdms/src/utils/results"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ParishRepository struct {
	*AbstractRepository[models.Parish]
}

func NewParishService(db *pgxpool.Pool) services.ParishService {

	abstractImplent := NewAbstractRepository[models.Parish](db)

	return &ParishRepository{
		&abstractImplent,
	}
}

const selectParishQuery = "SELECT id, municipality_id,state_id ,name, coordinates , capital FROM locations.parish WHERE id = $1"

const selectAllParishQuery = "SELECT id, municipality_id,state_id,name, coordinates , capital FROM locations.parish"

const insertParishQuery = `INSERT INTO locations.parish(
      state_id, municipality_id	,name, coordinates, capital)
	VALUES (@state_id, @name, @coordinates, @capital)`

const updateParishQuery = `UPDATE locations.parish
	SET municipality_id=@municipality_id, state_id=@state_id name=@name, coordinates=@coordinates, capital=@capital
	WHERE id = @id `

const deleteParishQuery = `DELETE FROM locations.parish WHERE id = $1`

func (u *ParishRepository) Get(id int64) *results.ResultWithValue[*models.Parish] {

	r := u.AbstractRepository.Get(id, selectParishQuery)

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *ParishRepository) GetAll(params ...string) ([]models.Parish, *results.GeneralError) {

	var municipalities []models.Parish = make([]models.Parish, 0)

	values, err := u.AbstractRepository.GetAll(selectAllParishQuery, params...)

	if err != nil {
		return municipalities, err
	}

	return values, nil
}

func (u *ParishRepository) Create(state *models.Parish) *results.ResultWithValue[*models.Parish] {

	r := u.AbstractRepository.Create(*state, insertParishQuery, state.GetNameArgs(), state.SetId)

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *ParishRepository) Update(state *models.Parish) *results.ResultWithValue[*models.Parish] {

	r := u.AbstractRepository.Update(*state, updateParishQuery, state.GetNameArgs())

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *ParishRepository) Delete(id int64) *results.Result {

	return u.AbstractRepository.Delete(id, deleteParishQuery)
}
