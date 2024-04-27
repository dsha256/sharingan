package util

import "strconv"

func ParseFloat64(s string) float64 {
	if f64, err := strconv.ParseFloat(s, 64); err == nil {
		return f64
	}

	return 0
}

func FormatFloat64(u float64) string {
	return strconv.FormatFloat(u, 'f', -1, 64)
}
