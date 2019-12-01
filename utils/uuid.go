package utils

import (
	uuid "github.com/satori/go.uuid"
)

// GenUUID return UUID V4.
func GenUUID() string {
	return uuid.NewV4().String()
}
