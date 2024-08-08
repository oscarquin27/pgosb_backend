package repository

import (
	"context"
	"fdms/src/mikro"
	"fdms/src/models"
	"fdms/src/services"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CentersRepository struct {
	db *pgxpool.Pool
}

func NewCenterService(connPool *pgxpool.Pool) services.CentersService {
	return &CentersRepository{
		db: connPool,
	}
}

// Create implements services.CentersService.
func (c *CentersRepository) Create(center *models.Center) error {
	m := mikro.NewMkModel(c.db)

	rows, err := m.Model(center).Insert("hq.centers")

	if err != nil {
		return err
	}

	if rows > 0 {
		return nil
	}

	return models.ErrorCenterNotCreated
}

// Delete implements services.CentersService.
func (c *CentersRepository) Delete(id int) error {
	ctx := context.Background()

	conn, err := c.db.Acquire(ctx)

	if err != nil {
		return err
	}

	defer conn.Release()

	rows, err := conn.Exec(ctx, "delete from hq.centers where id = $1", id)

	if err != nil {
		return err
	}

	if rows.RowsAffected() == 1 {
		return nil
	}

	return models.ErrorCenterNotDeleted
}

// Get implements services.CentersService.
func (c *CentersRepository) Get(id int) (*models.Center, error) {
	ctx := context.Background()

	conn, err := c.db.Acquire(ctx)

	if err != nil {
		return nil, err
	}

	defer conn.Release()

	rows, err := conn.Query(ctx, "select id, urb_id, type, name, address, phone  from hq.centers where id = $1", id)

	if err != nil {
		return nil, err
	}

	r, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Center])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorCenterFound
		}

		return nil, err
	}

	return &r, nil
}

// GetAll implements services.CentersService.
func (c *CentersRepository) GetAll() ([]models.Center, error) {
	ctx := context.Background()

	conn, err := c.db.Acquire(ctx)

	if err != nil {
		return nil, err
	}

	defer conn.Release()

	rows, err := conn.Query(ctx, "select id, urb_id, type, name, address, phone  from hq.centers")

	if err != nil {
		return nil, err
	}

	r, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Center])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorCenterFound
		}

		return nil, err
	}

	return r, nil
}

// Update implements services.CentersService.
func (c *CentersRepository) Update(center *models.Center) error {
	m := mikro.NewMkModel(c.db)

	rows, err := m.Model(center).Omit("id").Update("hq.centers")

	if err != nil {
		return err
	}

	if rows > 0 {
		return nil
	}

	return models.ErrorCenterNotUpdated
}


