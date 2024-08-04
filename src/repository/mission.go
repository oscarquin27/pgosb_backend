package repository

import (
	"context"
	logger "fdms/src/infrastructure/log"
	"fdms/src/models"
	"fdms/src/services"
	"fdms/src/utils"
	"fdms/src/utils/date_utils"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
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

	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, `SELECT id, created_at, code from missions.mission order by created_at desc`)

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

func (u *MissionRepository) Get(id int) (*models.Mission, error) {
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		logger.Error().Err(err).Msg("Error generando conexion")
		return nil, err
	}

	defer conn.Release()

	rows, err := conn.Query(ctx, `SELECT id, created_at, code from missions.mission where id = $1;`, id)

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

	// 2. Insert with the Retrieved ID
	_, err = conn.Exec(ctx, `
        INSERT INTO missions.mission (id, code, created_at)
        VALUES ($1, $2, $3)`,
		id, code, date,
	)
	if err != nil {
		logger.Error().Err(err).Msg("Error executing insert query")
		return nil, models.ErrorMissionNotCreated
	}

	// 3. Set the ID in the Model (if needed)
	s.Id = utils.ConvertToPgTypeInt4(int(id)) // Assuming you need this conversion
	s.CreatedAt = utils.ConvertToPgTypeDate(date)
	s.Code = utils.ConvertToPgTypeText(code)

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
	SET code = $1
	WHERE id = $2`, s.Code, s.Id)

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
