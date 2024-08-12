package repository

import (
	"fdms/src/models"
	"fdms/src/services"
	"fdms/src/utils/results"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SectorRepository struct {
	*AbstractRepository[models.Sector]
}

func NewSectorService(db *pgxpool.Pool) services.SectorService {

	abstractImplent := NewAbstractRepository[models.Sector](db)

	return &SectorRepository{
		&abstractImplent,
	}
}

const selectSectorQuery = "SELECT id, parish_id, name  FROM locations.sectors WHERE id = $1"

const selectAllSectorQuery = "SELECT id, parish_id, name FROM locations.sectors"

const insertSectorQuery = `INSERT INTO locations.sectors(parish_id, name)
	VALUES (@parish_id, @name)`

const updateSectorQuery = `UPDATE locations.sectors
	SET parish_id=@parish_id,  name=@name
	WHERE id = @id `

const deleteSectorQuery = `DELETE FROM locations.sectors WHERE id = $1`

func (u *SectorRepository) Get(id int64) *results.ResultWithValue[*models.Sector] {

	r := u.AbstractRepository.Get(id, selectSectorQuery)

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *SectorRepository) GetAll() ([]models.Sector, *results.GeneralError) {

	var municipalities []models.Sector = make([]models.Sector, 0)

	values, err := u.AbstractRepository.GetAll(selectAllSectorQuery)

	if err != nil {
		return municipalities, err
	}

	return values, nil
}

func (u *SectorRepository) Create(state *models.Sector) *results.ResultWithValue[*models.Sector] {

	r := u.AbstractRepository.Create(*state, insertSectorQuery, state.GetNameArgs(), state.SetId)

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *SectorRepository) Update(state *models.Sector) *results.ResultWithValue[*models.Sector] {

	r := u.AbstractRepository.Update(*state, updateSectorQuery, state.GetNameArgs())

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *SectorRepository) Delete(id int64) *results.Result {

	return u.AbstractRepository.Delete(id, deleteSectorQuery)
}
