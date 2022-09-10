package helper

import (
	"strconv"
	"strings"
	"time"
)

func GetString(val string) string {
	if !IsStringEmpty(val) {
		return val
	}
	return ""
}

func ParseDate(val string) time.Time {
	var defaultVal time.Time

	if !IsStringEmpty(val) {
		timeVal, err := time.Parse("DD-MM-YYYY", val)

		if err != nil {
			return timeVal
		}
		return defaultVal
	}
	return defaultVal
}

func ParseInt64(val string) int64 {
	defaultVal := 0
	if !IsStringEmpty(val) {
		intVal, err := strconv.ParseInt(val, 10, 64)

		if err == nil {
			return intVal
		}

		return int64(defaultVal)
	}
	return int64(defaultVal)
}

func ParseFloat64(val string) float64 {
	defaultVal := 0
	if !IsStringEmpty(val) {
		val, err := strconv.ParseFloat(val, 64)

		if err == nil {
			return val
		}

	}
	return float64(defaultVal)
}

func IsStringEmpty(val string) bool {
	return len(strings.TrimSpace(val)) == 0
}
