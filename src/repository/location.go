package repository

import (
	"context"
	"fdms/src/models"
	"fdms/src/services"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LocationRepository struct {
	db *pgxpool.Pool
}

func NewLocationService(db *pgxpool.Pool) services.LocationsService {
	return &LocationRepository{
		db: db,
	}
}

func (u *LocationRepository) GetState(id int64) (*models.State, error) {

	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, "select models_id, name, coordinates from locations.modelss where models_id = $1", id)

	if err != nil {
		return nil, err
	}

	r, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.State])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorStateFound
		}

		return nil, err
	}

	return &r, nil
}

func (u *LocationRepository) GetAllStates() ([]models.State, error) {

	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, "select models_id, name, coordinates from locations.modelss")

	if err != nil {
		return nil, err
	}

	r, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.State])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorStateFound
		}

		return nil, err
	}

	return r, nil
}

func (u *LocationRepository) CreateState(r *models.State) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, "insert into locations.models (name, coordinates) values ($1, $2)", r.Coordinates, r.Name)

	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return models.ErrorStateNotCreated
}

func (u *LocationRepository) UpdateState(r *models.State) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, "update locations.models set name = $1, coordinates = $2 where models_id = $3", r.Name, r.Coordinates, r.Id)

	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return models.ErrorStateNotUpdated
}

func (u *LocationRepository) DeleteState(id int64) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, "delete from locations.models where models_id = $1", id)

	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return models.ErrorStateNotDeleted
}

func (u *LocationRepository) GetCity(id int64) (*models.City, error) {

	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, "select models_id, models_id, name, area_code, zip_code, coordinates from locations.cities where models_id = $1", id)

	if err != nil {
		return nil, err
	}

	r, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.City])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorCityFound
		}

		return nil, err
	}

	return &r, nil
}

func (u *LocationRepository) GetAllCity() ([]models.City, error) {

	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, "select models_id, models_id, name, area_code, zip_code, coordinates from locations.cities")

	if err != nil {
		return nil, err
	}

	r, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.City])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorCityFound
		}

		return nil, err
	}

	return r, nil
}

func (u *LocationRepository) CreateCity(r *models.City) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, "insert into locations.models (models_id, name, area_code, zip_code, coordinates) values ($1, $2, $3, $4)", r.State_Id, r.Name, r.Area_Code, r.Zip_Code, r.Coordinates)

	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return models.ErrorCityNotCreated
}

func (u *LocationRepository) UpdateCity(r *models.City) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, "update locations.cities set name = $1, area_code = $2, zip_code = $3, coordinates = $4 where models_id = $5", r.Name, r.Area_Code, r.Zip_Code, r.Coordinates, r.Id)

	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return models.ErrorCityNotUpdated
}

func (u *LocationRepository) DeleteCity(id int64) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, "delete from locations.cities where models_id = $1", id)

	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return models.ErrorCityNotDeleted
}

func (u *LocationRepository) GetMunicipality(id int64) (*models.Municipality, error) {

	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, "select models_id, models_id, name, coordinates from locations.municipalities where models_id = $1", id)

	if err != nil {
		return nil, err
	}

	r, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Municipality])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorMunicipalityFound
		}

		return nil, err
	}

	return &r, nil
}

func (u *LocationRepository) GetAllMunicipality() ([]models.Municipality, error) {

	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, "select models_id, models_id, name, coordinates from locations.municipalities")
	if err != nil {
		return nil, err
	}

	r, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Municipality])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorMunicipalityFound
		}

		return nil, err
	}

	return r, nil
}

func (u *LocationRepository) CreateMunicipality(r *models.Municipality) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, "insert into locations.municipalities (models_id, name, coordinates) values ($1, $2, $3, $4)", r.State_Id, r.Name, r.Coordinates)

	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return models.ErrorCityNotCreated
}

func (u *LocationRepository) UpdateMunicipality(r *models.Municipality) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, "update locations.mnunicipalities set models_id = $1, name = $2, coordinates = $3 where models_id = $4", r.Name, r.State_Id, r.Coordinates, r.Id)

	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return models.ErrorMunicipalityNotCreated
}

func (u *LocationRepository) DeleteMunicipality(id int64) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, "delete from locations.municipalities where models_id = $1", id)

	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return models.ErrorMunicipalityNotDeleted
}

func (u *LocationRepository) GetParish(id int64) (*models.Parish, error) {

	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, "select models_id, models_id, models_id, name, coordinates from locations.models where models_id = $1", id)

	r, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Parish])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorParishFound
		}

		return nil, err
	}

	return &r, nil
}

func (u *LocationRepository) GetAllParish() ([]models.Parish, error) {

	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, "select models_id, models_id, models_id, name, coordinates from locations.models")

	r, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Parish])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorParishFound
		}

		return nil, err
	}

	return r, nil
}

func (u *LocationRepository) CreateParish(r *models.Parish) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, "insert into locations.models (models_id, models_id, name, coordinates) values ($1, $2, $3, $4)", r.State_Id, r.Municipality_Id, r.Name)

	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return models.ErrorCityNotCreated
}

func (u *LocationRepository) UpdateParish(r *models.Parish) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, "update locations.models set models_id = $1, models_id = $2, name = $3, coordinates = $4 where models_id = $4", r.State_Id, r.Municipality_Id, r.Name, r.Coordinates, r.Id)

	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return models.ErrorParishNotUpdated
}

func (u *LocationRepository) DeleteParish(id int64) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, "delete from locations.models where models_id = $1", id)

	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return models.ErrorParishNotDeleted
}

func (u *LocationRepository) GetStation(id int64) (*models.Station, error) {

	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, "select models_id, models_id, name, coordinates, description, code, abbreviation, phones, models_id, models_id, sector, community, street, address from locations.fire_modelss where models_id = $1", id)

	r, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Station])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorStationFound
		}

		return nil, err
	}

	return &r, nil
}

func (u *LocationRepository) GetAllStations() ([]models.Station, error) {

	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, "select models_id, models_id, name, coordinates, description, code, abbreviation, phones, models_id, models_id, sector, community, street, address from locations.fire_modelss")

	r, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Station])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorStationFound
		}

		return nil, err
	}

	return r, nil
}

func (u *LocationRepository) CreateStation(r *models.Station) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, `insert into locations.fire_stations
	(models_id, name, coordinates, description, code, abbreviation, phones, models_id, models_id, sector, community, street, address) 
	values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`,
		r.Municipality_id,
		r.Name,
		r.Coordinates,
		r.Description,
		r.Code,
		r.Abbreviation,
		r.Phones,
		r.State_id,
		r.Parish_id,
		r.Sector,
		r.Community,
		r.Street,
		r.Address)

	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return models.ErrorStationFound
}

func (u *LocationRepository) UpdateStation(r *models.Station) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, "update locations.fire_stations set models_id = $1, name = $2, coordinates = $3, description = $4, code = $5, abbreviation = $6, phones = $7, models_id = $8, models_id = $9, sector = $10, community = $11, street = $12, address = $13 where models_id = $14", r.Municipality_id, r.Name, r.Coordinates, r.Description, r.Code, r.Abbreviation, r.Phones, r.State_id, r.Parish_id, r.Sector, r.Community, r.Street, r.Address, r.Id)

	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return models.ErrorStationNotUpdated
}

func (u *LocationRepository) DeleteStation(id int64) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return err
	}

	rows, err := conn.Exec(ctx, "delete from locations.fire_stations where models_id = $1", id)

	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return models.ErrorStationNotDeleted
}
