package service

import "strconv"

func strToSafeString(original string) string {
	return original
}

func intToSafeString(original int) string {
	return strconv.Itoa(original)
}
