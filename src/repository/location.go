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
