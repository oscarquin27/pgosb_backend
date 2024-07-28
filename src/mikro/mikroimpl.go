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

func (mk *MkModel) Model(value interface{}) (*MkModel) {
	params := extractParams(value)

	return &MkModel{
		db:    mk.db,
		model: &value,
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

func executeSentence(pg *pgxpool.Pool, sql string, values []interface{}) (int64, error){
	ctx := context.Background()

	conn, err := pg.Acquire(ctx)

	if err != nil {
		return 0, err
	}

	defer conn.Conn().Close(ctx)

	rows, err := conn.Exec(ctx, sql, values...)

	
	if err != nil {
		return 0, err
	}

	return rows.RowsAffected(), nil

}

//Omite el campo a ser actualizado
func (mk *MkModel) Omit(field string) (*MkModel){
	delete(mk.params, field)
	return mk
}

//Omite multiples campos a ser actualizados
func (mk *MkModel) OmitMany(fields []string) (*MkModel){
	for _, v := range fields {
		delete(mk.params, v)
	}

	return mk
}

//Construye una condicion sencilla para la actualizacion, no toma en cuenta el Omitir
func (mk *MkModel) Where(field string, operator string, value any) (*MkModel) {
	mk.conditionField = field
	mk.conditionOperator = operator
	mk.conditionValue = value	

	return mk
}



func buildInsert(fields []string, table string) string{
	var sb strings.Builder
	var sbf strings.Builder
	
	sb.WriteString("insert into " + table + " (")
	sbf.WriteString("values (")

	for f, v := range fields {
		if f < len(fields) - 1 {
			sb.WriteString(v + ",")
			sbf.WriteString("$"+strconv.Itoa(1 + f)+",")
		} else {
			sb.WriteString(v + ") ")
			sbf.WriteString("$"+strconv.Itoa(1 + f)+")")
		}
	}


	return sb.String() + sbf.String()
}

func buildUpdate(fields []string, table string, mk MkModel) string{
	var sb strings.Builder
	
	sb.WriteString("update " + table + " set ")


	for f, v := range fields {
		if f < len(fields) - 1 {
			sb.WriteString(v + " = " + "$"+strconv.Itoa(1 + f)+", ")
		} else{
			sb.WriteString(v + " = " + "$"+strconv.Itoa(1 + f))
		}
	}
	
	if len(mk.conditionField) > 0 {
		sb.WriteString(" where " + mk.conditionField + " " + mk.conditionOperator + " $" + strconv.Itoa(len(fields)+1))
	}

	return sb.String()
}


func extractParams(value interface{}) (map[string]interface{}) {
	m := make(map[string]interface{})

	val := reflect.ValueOf(value)
    modelType := val.Type().Elem()

    for i := 0; i < modelType.NumField(); i++ {
        field := modelType.Field(i)

        fieldName := field.Tag.Get("mk") // Use mk tag for field name
        if fieldName == "" {
            continue //Skip if tag not set
        }
        fieldVal := val.Elem().Field(i).Interface()
		
		m[fieldName] = fieldVal
        
	}

	return m
}