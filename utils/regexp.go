package utils

import "regexp"

// IsPasswordRuleEnabled パスワードに半角小文字大文字英字と数字を含む8文字以上か
func IsPasswordRuleEnabled(password string) bool {
	// 8文字以上か
	if len(password) < 8 {
		return false
	}

	// 小文字を含むか
	containUpperCase := regexp.MustCompile(`.*[a-z]`)
	if !containUpperCase.MatchString(password) {
		return false
	}

	// 大文字を含むか
	containLowerCase := regexp.MustCompile(`.*[A-Z]`)
	if !containLowerCase.MatchString(password) {
		return false
	}

	// 数値を含むか
	containNumber := regexp.MustCompile(`.*\d`)
	if !containNumber.MatchString(password) {
		return false
	}

	// 全て含んでいたらtrue
	return true
}
