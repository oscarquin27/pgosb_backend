package repository

import (
	"context"
	"fdms/src/mikro"
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

	rows, err := conn.Query(ctx, "select state_id, name, coordinates from locations.states where state_id = $1", id)

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

	rows, err := conn.Query(ctx, "select state_id, name, coordinates from locations.states")

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
	m := mikro.NewMkModel(u.db)

	rows, err := m.Model(r).Omit("state_id").Insert("locations.states")

	if err != nil {
		return err
	}

	if rows > 0 {
		return nil
	}

	return models.ErrorStateNotCreated
}

func (u *LocationRepository) UpdateState(r *models.State) error {
	m := mikro.NewMkModel(u.db)

	rows, err := m.Model(r).Omit("state_id").Where("state_id", "=", r.Id).Update("locations.states")

	if err != nil {
		return err
	}

	if rows > 0 {
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

	rows, err := conn.Exec(ctx, "delete from locations.states where state_id = $1", id)

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

	rows, err := conn.Query(ctx, "select city_id, state_id, name, area_code, zip_code, coordinates from locations.cities where city_id = $1", id)

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

	rows, err := conn.Query(ctx, "select city_id, state_id, name, area_code, zip_code, coordinates from locations.cities")

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

	rows, err := conn.Exec(ctx, "insert into locations.cities (state_id, name, area_code, zip_code, coordinates) values ($1, $2, $3, $4)", r.State_Id, r.Name, r.Area_Code, r.Zip_Code, r.Coordinates)

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

	rows, err := conn.Exec(ctx, "update locations.cities set name = $1, area_code = $2, zip_code = $3, coordinates = $4 where city_id = $5", r.Name, r.Area_Code, r.Zip_Code, r.Coordinates, r.Id)

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

	rows, err := conn.Exec(ctx, "delete from locations.cities where city_id = $1", id)

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

	rows, err := conn.Query(ctx, "select municipality_id, state_id, name, coordinates from locations.municipalities where municipality_id = $1", id)

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

	rows, err := conn.Query(ctx, "select municipality_id, state_id, name, coordinates from locations.municipalities")
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

	rows, err := conn.Exec(ctx, "insert into locations.municipalities (state_id, name, coordinates) values ($1, $2, $3, $4)", r.State_Id, r.Name, r.Coordinates)

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

	rows, err := conn.Exec(ctx, "update locations.mnunicipalities set state_id = $1, name = $2, coordinates = $3 where municipality_id = $4", r.Name, r.State_Id, r.Coordinates, r.Id)

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

	rows, err := conn.Exec(ctx, "delete from locations.municipalities where municipality_id = $1", id)

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

	rows, err := conn.Query(ctx, "select parish_id, state_id, municipality_id, name, coordinates from locations.parish where parish_id = $1", id)

	if err != nil {
		return nil, err
	}

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

	rows, err := conn.Query(ctx, "select parish_id, state_id, municipality_id, name, coordinates from locations.parish")

	if err != nil {
		return nil, err
	}

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

	rows, err := conn.Exec(ctx, "insert into locations.parish (state_id, municipality_id, name, coordinates) values ($1, $2, $3, $4)", r.State_Id, r.Municipality_Id, r.Name)

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

	rows, err := conn.Exec(ctx, "update locations.parish set state_id = $1, municipality_id = $2, name = $3, coordinates = $4 where parish_id = $4", r.State_Id, r.Municipality_Id, r.Name, r.Coordinates, r.Id)

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

	rows, err := conn.Exec(ctx, "delete from locations.parish where parish_id = $1", id)

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

	rows, err := conn.Query(ctx, "select station_id, state_id, name, coordinates, description, code, abbreviation, phones, municipality_id, parish_id, sector, community, street, address from locations.fire_stations where station_id = $1", id)

	if err != nil {
		return nil, err
	}

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

	rows, err := conn.Query(ctx, "select station_id, state_id, name, coordinates, description, code, abbreviation, phones, municipality_id, parish_id, sector, community, street, address from locations.fire_stations")

	if err != nil {
		return nil, err
	}

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
	(state_id, name, coordinates, description, code, abbreviation, phones, municipality_id, parish_id, sector, community, street, address) 
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

	rows, err := conn.Exec(ctx, "update locations.fire_stations set state_id = $1, name = $2, coordinates = $3, description = $4, code = $5, abbreviation = $6, phones = $7, state_id = $8, parish_id = $9, sector = $10, community = $11, street = $12, address = $13 where station_id = $14", r.Municipality_id, r.Name, r.Coordinates, r.Description, r.Code, r.Abbreviation, r.Phones, r.State_id, r.Parish_id, r.Sector, r.Community, r.Street, r.Address, r.Id)

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

	rows, err := conn.Exec(ctx, "delete from locations.fire_stations where station_id = $1", id)

	if err != nil {
		return err
	}

	if rows.RowsAffected() > 0 {
		return nil
	}

	return models.ErrorStationNotDeleted
}
