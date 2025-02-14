package repository

import (
	"context"
	"fdms/src/models"
	"fdms/src/services"
	"fdms/src/utils/results"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MissionUnitRepository struct {
	*AbstractRepository[models.MissionUnit]
}

func NewMissionUnitService(db *pgxpool.Pool) services.MissionUnitService {

	abstractImplent := NewAbstractRepository[models.MissionUnit](db)

	return &MissionUnitRepository{
		&abstractImplent,
	}
}

const selectMissionUnitQuery = "SELECT * FROM missions.units WHERE id = $1"
const selectMissionUnitQueryTemplate = "SELECT * FROM missions.units_template WHERE id = $1"

const selectAllMissionUnitQuery = "SELECT * FROM missions.units"
const selectAllMissionUnitQueryTemplate = "SELECT * FROM missions.units"

const insertMissionUnitQuery = `INSERT INTO missions.units(
	 mission_id, unit_id)
	VALUES ( @mission_id, @unit_id) RETURNING id`
const insertMissionUnitQueryTemplate = `INSERT INTO missions.units_template(
		mission_id, unit_id)
	   VALUES ( @mission_id, @unit_id) RETURNING id`

const updateMissionUnitQuery = `UPDATE missions.units
	SET   unit_id=@unit_id, mission_id=@mission_id
	WHERE id = @id `
const updateMissionUnitQueryTemplate = `UPDATE missions.units_template
	SET   unit_id=@unit_id, mission_id=@mission_id
	WHERE id = @id `

const deleteMissionUnitQuery = `DELETE FROM missions.units WHERE id = $1`
const deleteMissionUnitQueryTemplate = `DELETE FROM missions.units_template WHERE id = $1`

const selectMissionUnitByMissionIdQuery = `SELECT * FROM missions.vw_units_mission WHERE mission_id = $1`
const selectMissionUnitByMissionIdQueryTemplate = `SELECT * FROM missions.vw_units_mission_template WHERE mission_id = $1`

func (u *MissionUnitRepository) Get(id int64, isTemplate bool) *results.ResultWithValue[*models.MissionUnit] {

	query := selectMissionUnitQuery

	if isTemplate {
		query = selectMissionUnitQueryTemplate
	}

	r := u.AbstractRepository.Get(id, query, false)

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)

}
func (u *MissionUnitRepository) GetAll(isTemplate bool, params ...string) ([]models.MissionUnit, *results.GeneralError) {

	var MissionUnits []models.MissionUnit = make([]models.MissionUnit, 0)

	query := selectAllMissionUnitQuery

	if isTemplate {
		query = selectAllMissionUnitQueryTemplate
	}

	values, err := u.AbstractRepository.GetAll(query, false, params...)

	if err != nil {
		return MissionUnits, err
	}

	return values, nil
}

func (u *MissionUnitRepository) Create(MissionUnit *models.MissionUnit, isTemplate bool) *results.ResultWithValue[*models.MissionUnit] {

	query := insertMissionUnitQuery

	if isTemplate {
		query = insertMissionUnitQueryTemplate
	}

	r, id := u.AbstractRepository.Create(*MissionUnit, query, MissionUnit.GetNameArgs(), MissionUnit.SetId, false)

	MissionUnit.Id = id

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)

}

func (u *MissionUnitRepository) Update(MissionUnit *models.MissionUnit, isTemplate bool) *results.ResultWithValue[*models.MissionUnit] {

	query := updateMissionUnitQuery

	if isTemplate {
		query = updateMissionUnitQueryTemplate
	}

	r := u.AbstractRepository.Update(*MissionUnit, query, MissionUnit.GetNameArgs(), false)

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *MissionUnitRepository) Delete(id int64, isTemplate bool) *results.Result {

	query := deleteMissionUnitQuery

	if isTemplate {
		query = deleteMissionUnitQueryTemplate
	}

	return u.AbstractRepository.Delete(id, query, false)
}

func (u *MissionUnitRepository) GetByMissionId(id int, isTemplate bool) ([]models.MissionUnitSummary, error) {

	ctx := context.Background()
	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return nil, err
	}

	defer conn.Release()

	query := selectMissionUnitByMissionIdQuery

	if isTemplate {
		query = selectMissionUnitByMissionIdQueryTemplate
	}

	rows, err := conn.Query(ctx, query, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.MissionUnitSummary])

	if err != nil {
		return nil, err
	}

	return users, nil

}
