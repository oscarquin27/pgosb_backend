package utils

import (
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
	return d.Time.String()
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
