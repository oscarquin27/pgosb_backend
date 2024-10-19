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

const selectAllMissionUnitQuery = "SELECT * FROM missions.units"

const insertMissionUnitQuery = `INSERT INTO missions.units(
	 mission_id, unit_id)
	VALUES ( @mission_id, @unit_id) RETURNING id`

const updateMissionUnitQuery = `UPDATE missions.units
	SET   unit_id=@unit_id, mission_id=@mission_id
	WHERE id = @id `

const deleteMissionUnitQuery = `DELETE FROM missions.units WHERE id = $1`

const selectMissionUnitByMissionIdQuery = `SELECT * FROM missions.vw_units_mission WHERE mission_id = $1`

func (u *MissionUnitRepository) Get(id int64) *results.ResultWithValue[*models.MissionUnit] {
	r := u.AbstractRepository.Get(id, selectMissionUnitQuery)

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)

}
func (u *MissionUnitRepository) GetAll(params ...string) ([]models.MissionUnit, *results.GeneralError) {

	var MissionUnits []models.MissionUnit = make([]models.MissionUnit, 0)

	values, err := u.AbstractRepository.GetAll(selectAllMissionUnitQuery, params...)

	if err != nil {
		return MissionUnits, err
	}

	return values, nil
}

func (u *MissionUnitRepository) Create(MissionUnit *models.MissionUnit) *results.ResultWithValue[*models.MissionUnit] {

	r := u.AbstractRepository.Create(*MissionUnit, insertMissionUnitQuery, MissionUnit.GetNameArgs(), MissionUnit.SetId)

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *MissionUnitRepository) Update(MissionUnit *models.MissionUnit) *results.ResultWithValue[*models.MissionUnit] {
	r := u.AbstractRepository.Update(*MissionUnit, updateMissionUnitQuery, MissionUnit.GetNameArgs())

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *MissionUnitRepository) Delete(id int64) *results.Result {

	return u.AbstractRepository.Delete(id, deleteMissionUnitQuery)
}

func (u *MissionUnitRepository) GetByMissionId(id int) ([]models.MissionUnitSummary, error) {

	ctx := context.Background()
	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return nil, err
	}

	defer conn.Release()

	rows, err := conn.Query(ctx, selectMissionUnitByMissionIdQuery, id)

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
