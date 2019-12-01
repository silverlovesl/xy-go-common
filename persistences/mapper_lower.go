package persistences

import (
	"strings"

	"github.com/jinzhu/inflection"
)

// GonicLowerCaseMapper E.g. ID -> ID, USER -> User and to table names: UserID -> USER_ID, MyUID -> MY_UID
type GonicLowerCaseMapper struct {
	wordMap map[string]bool
}

// GonicLowerCasePluralizeMapper E.g. USERS -> user, COMPANIES -> Company and to table names: User -> users, CompanyAccount -> company_accounts
type GonicLowerCasePluralizeMapper struct {
	wordMap map[string]bool
}

func isASCIILower(r rune) bool {
	return 'a' <= r && r <= 'z'
}

func toASCIILower(r rune) rune {
	if 'A' <= r && r <= 'Z' {
		r -= ('A' - 'a')
	}
	return r
}

func gonicLowerCasedName(name string) string {
	newstr := make([]rune, 0, len(name)+3)
	for idx, chr := range name {
		if isASCIIUpper(chr) && idx > 0 {
			if !isASCIIUpper(newstr[len(newstr)-1]) {
				newstr = append(newstr, '_')
			}
		}

		if !isASCIIUpper(chr) && idx > 1 {
			l := len(newstr)
			if isASCIIUpper(newstr[l-1]) && isASCIIUpper(newstr[l-2]) {
				newstr = append(newstr, newstr[l-1])
				newstr[l-1] = '_'
			}
		}

		newstr = append(newstr, chr)
	}

	return strings.ToLower(string(newstr))
}

// Obj2Table provides Obj field -> Table column mapping.
func (mapper GonicLowerCaseMapper) Obj2Table(name string) string {
	return gonicLowerCasedName(name)
}

// Table2Obj Provides Table column -> Obj field mapping.
func (mapper GonicLowerCaseMapper) Table2Obj(name string) string {
	newstr := make([]rune, 0)

	name = strings.ToLower(name)
	parts := strings.Split(name, "_")

	for _, p := range parts {
		_, isInitialism := mapper.wordMap[strings.ToLower(p)]
		for i, r := range p {
			if i == 0 || isInitialism {
				r = toASCIILower(r)
			}
			newstr = append(newstr, r)
		}
	}

	return string(newstr)
}

// Obj2Table provides Obj field -> Table column mapping.
func (mapper GonicLowerCasePluralizeMapper) Obj2Table(name string) string {
	return inflection.Plural(gonicLowerCasedName(name))
}

// Table2Obj Provides Table column -> Obj field mapping.
func (mapper GonicLowerCasePluralizeMapper) Table2Obj(name string) string {
	newstr := make([]rune, 0)

	name = strings.ToLower(name)
	parts := strings.Split(name, "_")

	for _, p := range parts {
		_, isInitialism := mapper.wordMap[strings.ToLower(p)]
		for i, r := range p {
			if i == 0 || isInitialism {
				r = toASCIILower(r)
			}
			newstr = append(newstr, r)
		}
	}

	return inflection.Singular(string(newstr))
}

// LintGonicLowerCaseMapper lists of commonly used acronyms
var LintGonicLowerCaseMapper = GonicLowerCaseMapper{
	wordMap: wordMap,
}

// LintGonicLowerCasePluralizeMapper lists of commonly used acronyms
var LintGonicLowerCasePluralizeMapper = GonicLowerCasePluralizeMapper{
	wordMap: wordMap,
}
