package models

import (
	"errors"

	"github.com/jackc/pgx/v5/pgtype"
)

var (
	ErrorLayoutFound = errors.New("layout no encontrado")
)

type Layout struct {
	Id           pgtype.Int8 `json:"id" db:"id"`
	Column_name  pgtype.Text `json:"column_name"`
	Display_name pgtype.Text `json:"display_name"`
	Group_name   pgtype.Text `json:"group_name"`
	Entity_name  pgtype.Text `json:"entity_name"`
	Visibility   pgtype.Bool `json:"visibility"`
	Type         pgtype.Text `json:"type"`
}