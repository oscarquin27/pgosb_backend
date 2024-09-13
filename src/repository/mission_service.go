package repository

import (
	"context"
	"fdms/src/mikro"
	"fdms/src/models"
	"fdms/src/services"
	"fdms/src/utils/results"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MissionServiceRepository struct {
	db *pgxpool.Pool
}

func NewMissionServiceService(db *pgxpool.Pool) services.MissionServiceService {
	return &MissionServiceRepository{
		db: db,
	}
}

func (u *MissionServiceRepository) Get(id int) (*models.MissionService, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return nil, err
	}

	defer conn.Release()

	rows, err := conn.Query(ctx, `SELECT id, mission_id, 
	antares_id, 
	units, 
	bombers,
	operative_areas,
	summary, 
	description,
	unharmed,
	injured,
	transported,
	deceased,
	station_id,
	center_id,
	location_id,
	service_date,
	manual_service_date,
	is_important,
	sending_user_id,
	receiving_user_id

	
	FROM missions.services where id = $1;`, id)

	if err != nil {
		return nil, err
	}

	service, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.MissionService])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorMissionServiceNotFound
		}

		return nil, err
	}

	return &service, nil
}

func (u *MissionServiceRepository) GetAll() ([]models.MissionService, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return nil, err
	}

	defer conn.Release()

	rows, err := conn.Query(ctx, `SELECT id, mission_id, 
	antares_id, 
	units, 
	bombers, 
	operative_areas,
	summary,
	description,
	unharmed,
	injured ,
	transported ,
	deceased ,
	station_id,
	center_id,
	location_id,
    service_date,
	manual_service_date,
	is_important,
	sending_user_id,
	receiving_user_id
	FROM missions.services order by id desc;`)

	if err != nil {
		return nil, err
	}

	services, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.MissionService])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorMissionServiceNotFound
		}

		return nil, err
	}

	return services, nil
}

func (u *MissionServiceRepository) GetAllMissionServiceSummary() ([]models.MissionServiceSummary, error) {

	defaultValue := make([]models.MissionServiceSummary, 0)

	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return defaultValue, err
	}

	defer conn.Release()

	rows, err := conn.Query(ctx, `SELECT * FROM missions.vw_service_summary`)

	if err != nil {
		return defaultValue, err
	}

	services, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.MissionServiceSummary])

	if err != nil {
		if err == pgx.ErrNoRows {
			return defaultValue, models.ErrorMissionNotFound
		}

		return defaultValue, err
	}

	return services, nil
}

func (u *MissionServiceRepository) GetRelevantServices(id string) ([]models.RelevantServices, error) {

	defaultValue := make([]models.RelevantServices, 0)

	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return defaultValue, err
	}

	defer conn.Release()

	rows, err := conn.Query(ctx, fmt.Sprintf(`SELECT id, 
	region_area, 
	mission_code, 
	antares_id, 
	antares_type, 
	antares_description, 
	service_id, 
	operative_area_name, 
	service_description, 
	service_date::varchar, 
	units, 
	firefighters, 
	people, 
	infrastructures, 
	vehicles, 
	service_locations, 
	service_stations, 
	centers 
	FROM missions.vw_relevant_services
	where service_id::text in (%s)`, id))

	if err != nil {
		return defaultValue, err
	}

	services, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.RelevantServices])

	if err != nil {
		if err == pgx.ErrNoRows {
			return defaultValue, models.ErrorMissionNotFound
		}

		return defaultValue, err
	}

	return services, nil
}

func (u *MissionServiceRepository) GetByMissionId(id int) ([]models.MissionService, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return nil, err
	}

	defer conn.Release()

	rows, err := conn.Query(ctx, `SELECT id, mission_id, 
	antares_id, 
	units, 
	bombers,
	operative_areas, 
	summary, 
	description,
	unharmed,
	injured ,
	transported ,
	deceased ,
	station_id,
	center_id,
	location_id,
	service_date,
	manual_service_date,
	is_important,
	sending_user_id,
	receiving_user_id
	FROM missions.services where mission_id = $1 `, id)

	if err != nil {
		return nil, err
	}

	services, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.MissionService])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorMissionServiceNotFound
		}

		return nil, err
	}

	return services, nil
}

// GetUnits implements services.MissionServiceService.
func (u *MissionServiceRepository) GetUnits(id int) *results.ResultWithValue[[]models.UnitSimple] {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	r := results.NewResultWithValue[[]models.UnitSimple]("Get-Units-Simple", false, make([]models.UnitSimple, 0), nil).
		Failure()

	if err != nil {
		return r.WithError(results.NewError(err.Error(), err))
	}

	defer conn.Release()

	rows, err := conn.Query(ctx, `SELECT coalesce(id::varchar, '') as id, 
		coalesce(plate::varchar, '') as plate, 
		coalesce(station::varchar, '') as station, 
		coalesce(unit_type::varchar, '') as unit_type, 
		coalesce(alias::varchar, '') as alias
		FROM  vehicles.unit u
		where id in (select 
			unnest(s.units) as id
		from missions.services s
		where s.id = $1)`, id)

	if err != nil {
		return r.WithError(results.NewError(err.Error(), err))
	}

	services, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.UnitSimple])

	if err != nil {
		if err == pgx.ErrNoRows || len(services) == 0 {
			return r.WithError(results.NewError(err.Error(), err))
		}

		return r.WithError(results.NewError(err.Error(), err))
	}

	return r.Success().WithValue(services)
}

// GetUsers implements services.MissionServiceService.
func (u *MissionServiceRepository) GetUsers(id int) *results.ResultWithValue[[]models.MissionUserService] {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	r := results.NewResultWithValue[[]models.MissionUserService]("Get-UserMission-Simple", false, make([]models.MissionUserService, 0), nil).
		Failure()

	if err != nil {
		return r.WithError(results.NewError(err.Error(), err))
	}

	defer conn.Release()

	query := "SELECT * FROM missions.vw_services_firefighter WHERE service_id = $1"

	rows, err := conn.Query(ctx, query, id)

	if err != nil {
		return r.WithError(results.NewError(err.Error(), err))
	}

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.MissionUserService])

	if err != nil {
		if err == pgx.ErrNoRows || len(users) == 0 {
			return r.Success()
		}

		return r.WithError(results.NewError(err.Error(), err))
	}

	return r.Success().WithValue(users)
}

func (u *MissionServiceRepository) Create(s *models.MissionService) (*models.MissionService, error) {
	m := mikro.NewMkModel(u.db)

	err := m.Model(s).Omit("id").Omit("service_date").
		Returning().InsertReturning("missions.services")

	if err != nil {
		return nil, err
	}

	if s.Id.Int32 >= 0 {
		return s, nil
	}

	return nil, models.ErrorMissionServiceNotCreated
}

func (u *MissionServiceRepository) Update(s *models.MissionService) error {
	m := mikro.NewMkModel(u.db)

	rows, err := m.Model(s).Omit("id").Omit("service_date").Where("id", "=", s.Id).Update("missions.services")

	if err != nil {
		return err
	}

	if rows == 1 {
		return nil
	}

	return models.ErrorMissionServiceNotUpdated
}

func (u *MissionServiceRepository) Delete(id int) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return err
	}

	_, err = conn.Exec(ctx, "delete from missions.services where id = $1", id)

	if err != nil {
		return models.ErrorMissionServiceNotDeleted
	}

	return nil
}
