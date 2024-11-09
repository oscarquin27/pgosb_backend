package repository

import "github.com/jackc/pgx/v5"

type ServiceRelationEntityRepository struct {
	db *pgx.Conn
}
