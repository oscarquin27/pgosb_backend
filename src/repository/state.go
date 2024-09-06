package repository

import (
	"fdms/src/models"
	"fdms/src/services"
	"fdms/src/utils/results"

	"github.com/jackc/pgx/v5/pgxpool"
)

type StateRepository struct {
	*AbstractRepository[models.State]
}

func NewStateService(db *pgxpool.Pool) services.StateService {

	abstractImplent := NewAbstractRepository[models.State](db)

	return &StateRepository{
		&abstractImplent,
	}
}

const selectStateQuery = "SELECT id, name, coordinates , capital FROM locations.states WHERE id = $1"

const selectAllStateQuery = "SELECT id, name, coordinates , capital FROM locations.states"

const insertStateQuery = `INSERT INTO locations.states(
	name, coordinates, capital)
	VALUES (@name, @coordinates, @capital)`

const updateStateQuery = `UPDATE locations.states
	SET  name=@name, coordinates=@coordinates, capital=@capital
	WHERE id = @id `

const deleteStateQuery = `DELETE FROM locations.states WHERE id = $1`

func (u *StateRepository) Get(id int64) *results.ResultWithValue[*models.State] {
	r := u.AbstractRepository.Get(id, selectStateQuery)

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)

}
func (u *StateRepository) GetAll(params ...string) ([]models.State, *results.GeneralError) {

	var states []models.State = make([]models.State, 0)

	values, err := u.AbstractRepository.GetAll(selectAllStateQuery, params...)

	if err != nil {
		return states, err
	}

	return values, nil
}

func (u *StateRepository) Create(state *models.State) *results.ResultWithValue[*models.State] {

	r := u.AbstractRepository.Create(*state, insertStateQuery, state.GetNameArgs(), state.SetId)

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *StateRepository) Update(state *models.State) *results.ResultWithValue[*models.State] {
	r := u.AbstractRepository.Update(*state, updateStateQuery, state.GetNameArgs())

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *StateRepository) Delete(id int64) *results.Result {

	return u.AbstractRepository.Delete(id, deleteStateQuery)
}
