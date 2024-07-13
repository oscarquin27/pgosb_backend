package utils

import (
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func ParseInt(s string) int {
	id, err := strconv.Atoi(s)
	if err != nil {
		id = 0
	}
	return id
}

func ConvertToPgTypeInt4(i int) pgtype.Int4 {
	return pgtype.Int4{Int32: int32(i), Valid: true}
}

func ConvertFromInt4Array(i []pgtype.Int4) []string {
	var s []string

	for _, n := range(i) {
		s = append(s, ConvertFromInt4(n))
	}

	return s
}

func ConvertToInt4Array(i []string) []pgtype.Int4 {
	var s []pgtype.Int4

	for _, n := range(i) {
		s = append(s, ConvertToPgTypeInt4(ParseInt(n)))
	}

	return s
}

func ConvertFromInt4(i pgtype.Int4) string {
	return strconv.Itoa(int(i.Int32))
}

func ConvertToPgTypeText(s string) pgtype.Text {
	return pgtype.Text{String: s, Valid: true}
}

func ConvertFromText(s pgtype.Text) string {
	return s.String
}

func ConvertToPgTypeInt2(i int) pgtype.Int2 {
	return pgtype.Int2{Int16: int16(i), Valid: true}
}

func ConvertFromInt2(i pgtype.Int2) string {
	return strconv.Itoa(int(i.Int16))
}

func ConvertToPgTypeDate(d string) pgtype.Date {
	t, err := time.Parse("2000-01-01", d)
	if err != nil {
		t = time.Now()
	}
	return pgtype.Date{Time: t, Valid: true}
}

func ConvertFromDate(d pgtype.Date) string {
	d.Valid = true
	return d.Time.String()
}

func ConvertFromDateTime(d pgtype.Time) string {
	hours := d.Microseconds / int64(time.Hour) / 1000
    minutes := (d.Microseconds - hours * int64(time.Hour) * 1000) / int64(time.Minute)
    seconds := (d.Microseconds - hours * int64(time.Hour) * 1000 - minutes * int64(time.Minute)) / int64(time.Second)
  
  	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)	
}

func ConvertToPgTypeDateTime(d string) pgtype.Time {
	
	t, err := time.Parse("2024-01-01 10:00:00 00.000", d)

	if err != nil {
		t = time.Now().UTC()
	}
  
  	return pgtype.Time{Microseconds: t.Unix(), Valid: true}	
}

func ConvertToPgTypeBool(b bool) pgtype.Bool {
	return pgtype.Bool{Bool: b, Valid: true}
}

func ConvertFromBool(b pgtype.Bool) bool {
	return b.Bool
}

func ConvertToPgTypeNumeric(i int) pgtype.Numeric {
	return pgtype.Numeric{Int: big.NewInt(int64(i)), Valid: true}
}

func ConvertFromNumeric(i pgtype.Numeric) string {
	return i.Int.String()
}