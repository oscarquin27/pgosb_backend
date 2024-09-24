package repository

import (
	"context"
	"fdms/src/models"
	"fdms/src/services"
	"fdms/src/utils/results"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MissionAuthorityServiceRepository struct {
	*AbstractRepository[models.MissionAuthorityService]
}

func NewMissionAuthorityServiceService(db *pgxpool.Pool) services.MissionAuthorityRelateService {

	abstractImplent := NewAbstractRepository[models.MissionAuthorityService](db)

	return &MissionAuthorityServiceRepository{
		&abstractImplent,
	}
}

const selectMissionAuthorityServiceQuery = "SELECT * FROM missions.authorities_service WHERE id = $1"

//const selectMissionAuthorityServiceQuerybyAuthorityId = "SELECT * FROM missions.authorities_services WHERE authority_id = $1"

const selectMissionAuthorityServiceQuerybyServiceId = `SELECT auth_serv.id as main_id, mas.*
    FROM missions.authorities_service auth_serv
    LEFT JOIN missions.vw_mission_authority_summary mas ON auth_serv.authority_id = mas.id
    WHERE auth_serv.service_id = $1
	`

const selectAllMissionAuthorityServiceQuery = "SELECT * FROM missions.authorities_service"

const insertMissionAuthorityServiceQuery = `INSERT INTO missions.authorities_service ( 

    authority_id,
    service_id,
	mission_id
    )

VALUES (
   @authority_id,
   @service_id,
   @mission_id
) RETURNING id`

const updateMissionAuthorityServiceQuery = `UPDATE missions.authorities_service
SET 
	authority_id = @authority_id,
	service_id = @service_id,
	mission_id = @mission_id,

WHERE id = @id; `

const deleteMissionAuthorityServiceQuery = `DELETE FROM missions.authorities_service WHERE id = $1`

func (u *MissionAuthorityServiceRepository) Get(id int64) *results.ResultWithValue[*models.MissionAuthorityService] {
	r := u.AbstractRepository.Get(id, selectMissionAuthorityServiceQuery)

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)

}
func (u *MissionAuthorityServiceRepository) GetAll(params ...string) ([]models.MissionAuthorityService, *results.GeneralError) {

	var states []models.MissionAuthorityService = make([]models.MissionAuthorityService, 0)

	values, err := u.AbstractRepository.GetAll(selectAllMissionAuthorityServiceQuery, params...)

	if err != nil {
		return states, err
	}

	return values, nil
}

func (u *MissionAuthorityServiceRepository) Create(state *models.MissionAuthorityService) *results.ResultWithValue[*models.MissionAuthorityService] {

	r := u.AbstractRepository.Create(*state, insertMissionAuthorityServiceQuery, state.GetNameArgs(), state.SetId)

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *MissionAuthorityServiceRepository) Update(state *models.MissionAuthorityService) *results.ResultWithValue[*models.MissionAuthorityService] {
	r := u.AbstractRepository.Update(*state, updateMissionAuthorityServiceQuery, state.GetNameArgs())

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *MissionAuthorityServiceRepository) Delete(id int64) *results.Result {

	return u.AbstractRepository.Delete(id, deleteMissionAuthorityServiceQuery)
}

func (u *MissionAuthorityServiceRepository) GetByServiceId(id int64) *results.ResultWithValue[[]models.MissionAuthorityServiceSummary] {

	defaultList := make([]models.MissionAuthorityServiceSummary, 0)

	rest := results.NewResultWithValue("GetByMissionId", false, defaultList, nil)

	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return rest.WithError(
			results.NewError("no se pudo adquirir conexion", err))
	}

	defer conn.Release()

	rows, err := conn.Query(ctx, selectMissionAuthorityServiceQuerybyServiceId, id)

	if err != nil {
		return rest.WithError(
			results.NewError("no se pudo ejecutar el query", err))
	}

	defer rows.Close()

	registers, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.MissionAuthorityServiceSummary])

	if err != nil {
		if err == pgx.ErrNoRows {
			return rest.WithError(results.NewNotFoundError("no encontraron registros", err))
		}

		return rest.WithError(
			results.NewError("no se pudo ejecutar el query", err))
	}

	return rest.WithValue(registers).Success()
}
