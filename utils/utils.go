package utils

import (
	"math/big"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)
func ParseInt(s string) (int){
	id, err := strconv.Atoi(s)
	if err != nil {
		id = 0
	}
	return id
}

func ConvertToPgTypeInt4(i int) (pgtype.Int4) {
	return pgtype.Int4{Int32: int32(i)}
}

func ConvertFromInt4(i pgtype.Int4) (string) {
	return strconv.Itoa(int(i.Int32))
}

func ConvertToPgTypeText(s string) (pgtype.Text){
	return pgtype.Text{String: s}
}

func ConvertFromText(s pgtype.Text) (string) {
	return s.String
}

func ConvertToPgTypeInt2(i int) (pgtype.Int2) {
	return pgtype.Int2{Int16: int16(i)}
}

func ConvertFromInt2(i pgtype.Int2) (string) {
	return strconv.Itoa(int(i.Int16))
}

func ConvertToPgTypeDate(d time.Time) (pgtype.Date){
	return pgtype.Date{Time: d}
}

func ConvertFromDate(d pgtype.Date) (time.Time) {
	return d.Time
}

func ConvertToPgTypeBool(b bool) (pgtype.Bool){
	return pgtype.Bool{Bool: b}
}

func ConvertFromBool(b pgtype.Bool) (bool){
	return b.Bool
}

func ConvertToPgTypeNumeric(i int) (pgtype.Numeric){
	return pgtype.Numeric{Int: big.NewInt(int64(i))}
}

func ConvertFromNumeric(i pgtype.Numeric) (string){
	return i.Int.String()
}