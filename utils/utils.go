package utils

import (
	"fmt"
	"strconv"
)

func ParseInt(s string) int {
	id, err := strconv.Atoi(s)
	if err != nil {
		id = 0
	}
	return id
}

func ParseFloat(s string) float32 {
	id, err := strconv.ParseFloat(s, 32)
	
	if err != nil {
		return 0
	}

	return float32(id)
}

func ConvertFromInt(i int) string {
	return strconv.Itoa(i)
}

func ConvertFromDecimal(f float32) string {
	return fmt.Sprintf("%f", f)
}