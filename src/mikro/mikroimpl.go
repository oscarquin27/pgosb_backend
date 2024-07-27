package mikro

import (
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

func (mk *MkModel) Insert(table string) (string,error) {
	var fields []string
	var values []interface{}

	for k, v := range mk.params {
		fields = append(fields, k)
		values = append(values, v)
	}

	sentence := buildInsert(fields, table)

	return sentence, nil
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

func (mk *MkModel) Update(table string) (string,error) {
	var fields []string
	var values []interface{}

	for k, v := range mk.params {
		fields = append(fields, k)
		values = append(values, v)
	}

	sentence := buildUpdate(fields, table, *mk)

	return sentence, nil
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
		} else {
			sb.WriteString(v + " = " + "$"+strconv.Itoa(1 + f))
		}
	}
	
	if len(mk.conditionField) > 0 {
		sb.WriteString(" where " + mk.conditionField + " " + mk.conditionOperator + " $" + strconv.Itoa(len(fields)+1))
		mk.params[mk.conditionField] = mk.conditionValue
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
        fieldVal := val.Elem().Field(i)
		
		m[fieldName] = fieldVal
        
	}

	return m
}