package repository

import (
	"context"
	"database/sql"
	logger "fdms/src/infrastructure/log"
	"fdms/src/models"
	"fdms/src/services"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	selectAllMissions = `
	SELECT * from missions.mission order by created_at desc
	`
	selectAllMissionsTemplate = `
	SELECT * from missions.mission_template order by created_at desc
	`

	insertMission = `

        INSERT INTO missions.mission (
		 
		code, 
		
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
		center_id, 
		sending_user_id, 
		receiving_user_id, 
		level, 
		peace_quadrant, 
		location_destiny_id, 
		is_important, 
		pending_for_data,
		cancel_reason,
		manual_mission_date
		
		)
        VALUES (
		
		@code, 
	
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
		@center_id, 
		@sending_user_id, 
		@receiving_user_id, 
		@level, 
		@peace_quadrant, 
		@location_destiny_id, 
		@is_important, 
		@pending_for_data,
		@cancel_reason,
		@manual_mission_date 
		
		
		) 
		RETURNING id , created_at
    `

	insertMissionTemplate = `
	INSERT INTO missions.mission_template (
		name, description, status, created_at, updated_at, deleted_at
	) VALUES (
		$1, $2, $3, $4, $5, $6
	) RETURNING *
	`

	updateMission = `
	UPDATE missions.mission
	SET 
		code = @code,
		alias = @alias,
		operative_areas = @operative_areas,
		summary = @summary,
		description = @description,
		unharmed = @unharmed,
		injured = @injured,
		transported = @transported,
		deceased = @deceased,
		station_id = @station_id,
		location_id = @location_id,
		center_id = @center_id,
		sending_user_id = @sending_user_id,
		receiving_user_id = @receiving_user_id,
		level = @level,
		peace_quadrant = @peace_quadrant,
		location_destiny_id = @location_destiny_id,
		is_important = @is_important,
		pending_for_data = @pending_for_data,
		cancel_reason = @cancel_reason,
		manual_mission_date = @manual_mission_date
	
		WHERE id = @id
	`

	updateMissionTemplate = `
	UPDATE missions.mission_template 
	SET name = $1,
		description = $2,
		status = $3,
		updated_at = $4
	WHERE id = $5
	RETURNING *
	`

	deleteMission = `
	DELETE FROM missions.mission WHERE id = $1
	`

	deleteMissionTemplate = `
	DELETE FROM missions.mission_template WHERE id = $1
	`

	selectMissionSummary = `SELECT * FROM missions.vw_mission_summary ORDER BY id DESC`

	selectMissionSummaryTemplate = `SELECT * FROM missions.vw_mission_summary_template ORDER BY id DESC`

	selectMissionById = `SELECT * FROM missions.mission WHERE id = $1`

	selectMissionByIdTemplate = `SELECT * FROM missions.mission_template WHERE id = $1`
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
func (u *MissionRepository) GetAll(isTemplate bool) ([]models.Mission, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return nil, err
	}

	defer conn.Release()

	query := selectAllMissions

	if isTemplate {
		query = selectAllMissionsTemplate
	}

	rows, err := conn.Query(ctx, query)

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

func (u *MissionRepository) GetAllMissionSummary(isTemplate bool) ([]models.MissionSummary, error) {

	defaultValue := make([]models.MissionSummary, 0)

	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return defaultValue, err
	}

	defer conn.Release()

	query := selectMissionSummary

	if isTemplate {
		query = selectMissionSummaryTemplate
	}

	rows, err := conn.Query(ctx, query)

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

func (u *MissionRepository) Get(id int64, isTemplate bool) (*models.Mission, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		logger.Error().Err(err).Msg("Error generando conexion")
		return nil, err
	}

	defer conn.Release()

	query := selectMissionById

	if isTemplate {
		query = selectMissionByIdTemplate
	}

	rows, err := conn.Query(ctx, query, id)

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

func (u *MissionRepository) Create(s *models.Mission, isTemplate bool) (*models.Mission, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {

		logger.Error().Err(err).Msg("Error acquiring database connection")

		return nil, err
	}

	defer conn.Release()

	query := insertMission

	if isTemplate {
		query = insertMissionTemplate
	}

	s.Code = sql.NullString{String: "", Valid: true}

	// 2. Insert with the Retrieved ID

	var id int64
	var createdAt time.Time

	err = conn.QueryRow(ctx, query,
		s.GetNameArgs(),
	).Scan(&id, &createdAt)

	if err != nil {

		logger.Error().Err(err).Msg("Error executing insert query")

		return nil, models.ErrorMissionNotCreated
	}

	s.Id = id
	s.CreatedAt = sql.NullTime{Time: createdAt, Valid: true}

	// 3. Set the ID in the Model (if needed)

	return s, nil // Return the mission with the set ID
}

func (u *MissionRepository) Update(s *models.Mission, isTemplate bool) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return err
	}

	query := updateMission

	if isTemplate {
		query = updateMissionTemplate
	}

	_, err = conn.Exec(ctx, query, s.GetNameArgs())

	if err != nil {
		return err
	}

	return nil
}

func (u *MissionRepository) Delete(id int64, isTemplate bool) error {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return err
	}

	query := deleteMission

	if isTemplate {
		query = deleteMissionTemplate
	}

	_, err = conn.Exec(ctx, query, id)

	if err != nil {
		return models.ErrorMissionServiceNotDeleted
	}

	return nil
}
