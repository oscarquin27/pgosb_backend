package layout_domain

import (
	layouts "fdms/domain/entities/layouts"

	"github.com/jackc/pgx/v5/pgxpool"
)
type LayoutRepository interface {
	GetLayout(entity string) ([]layouts.Layout, error)
}

type LayoutImpl struct {
	db *pgxpool.Pool
}

func NewLayoutService(db *pgxpool.Pool) LayoutRepository {
	return &LayoutImpl{
		db : db,
	}
}