package utils

import (
	logger "fdms/src/infrastructure/log"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

const PGOSB_ACCESS_TOKEN_COOKIE string = "PGOSB_ACCESS_TOKEN"
const PGOSB_REFRESH_TOKEN_COOKIE string = "PGOSB_REFRESH_TOKEN"
const PGOSB_SESSION_STATE_COOKIE string = "PGOSB_SESSION_STATE"

func ConvertIntToString(s int) string {
	id := strconv.Itoa(s)
	return id
}

func ParseInt(s string) int {
	id, err := strconv.Atoi(s)
	if err != nil {
		id = 0
	}
	return id
}

func ParseBool(s string) bool {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return false
	}
	return b
}

func ParseBoolString(s string) string {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return "false"
	}
	return strconv.FormatBool(b)
}

func GetStringFromPointer(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func ParseInt64(s string) int64 {

	s = strings.ReplaceAll(s, ".", "")

	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		logger.Warn().Msgf("fallo la conversacion del string: %s a entero 64 ", s)
		return 0
	}
	return n
}

func ParseInt64Pointer(s *string) int64 {
	if s == nil {
		return 0
	}

	sNew := strings.ReplaceAll(*s, ".", "")

	n, err := strconv.ParseInt(sNew, 10, 64)
	if err != nil {
		logger.Warn().Msgf("fallo la conversacion del string: %s a entero 64 ", sNew)
		return 0
	}
	return n
}

func ParseInt64String(n int64) string {
	s := strconv.FormatInt(n, 10)
	return s
}

func ParseInt64StringPointer(n *int64) string {
	if n == nil {
		return ""
	}

	s := strconv.FormatInt(*n, 10)
	return s
}

func ReadJwt(tokenString string) (*jwt.Token, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})

	if err != nil {
		fmt.Println("Error parsing token:", err)
		return nil, err

	}

	return token, nil

}

func ParseFloat(s string) float64 {
	value, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0.00
	}
	return value
}

func ConvertToPgTypeInt4(i int) pgtype.Int4 {
	return pgtype.Int4{Int32: int32(i), Valid: true}
}

func ConvertFromInt4Array(i []pgtype.Int4) []string {
	s := []string{}

	for _, n := range i {
		s = append(s, ConvertFromInt4(n))
	}

	return s
}

func ConvertFromInt2Array(i []pgtype.Int2) []string {
	s := []string{}

	for _, n := range i {
		s = append(s, ConvertFromInt2(n))
	}

	return s
}

func ConvertFromTextArray(i []pgtype.Text) []string {
	s := []string{}

	for _, n := range i {
		s = append(s, n.String)
	}

	return s
}

func ConvertToTextArray(i []string) []pgtype.Text {
	s := []pgtype.Text{}

	for _, n := range i {

		s = append(s, pgtype.Text{String: n, Valid: true})
	}

	return s
}

func ConvertToInt4Array(i []string) []pgtype.Int4 {
	s := []pgtype.Int4{}

	for _, n := range i {
		s = append(s, ConvertToPgTypeInt4(ParseInt(n)))
	}

	return s
}

func ConvertToInt2Array(i []string) []pgtype.Int2 {
	s := []pgtype.Int2{}

	for _, n := range i {
		s = append(s, ConvertToPgTypeInt2(ParseInt(n)))
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
	minutes := (d.Microseconds - hours*int64(time.Hour)*1000) / int64(time.Minute)
	seconds := (d.Microseconds - hours*int64(time.Hour)*1000 - minutes*int64(time.Minute)) / int64(time.Second)

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
