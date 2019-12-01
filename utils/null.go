package utils

import (
	"github.com/guregu/null"
)

// NullStringIfEmpty returns null.String. if string is empty, returns null.String having null value.
func NullStringIfEmpty(s string) null.String {
	return null.NewString(s, s != "")
}

// NullIntIfZero returns null.String. if int is empty, returns null.Int having null value.
func NullIntIfZero(i int) null.Int {
	return null.NewInt(int64(i), i != 0)
}
