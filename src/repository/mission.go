package repository

import (
	"context"
	"database/sql"
	logger "fdms/src/infrastructure/log"
	"fdms/src/models"
	"fdms/src/services"
	"fdms/src/utils/date_utils"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	insertMission = `
        INSERT INTO missions.mission (
		id, 
		code, 
		created_at, 
		alias, 
		operative_areas, 
		summary, 
		description, 
		unharmed, 
		injured, 
		transported, 
		deceased, 
		station_id, 
		location_id, 
		health_care_center_id, 
		sending_user_id, 
		receiving_user_id, 
		level, 
		peace_quadrant, 
		location_destiny_id, 
		is_important, 
		not_attended, 
		false_alarm, 
		pending_for_data
		)
        VALUES (
		@id, 
		@code, 
		@created_at, 
		@alias, 
		@operative_areas, 
		@summary, 
		@description, 
		@unharmed, 
		@injured, 
		@transported, 
		@deceased, 
		@station_id, 
		@location_id, 
		@health_care_center_id, 
		@sending_user_id, 
		@receiving_user_id, 
		@level, 
		@peace_quadrant, 
		@location_destiny_id, 
		@is_important, 
		@not_attended, 
		@false_alarm, 
		@pending_for_data
		)
    `
)

type MissionRepository struct {
	db *pgxpool.Pool
}

func NewMissionService(db *pgxpool.Pool) services.MissionService {
	return &MissionRepository{
		db: db,
	}
}

// GetAll implements services.MissionService.
func (u *MissionRepository) GetAll() ([]models.Mission, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return nil, err
	}

	defer conn.Release()

	rows, err := conn.Query(ctx, `SELECT * from missions.mission order by created_at desc`)

	if err != nil {
		return nil, err
	}

	services, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Mission])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorMissionNotFound
		}

		return nil, err
	}

	return services, nil
}

func (u *MissionRepository) GetAllMissionSummary() ([]models.MissionSummary, error) {

	defaultValue := make([]models.MissionSummary, 0)

	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return defaultValue, err
	}

	defer conn.Release()

	rows, err := conn.Query(ctx, `SELECT * FROM missions.vw_mission_summary ORDER BY id DESC`)

	if err != nil {
		return defaultValue, err
	}

	services, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.MissionSummary])

	if err != nil {
		if err == pgx.ErrNoRows {
			return defaultValue, models.ErrorMissionNotFound
		}

		return defaultValue, err
	}

	return services, nil
}

func (u *MissionRepository) Get(id int) (*models.Mission, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		logger.Error().Err(err).Msg("Error generando conexion")
		return nil, err
	}

	defer conn.Release()

	rows, err := conn.Query(ctx, `SELECT * from missions.mission where id = $1;`, id)

	if err != nil {
		logger.Error().Err(err).Msg("Error ejecutando0 querys")

		return nil, err
	}

	services, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Mission])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorMissionNotFound
		}

		return nil, err
	}

	return &services, nil
}

func (u *MissionRepository) Create(s *models.Mission) (*models.Mission, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {

		logger.Error().Err(err).Msg("Error acquiring database connection")

		return nil, err
	}

	defer conn.Release()

	// 1. Get Next ID from Sequence
	var id int64 // Use int64 for sequence value

	err = conn.QueryRow(ctx, `SELECT nextval('missions.mission_id_seq'::regclass)`).Scan(&id)

	if err != nil {
		logger.Error().Err(err).Msg("Error getting next sequence value")

		return nil, err // Return a more specific error
	}

	date := time.Now().Format(date_utils.CompleteFormatDate)

	code := fmt.Sprintf("%d-%s", id, date)

	s.Id = id

	createdAt, err := time.Parse(date_utils.CompleteFormatDate, date)

	if err != nil {

		logger.Error().Err(err).Msg("Error parsing date")

		return nil, err
	}

	s.CreatedAt = sql.NullTime{Time: createdAt, Valid: true}

	s.Code = sql.NullString{String: code, Valid: true}

	// 2. Insert with the Retrieved ID
	_, err = conn.Exec(ctx, insertMission,
		s.GetNameArgs(),
	)

	if err != nil {

		logger.Error().Err(err).Msg("Error executing insert query")

		return nil, models.ErrorMissionNotCreated
	}

	// 3. Set the ID in the Model (if needed)

	return s, nil // Return the mission with the set ID
}

func (u *MissionRepository) Update(s *models.Mission) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return err
	}

	_, err = conn.Exec(ctx, `UPDATE missions.mission
	SET code = $1,
	    alias = $2
	WHERE id = $3`, s.Code, s.Alias, s.Id)

	if err != nil {
		return err
	}

	return nil
}

func (u *MissionRepository) Delete(id int) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return err
	}

	_, err = conn.Exec(ctx, "delete from missions.mission where id = $1", id)

	if err != nil {
		return models.ErrorMissionServiceNotDeleted
	}

	return nil
}
