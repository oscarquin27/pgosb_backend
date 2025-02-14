package repository

import (
	"context"
	"fdms/src/models"
	"fdms/src/services"
	"fdms/src/utils/results"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MissionFirefighterRepository struct {
	*AbstractRepository[models.MissionFirefighter]
}

func NewMissionFirefighterService(db *pgxpool.Pool) services.MissionFirefighterService {

	abstractImplent := NewAbstractRepository[models.MissionFirefighter](db)

	return &MissionFirefighterRepository{
		&abstractImplent,
	}
}

const selectMissionFirefighterQuery = "SELECT * FROM missions.firefighters WHERE id = $1"
const selectMissionFirefighterQueryTemplate = "SELECT * FROM missions.firefighters_template WHERE id = $1"

const selectAllMissionFirefighterQuery = "SELECT * FROM missions.firefighters"
const selectAllMissionFirefighterTemplateQuery = "SELECT * FROM missions.firefighters_template"

const insertMissionFirefighterQuery = `INSERT INTO missions.firefighters(
	 user_id, service_role,mission_id)
	VALUES ( @user_id, @service_role,@mission_id) RETURNING id`
const insertMissionFirefighterQueryTemplate = `INSERT INTO missions.firefighters_template(
		user_id, service_role,mission_id)
	   VALUES ( @user_id, @service_role,@mission_id) RETURNING id`

const updateMissionFirefighterQuery = `UPDATE missions.firefighters
	SET   user_id=@user_id, service_role=@service_role, mission_id=@mission_id
	WHERE id = @id `
const updateMissionFirefighterQueryTemplate = `UPDATE missions.firefighters_template
	SET   user_id=@user_id, service_role=@service_role, mission_id=@mission_id
	WHERE id = @id `

const deleteMissionFirefighterQuery = `DELETE FROM missions.firefighters WHERE id = $1`
const deleteMissionFirefighterQueryTemplate = `DELETE FROM missions.firefighters_template WHERE id = $1`

const selectMissionFirefighterByMissionIdQuery = `SELECT * FROM missions.vw_firefighters_mission WHERE mission_id = $1`
const selectMissionFirefighterByMissionIdQueryTemplate = `SELECT * FROM missions.vw_firefighters_mission_template WHERE mission_id = $1`

func (u *MissionFirefighterRepository) Get(id int64, isTemplate bool) *results.ResultWithValue[*models.MissionFirefighter] {
	query := selectMissionFirefighterQuery

	if isTemplate {
		query = selectMissionFirefighterQueryTemplate
	}

	r := u.AbstractRepository.Get(id, query, false)

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)

}
func (u *MissionFirefighterRepository) GetAll(isTemmplate bool, params ...string) ([]models.MissionFirefighter, *results.GeneralError) {

	var MissionFirefighters []models.MissionFirefighter = make([]models.MissionFirefighter, 0)

	query := selectAllMissionFirefighterQuery

	if isTemmplate {
		query = selectAllMissionFirefighterTemplateQuery
	}

	values, err := u.AbstractRepository.GetAll(query, false, params...)

	if err != nil {
		return MissionFirefighters, err
	}

	return values, nil
}

func (u *MissionFirefighterRepository) Create(MissionFirefighter *models.MissionFirefighter, isTemplate bool) *results.ResultWithValue[*models.MissionFirefighter] {

	query := insertMissionFirefighterQuery

	if isTemplate {
		query = insertMissionFirefighterQueryTemplate
	}

	r, id := u.AbstractRepository.Create(*MissionFirefighter, query, MissionFirefighter.GetNameArgs(), MissionFirefighter.SetId, false)

	MissionFirefighter.Id = id

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *MissionFirefighterRepository) Update(MissionFirefighter *models.MissionFirefighter, isTemplate bool) *results.ResultWithValue[*models.MissionFirefighter] {

	query := updateMissionFirefighterQuery

	if isTemplate {
		query = updateMissionFirefighterQueryTemplate
	}
	r := u.AbstractRepository.Update(*MissionFirefighter, query, MissionFirefighter.GetNameArgs(), false)

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *MissionFirefighterRepository) Delete(id int64, isTemplate bool) *results.Result {

	query := deleteMissionFirefighterQuery

	if isTemplate {
		query = deleteMissionFirefighterQueryTemplate
	}

	return u.AbstractRepository.Delete(id, query, false)
}

func (u *MissionFirefighterRepository) GetByMissionId(id int, isTemplate bool) ([]models.MissionFirefighterUser, error) {

	ctx := context.Background()
	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return nil, err
	}

	defer conn.Release()

	query := selectMissionFirefighterByMissionIdQuery

	if isTemplate {
		query = selectMissionFirefighterByMissionIdQueryTemplate
	}

	rows, err := conn.Query(ctx, query, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.MissionFirefighterUser])

	if err != nil {
		return nil, err
	}

	return users, nil

}
