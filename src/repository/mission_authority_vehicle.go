package repository

import (
	"context"
	"fdms/src/models"
	"fdms/src/services"
	"fdms/src/utils/results"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MissionAuthorityVehicleRepository struct {
	*AbstractRepository[models.MissionAuthorityVehicle]
}

func NewMissionAuthorityVehicleService(db *pgxpool.Pool) services.MissionAuthorityVehicleService {

	abstractImplent := NewAbstractRepository[models.MissionAuthorityVehicle](db)

	return &MissionAuthorityVehicleRepository{
		&abstractImplent,
	}
}

const selectMissionAuthorityVehicleQuery = "SELECT * FROM missions.authorities_vehicle WHERE id = $1"

const selectMissionAuthorityVehicleQuerybyAuthorityId = "SELECT * FROM missions.authorities_vehicle WHERE authority_id = $1"

const selectAllMissionAuthorityVehicleQuery = "SELECT * FROM missions.authorities_vehicle"

const insertMissionAuthorityVehicleQuery = `INSERT INTO missions.authorities_vehicle ( 

    mission_id,
    authority_id,
    type,
    make,
    model,

    plate,
    year,
    color,
    description

    )
VALUES (
	@mission_id,
	@authority_id, 
	@type,
	@make,

	@model,

	@plate,
	@year,

	@color,
	@description


) RETURNING id`

const updateMissionAuthorityVehicleQuery = `UPDATE missions.authorities_vehicle
SET 
	type = @type,
	make = @make,
	model = @model,
	plate = @plate,
	year = @year,
	color = @color,
description = @description
WHERE id = @id; `

const deleteMissionAuthorityVehicleQuery = `DELETE FROM missions.authorities_vehicle WHERE id = $1`

func (u *MissionAuthorityVehicleRepository) Get(id int64) *results.ResultWithValue[*models.MissionAuthorityVehicle] {
	r := u.AbstractRepository.Get(id, selectMissionAuthorityVehicleQuery)

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)

}
func (u *MissionAuthorityVehicleRepository) GetAll(params ...string) ([]models.MissionAuthorityVehicle, *results.GeneralError) {

	var states []models.MissionAuthorityVehicle = make([]models.MissionAuthorityVehicle, 0)

	values, err := u.AbstractRepository.GetAll(selectAllMissionAuthorityVehicleQuery, params...)

	if err != nil {
		return states, err
	}

	return values, nil
}

func (u *MissionAuthorityVehicleRepository) Create(state *models.MissionAuthorityVehicle) *results.ResultWithValue[*models.MissionAuthorityVehicle] {

	r := u.AbstractRepository.Create(*state, insertMissionAuthorityVehicleQuery, state.GetNameArgs(), state.SetId)

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *MissionAuthorityVehicleRepository) Update(state *models.MissionAuthorityVehicle) *results.ResultWithValue[*models.MissionAuthorityVehicle] {
	r := u.AbstractRepository.Update(*state, updateMissionAuthorityVehicleQuery, state.GetNameArgs())

	return results.NewResultWithValue(r.StepIdentifier, r.IsSuccessful, &r.Value, r.Err)
}

func (u *MissionAuthorityVehicleRepository) Delete(id int64) *results.Result {

	return u.AbstractRepository.Delete(id, deleteMissionAuthorityVehicleQuery)
}

func (u *MissionAuthorityVehicleRepository) GetByAuthorityId(id int64) *results.ResultWithValue[[]models.MissionAuthorityVehicle] {

	defaultList := make([]models.MissionAuthorityVehicle, 0)

	rest := results.NewResultWithValue("GetByMissionId", false, defaultList, nil)

	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return rest.WithError(
			results.NewError("no se pudo adquirir conexion", err))
	}

	defer conn.Release()

	rows, err := conn.Query(ctx, selectMissionAuthorityVehicleQuerybyAuthorityId, id)

	if err != nil {
		return rest.WithError(
			results.NewError("no se pudo ejecutar el query", err))
	}

	defer rows.Close()

	registers, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.MissionAuthorityVehicle])

	if err != nil {
		if err == pgx.ErrNoRows {
			return rest.WithError(results.NewNotFoundError("no encontraron registros", err))
		}

		return rest.WithError(
			results.NewError("no se pudo ejecutar el query", err))
	}

	return rest.WithValue(registers).Success()
}
