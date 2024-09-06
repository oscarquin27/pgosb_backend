package repository

import (
	"fdms/src/models"
	"fdms/src/services"
	"fdms/src/utils/results"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MunicipalityRepository struct {
	*AbstractRepository[models.Municipality]
}

func NewMunicipalityService(db *pgxpool.Pool) services.MunicipalityService {

	abstractImplent := NewAbstractRepository[models.Municipality](db)

	return &MunicipalityRepository{
		&abstractImplent,
	}
}

const selectMunicipalityQuery = "SELECT id, state_id ,name, coordinates , capital FROM locations.municipalities WHERE id = $1"

const selectAllMunicipalityQuery = "SELECT id, state_id,name, coordinates , capital FROM locations.municipalities"

const insertMunicipalityQuery = `INSERT INTO locations.municipalities(
      state_id,	name, coordinates, capital)
	VALUES (@state_id, @name, @coordinates, @capital)`

const updateMunicipalityQuery = `UPDATE locations.municipalities
	SET state_id=@state_id name=@name, coordinates=@coordinates, capital=@capital
	WHERE id = @id `

const deleteMunicipalityQuery = `DELETE FROM locations.municipalities WHERE id = $1`

func (u *MunicipalityRepository) Get(id int64) *results.ResultWithValue[*models.Municipality] {

	r := u.AbstractRepository.Get(id, selectMunicipalityQuery)

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *MunicipalityRepository) GetAll(params ...string) ([]models.Municipality, *results.GeneralError) {

	var municipalities []models.Municipality = make([]models.Municipality, 0)

	values, err := u.AbstractRepository.GetAll(selectAllMunicipalityQuery, params...)

	if err != nil {
		return municipalities, err
	}

	return values, nil
}

func (u *MunicipalityRepository) Create(state *models.Municipality) *results.ResultWithValue[*models.Municipality] {

	r := u.AbstractRepository.Create(*state, insertMunicipalityQuery, state.GetNameArgs(), state.SetId)

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *MunicipalityRepository) Update(state *models.Municipality) *results.ResultWithValue[*models.Municipality] {

	r := u.AbstractRepository.Update(*state, updateMunicipalityQuery, state.GetNameArgs())

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *MunicipalityRepository) Delete(id int64) *results.Result {

	return u.AbstractRepository.Delete(id, deleteMunicipalityQuery)
}
