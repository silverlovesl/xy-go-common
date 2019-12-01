package utils

import "strings"

// GetEscapeString return escape string
func GetEscapeString(str string) string {
	str = strings.Replace(str, "%", "\\%", -1)
	str = strings.Replace(str, "_", "\\_", -1)

	return str
}

// GetFowardMatchString return foward match query string
func GetFowardMatchString(str string) string {
	escapeStr := GetEscapeString(str)
	matchString := make([]byte, 0, 10)
	matchString = append(matchString, escapeStr...)
	matchString = append(matchString, "%"...)

	return string(matchString)
}

// GetMatchString return match query string
func GetMatchString(str string) string {
	escapeStr := GetEscapeString(str)
	matchString := make([]byte, 0, 10)
	matchString = append(matchString, "%"...)
	matchString = append(matchString, escapeStr...)
	matchString = append(matchString, "%"...)

	return string(matchString)
}
