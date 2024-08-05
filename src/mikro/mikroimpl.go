package mikro

import (
	"context"
	"reflect"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewMkModel(db *pgxpool.Pool) *MkModel {
	return &MkModel{
		db: db,
	}
}

func (mk *MkModel) Model(value interface{}) *MkModel {
	params := extractParams(value)

	return &MkModel{
		db:     mk.db,
		model:  &value,
		params: params,
	}
}

func (mk *MkModel) Insert(table string) (int64, error) {
	var fields []string
	var values []interface{}

	for k, v := range mk.params {
		fields = append(fields, k)
		values = append(values, v)
	}

	sentence := buildInsert(fields, table)

	return executeSentence(mk.db, sentence, values)
}

func (mk *MkModel) InsertReturning(table string) error {
	var fields []string
	var values []interface{}

	for k, v := range mk.params {
		fields = append(fields, k)
		values = append(values, v)
	}

	sentence := buildInsert(fields, table)

	sentence += mk.returning

	return executeSentenceReturning(mk.db, sentence, values, mk)
}

func (mk *MkModel) UpdateReturning(table string) error {
	var fields []string
	var values []interface{}

	for k, v := range mk.params {
		fields = append(fields, k)
		values = append(values, v)
	}

	if len(mk.conditionField) > 0 {
		values = append(values, mk.conditionValue)
	}

	sentence := buildUpdate(fields, table, *mk)
	sentence += mk.returning

	return executeSentenceReturning(mk.db, sentence, values, mk)
}

func (mk *MkModel) Update(table string) (int64, error) {
	var fields []string
	var values []interface{}

	for k, v := range mk.params {
		fields = append(fields, k)
		values = append(values, v)
	}

	if len(mk.conditionField) > 0 {
		values = append(values, mk.conditionValue)
	}

	sentence := buildUpdate(fields, table, *mk)

	return executeSentence(mk.db, sentence, values)
}

func executeSentence(pg *pgxpool.Pool, sql string, values []interface{}) (int64, error) {
	ctx := context.Background()

	conn, err := pg.Acquire(ctx)

	if err != nil {
		return 0, err
	}

	defer conn.Release()

	rows, err := conn.Exec(ctx, sql, values...)

	if err != nil {
		return 0, err
	}

	return rows.RowsAffected(), nil

}

func executeSentenceReturning(pg *pgxpool.Pool, sql string, values []interface{}, model *MkModel) error {
	ctx := context.Background()

	conn, err := pg.Acquire(ctx)

	if err != nil {
		return err
	}

	defer conn.Conn().Close(ctx)

	err = conn.QueryRow(ctx, sql, values...).Scan(&model.model)

	if err != nil {
		return err
	}

	return nil

}

// func executeSelectRows(pg *pgxpool.Pool, sql string, values []interface{}, model interface{}) (any, error) {
// 	ctx := context.Background()

// 	conn, err := pg.Acquire(ctx)

// 	if err != nil {
// 		return nil, err
// 	}

// 	defer conn.Conn().Close(ctx)

// 	rows, err := conn.Query(ctx, sql, values...)

// 	if err != nil {
// 		return nil, err
// 	}

// 	r, err := pgx.CollectRows(rows, pgx.RowToStructByName[model])

// }

// Omite el campo a ser actualizado
func (mk *MkModel) Omit(field string) *MkModel {
	delete(mk.params, field)
	return mk
}

// Omite multiples campos a ser actualizados
func (mk *MkModel) OmitMany(fields []string) *MkModel {
	for _, v := range fields {
		delete(mk.params, v)
	}

	return mk
}

func (mk *MkModel) Returning() *MkModel {
	params := extractParamFields(*mk.model)
	var sb strings.Builder

	sb.WriteString(" returning (")
	for f, v := range params {
		if f < len(params)-1 {
			sb.WriteString(v + ",")
		} else {
			sb.WriteString(v + ") ")
		}
	}

	mk.returning = sb.String()

	return mk
}

// Construye una condicion sencilla para la actualizacion, no toma en cuenta el Omitir
func (mk *MkModel) Where(field string, operator string, value any) *MkModel {
	mk.conditionField = field
	mk.conditionOperator = operator
	mk.conditionValue = value

	return mk
}

func buildSelect(fields []string, table string, mk MkModel) string {
	var sb strings.Builder

	sb.WriteString("select ")

	for f, v := range fields {
		if f < len(fields)-1 {
			sb.WriteString(v + ",")
		} else {
			sb.WriteString(v)
		}
	}

	sb.WriteString(" from " + table)

	if len(mk.conditionField) > 0 {
		sb.WriteString(" where " + mk.conditionField + " " + mk.conditionOperator + " $1")
	}

	return sb.String()
}

func buildInsert(fields []string, table string) string {
	var sb strings.Builder
	var sbf strings.Builder

	sb.WriteString("insert into " + table + " (")
	sbf.WriteString("values (")

	for f, v := range fields {
		if f < len(fields)-1 {
			sb.WriteString(v + ",")
			sbf.WriteString("$" + strconv.Itoa(1+f) + ",")
		} else {
			sb.WriteString(v + ") ")
			sbf.WriteString("$" + strconv.Itoa(1+f) + ")")
		}
	}

	return sb.String() + sbf.String()
}

func buildUpdate(fields []string, table string, mk MkModel) string {
	var sb strings.Builder

	sb.WriteString("update " + table + " set ")

	for f, v := range fields {
		if f < len(fields)-1 {
			sb.WriteString(v + " = " + "$" + strconv.Itoa(1+f) + ", ")
		} else {
			sb.WriteString(v + " = " + "$" + strconv.Itoa(1+f))
		}
	}

	if len(mk.conditionField) > 0 {
		sb.WriteString(" where " + mk.conditionField + " " + mk.conditionOperator + " $" + strconv.Itoa(len(fields)+1))
	}

	return sb.String()
}

func extractParams(value interface{}) map[string]interface{} {
	m := make(map[string]interface{})

	val := reflect.ValueOf(value)
	modelType := val.Type().Elem()

	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)

		fieldName := field.Tag.Get("json") // Use mk tag for field name
		s := strings.Split(fieldName, ",")
		fieldName = s[0]
		if fieldName == "" {
			continue //Skip if tag not set
		}
		fieldVal := val.Elem().Field(i).Interface()

		m[fieldName] = fieldVal

	}

	return m
}

func extractParamFields(value interface{}) []string {
	var s []string

	val := reflect.ValueOf(value)
	modelType := val.Type().Elem()

	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)

		fieldName := field.Tag.Get("json") // Use mk tag for field name
		v := strings.Split(fieldName, ",")
		fieldName = v[0]
		if fieldName == "" {
			continue //Skip if tag not set
		}

		s = append(s, fieldName)

	}

	return s
}
