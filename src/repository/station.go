package repository

import (
	"context"
	"fdms/src/models"
	"fdms/src/services"
	"fdms/src/utils/results"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StationRepository struct {
	connPool *pgxpool.Pool
}

func NewStationService(connPool *pgxpool.Pool) services.StationService {
	return &StationRepository{
		connPool: connPool,
	}
}

func (u *StationRepository) Get(id int64) *results.ResultWithValue[*models.Station] {

	r := results.NewResultWithValue[*models.Station]("Get-Station", false, nil, nil).Failure()

	ctx := context.Background()

	conn, err := u.connPool.Acquire(ctx)

	if err != nil {
		return r.WithError(
			results.NewUnknowError("no se pudo adquirir conexion", err))
	}

	defer conn.Release()

	rows, err := conn.Query(ctx, `SELECT 
	id, 
	municipality_id, 
	name, 
	coordinates, 
	description, 
	code, 
	abbreviation, 
	phones, 
	state_id, 
	parish_id, 
	sector, c
	ommunity, 
	street, 
	institution,
	state, 
	municipality, 
	parish,
	address  
    FROM locations.fire_stations 
    WHERE station_id = $1`, id)

	if err != nil {
		return r.WithError(
			results.NewUnknowError("no se pudo ejecutar query", err))
	}

	defer rows.Close()

	station, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Station])

	if err != nil {
		if err == pgx.ErrNoRows {
			return r.WithError(results.NewNotFoundError("no se encontro el role espeificado", err))
		}

		return r.WithError(
			results.NewUnknowError("error obteniendo unidad", err))
	}

	return r.Success().WithValue(&station)
}

func (u *StationRepository) GetAll() ([]models.Station, *results.GeneralError) {

	var statiosDefault []models.Station = make([]models.Station, 0)

	ctx := context.Background() // Or use a specific context

	conn, err := u.connPool.Acquire(ctx)

	if err != nil {
		return statiosDefault, results.
			NewUnknowError("no se pudo adquirir conexion", err)
	}

	defer conn.Release()

	rows, err := conn.Query(ctx, `SELECT 
	id, 
	municipality_id, 
	name, 
	coordinates, 
	description, 
	code, 
	abbreviation, 
	phones, 
	regions,
	state_id, 
	parish_id, 
	sector, 
	community, 
	street, 
	institution, 
	state, 
	municipality, 
	address,
	parish FROM locations.fire_stations`)

	if err != nil {
		return statiosDefault, results.
			NewUnknowError("no se pudo ejecutar el query", err)
	}

	defer rows.Close()

	stationsValues, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Station])

	if err != nil {
		if err == pgx.ErrNoRows {
			return statiosDefault, results.NewNotFoundError("no encontraron registros", err)
		}

		return statiosDefault, results.
			NewUnknowError("no se pudo ejecutar el query", err)
	}

	return stationsValues, nil
}

func (u *StationRepository) Create(station *models.Station) *results.ResultWithValue[*models.Station] {

	r := results.NewResultWithValue[*models.Station]("Create-Station", false, nil, nil).Failure()

	ctx := context.Background()

	conn, err := u.connPool.Acquire(ctx)

	if err != nil {
		return r.WithError(
			results.NewUnknowError("no se pudo adquirir conexion", err))
	}

	defer conn.Release()

	row := conn.QueryRow(ctx, `
	INSERT INTO locations.fire_stations 
	(municipality_id, name, coordinates, description, code, abbreviation, phones, state_id, parish_id, sector, community, street, institution, state, municipality, parish, regions, address) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18) RETURNING id`,
		station.Municipality_id, station.Name,
		station.Coordinates,
		station.Description, station.Code, station.Abbreviation, station.Phones, station.State_id, station.Parish_id, station.Sector,
		station.Community, station.Street, station.Institution, station.State, station.Municipality, station.Parish,
		station.Regions, station.Address,
	)

	err = row.Scan(&station.Id)

	if err != nil {
		return r.WithError(
			results.NewUnknowError("no se pudo ejecutar query", err))
	}

	return r.Success().WithValue(station)
}

func (u *StationRepository) Update(station *models.Station) *results.ResultWithValue[*models.Station] {

	r := results.NewResultWithValue[*models.Station]("Update-Stattion", false, nil, nil).Failure()

	ctx := context.Background()

	conn, err := u.connPool.Acquire(ctx)

	if err != nil {
		return r.WithError(
			results.NewUnknowError("no se pudo adquirir conexion", err))
	}

	defer conn.Release()

	rows, err := conn.Exec(ctx, `
	UPDATE locations.fire_stations 
	SET municipality_id = $1, name = $2, coordinates = $3, description = $4, code = $5, abbreviation = $6, phones = $7, state_id = $8, parish_id = $9, 
	sector = $10, 
	community = $11, street = $12, institution = $13, 
	state = $14, municipality = $15, parish = $16, regions = $17 , address = $18
	WHERE id = $19`,
		station.Municipality_id, station.Name, station.Coordinates, station.Description, station.Code, station.Abbreviation,
		station.Phones, station.State_id, station.Parish_id,
		station.Sector, station.Community,
		station.Street, station.Institution, station.State, station.Municipality, station.Parish, station.Regions, station.Address, station.Id,
	)

	if err != nil {
		return r.WithError(
			results.NewUnknowError("no se pudo ejecutar query", err))
	}

	if rows.RowsAffected() == 0 {
		return r.WithError(
			results.NewNotFoundError("no se consiguio registro", err))
	}

	if rows.RowsAffected() > 1 {
		return r.WithError(
			results.NewUnknowError("se afecto mas de un registro", err))
	}

	return r.Success().WithValue(station)
}

func (u *StationRepository) Delete(id int64) *results.Result {

	r := results.NewResult("Delete-Station", false, nil).Failure()

	ctx := context.Background()

	conn, err := u.connPool.Acquire(ctx)

	if err != nil {
		return r.FailureWithError(err)
	}

	defer conn.Release()

	rows, err := conn.Exec(ctx, "delete from locations.fire_stations where id = $1", id)

	if err != nil {
		return r.FailureWithError(err)
	}

	if rows.RowsAffected() == 0 {
		return r.WithError(results.NewNotFoundError("no se consiguio registro", err))
	}

	if rows.RowsAffected() > 1 {
		return r.WithError(results.NewUnknowError("se borraron multiples registros", err))
	}

	return r.Success()
}
