package utils

import (
	"crypto/sha256"
	"fmt"
)

// SHA256Str returns SHA256 has of the string.
func SHA256Str(str string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(str)))
}
