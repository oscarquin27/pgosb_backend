package repository

import (
	"fdms/src/models"
	"fdms/src/services"
	"fdms/src/utils/results"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthorityRepository struct {
	*AbstractRepository[models.Authority]
}

func NewAuthorityService(db *pgxpool.Pool) services.AuthorityService {

	abstractImplent := NewAbstractRepository[models.Authority](db)

	return &AuthorityRepository{
		&abstractImplent,
	}
}

const selectAuthorityQuery = "SELECT * FROM authorities.authority WHERE id = $1"

const selectAllAuthorityQuery = "SELECT * FROM authorities.authority"

const insertAuthorityQuery = `INSERT INTO authorities.authority (
    name,
    abbreviation,
    government
)
VALUES (
    @name,
    @abbreviation,
    @government
) RETURNING id`

const updateAuthorityQuery = `UPDATE authorities.authority
SET 
    name = @name,
    abbreviation = @abbreviation,
    government = @government
WHERE id = @id`

const deleteAuthorityQuery = `DELETE FROM authorities.authority WHERE id = $1`

func (u *AuthorityRepository) Get(id int64, isTemplate bool) *results.ResultWithValue[*models.Authority] {
	r := u.AbstractRepository.Get(id, selectAuthorityQuery, false)

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)

}
func (u *AuthorityRepository) GetAll(isTemplate bool, params ...string) ([]models.Authority, *results.GeneralError) {

	var states []models.Authority = make([]models.Authority, 0)

	values, err := u.AbstractRepository.GetAll(selectAllAuthorityQuery, false, params...)

	if err != nil {
		return states, err
	}

	return values, nil
}

func (u *AuthorityRepository) Create(state *models.Authority, isTemplate bool) *results.ResultWithValue[*models.Authority] {

	r, id := u.AbstractRepository.Create(*state, insertAuthorityQuery, state.GetNameArgs(), state.SetId, false)

	state.Id = id

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)

}

func (u *AuthorityRepository) Update(state *models.Authority, isTemplate bool) *results.ResultWithValue[*models.Authority] {
	r := u.AbstractRepository.Update(*state, updateAuthorityQuery, state.GetNameArgs(), false)

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *AuthorityRepository) Delete(id int64, isTemplate bool) *results.Result {

	return u.AbstractRepository.Delete(id, deleteAuthorityQuery, false)
}
