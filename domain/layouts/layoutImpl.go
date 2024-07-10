package layout_domain

import (
	"context"
	layout "fdms/domain/entities/layouts"

	"github.com/jackc/pgx/v5"
)

func (u *LayoutImpl) GetLayout(entity string) ([]layout.Layout, error) {
	
	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, "select id, column_name, display_name, group_name, visibility, entity_name, type from layouts.layout where entity_name = $1", entity)

	r, err := pgx.CollectRows(rows, pgx.RowToStructByName[layout.Layout])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, layout.ErrorLayoutFound
		}

		return nil, err
	}

	return r, nil
}