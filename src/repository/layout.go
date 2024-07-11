package repository

import (
	"context"
	"fdms/src/models"
	"fdms/src/services"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LayoutRepository struct {
	connPool *pgxpool.Pool
}

func NewLayoutService(connPool *pgxpool.Pool) services.LayoutService {
	return &LayoutRepository{
		connPool: connPool,
	}
}

func (u *LayoutRepository) Get(entity string) ([]models.Layout, error) {

	ctx := context.Background()

	conn, err := u.connPool.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, "select id, column_name, display_name, group_name, visibility, entity_name, type from modelss.models where entity_name = $1", entity)

	if err != nil {
		return nil, err
	}

	r, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Layout])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorLayoutFound
		}

		return nil, err
	}

	return r, nil
}
