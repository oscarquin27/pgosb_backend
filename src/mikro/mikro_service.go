package mikro

import "github.com/jackc/pgx/v5/pgxpool"

type MkModel struct {
	db *pgxpool.Pool
	model *interface{}
	params map[string]interface{}
	conditionField string
	conditionOperator string
	conditionValue any
}
