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

const selectAllMissionFirefighterQuery = "SELECT * FROM missions.firefighters"

const insertMissionFirefighterQuery = `INSERT INTO missions.firefighters(
	 user_id, service_role,mission_id)
	VALUES ( @user_id, @service_role,@mission_id) RETURNING id`

const updateMissionFirefighterQuery = `UPDATE missions.firefighters
	SET   user_id=@user_id, service_role=@service_role, mission_id=@mission_id
	WHERE id = @id `

const deleteMissionFirefighterQuery = `DELETE FROM missions.firefighters WHERE id = $1`

const selectMissionFirefighterByMissionIdQuery = `SELECT * FROM missions.vw_firefighters_mission WHERE mission_id = $1`

func (u *MissionFirefighterRepository) Get(id int64) *results.ResultWithValue[*models.MissionFirefighter] {
	r := u.AbstractRepository.Get(id, selectMissionFirefighterQuery)

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)

}
func (u *MissionFirefighterRepository) GetAll(params ...string) ([]models.MissionFirefighter, *results.GeneralError) {

	var MissionFirefighters []models.MissionFirefighter = make([]models.MissionFirefighter, 0)

	values, err := u.AbstractRepository.GetAll(selectAllMissionFirefighterQuery, params...)

	if err != nil {
		return MissionFirefighters, err
	}

	return values, nil
}

func (u *MissionFirefighterRepository) Create(MissionFirefighter *models.MissionFirefighter) *results.ResultWithValue[*models.MissionFirefighter] {

	r := u.AbstractRepository.Create(*MissionFirefighter, insertMissionFirefighterQuery, MissionFirefighter.GetNameArgs(), MissionFirefighter.SetId)

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *MissionFirefighterRepository) Update(MissionFirefighter *models.MissionFirefighter) *results.ResultWithValue[*models.MissionFirefighter] {
	r := u.AbstractRepository.Update(*MissionFirefighter, updateMissionFirefighterQuery, MissionFirefighter.GetNameArgs())

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *MissionFirefighterRepository) Delete(id int64) *results.Result {

	return u.AbstractRepository.Delete(id, deleteMissionFirefighterQuery)
}

func (u *MissionFirefighterRepository) GetByMissionId(id int) ([]models.MissionFirefighterUser, error) {

	ctx := context.Background()
	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return nil, err
	}

	defer conn.Release()

	rows, err := conn.Query(ctx, selectMissionFirefighterByMissionIdQuery, id)

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
