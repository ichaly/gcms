package util

import "strconv"

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func FormatLong(v int64) string {
	return strconv.FormatInt(v, 10)
}

func ParseLong(s string) int64 {
	num, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		num = 0
	}
	return num
}
