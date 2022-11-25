package util

import "strings"

func EqualsIgnoreCase(s1, s2 string) bool {
	return strings.ToLower(s1) == strings.ToLower(s2)
}

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
