package repository

import (
	"context"
	"fdms/src/models"
	"fdms/src/services"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AntaresRepository struct {
	db *pgxpool.Pool
}

func NewAntaresService(db *pgxpool.Pool) services.AntaresService {
	return &AntaresRepository{
		db: db,
	}
}

func (u *AntaresRepository) GetAll() ([]models.Antares, error) {

	ctx := context.Background()

	conn, err := u.db.Acquire(ctx)

	if err != nil {
		return nil, err
	}

	defer conn.Release()

	rows, err := conn.Query(ctx, "select id, type, description from missions.antares ")

	if err != nil {
		return nil, err
	}

	r, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Antares])

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorStateFound
		}

		return nil, err
	}

	return r, nil
}
