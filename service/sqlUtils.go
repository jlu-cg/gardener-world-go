package service

import (
	"strconv"
	"strings"
)

func strToSafeString(original string) string {
	result := strings.Replace(original, "'", "''", -1)
	return result
}

func intToSafeString(original int) string {
	return strconv.Itoa(original)
}
