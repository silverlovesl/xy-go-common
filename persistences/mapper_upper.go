package persistences

// GonicUpperCaseMapper implements IMapper. It will consider initialisms when mapping names.
import (
	"strings"

	"github.com/jinzhu/inflection"
)

var wordMap = map[string]bool{
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SSH":   true,
	"TLS":   true,
	"TTL":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
	"XSRF":  true,
	"XSS":   true,
}

// GonicUpperCaseMapper E.g. ID -> ID, USER -> User and to table names: UserID -> USER_ID, MyUID -> MY_UID
type GonicUpperCaseMapper struct {
	wordMap map[string]bool
}

// GonicUpperCasePluralizeMapper E.g. USERS -> User, COMPANIES -> Company and to table names: User -> USERS, CompanyAccount -> COMPANY_ACCOUNTS
type GonicUpperCasePluralizeMapper struct {
	wordMap map[string]bool
}

func isASCIIUpper(r rune) bool {
	return 'A' <= r && r <= 'Z'
}

func toASCIIUpper(r rune) rune {
	if 'a' <= r && r <= 'z' {
		r -= ('a' - 'A')
	}
	return r
}

func gonicUpperCasedName(name string) string {
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

	return strings.ToUpper(string(newstr))
}

// Obj2Table provides Obj field -> Table column mapping.
func (mapper GonicUpperCaseMapper) Obj2Table(name string) string {
	return gonicUpperCasedName(name)
}

// Table2Obj Provides Table column -> Obj field mapping.
func (mapper GonicUpperCaseMapper) Table2Obj(name string) string {
	newstr := make([]rune, 0)

	name = strings.ToLower(name)
	parts := strings.Split(name, "_")

	for _, p := range parts {
		_, isInitialism := mapper.wordMap[strings.ToUpper(p)]
		for i, r := range p {
			if i == 0 || isInitialism {
				r = toASCIIUpper(r)
			}
			newstr = append(newstr, r)
		}
	}

	return string(newstr)
}

// Obj2Table provides Obj field -> Table column mapping.
func (mapper GonicUpperCasePluralizeMapper) Obj2Table(name string) string {
	return inflection.Plural(gonicUpperCasedName(name))
}

// Table2Obj Provides Table column -> Obj field mapping.
func (mapper GonicUpperCasePluralizeMapper) Table2Obj(name string) string {
	newstr := make([]rune, 0)

	name = strings.ToLower(name)
	parts := strings.Split(name, "_")

	for _, p := range parts {
		_, isInitialism := mapper.wordMap[strings.ToUpper(p)]
		for i, r := range p {
			if i == 0 || isInitialism {
				r = toASCIIUpper(r)
			}
			newstr = append(newstr, r)
		}
	}

	return inflection.Singular(string(newstr))
}

// LintGonicUpperCaseMapper lists of commonly used acronyms
var LintGonicUpperCaseMapper = GonicUpperCaseMapper{
	wordMap: wordMap,
}

// LintGonicUpperCasePluralizeMapper lists of commonly used acronyms
var LintGonicUpperCasePluralizeMapper = GonicUpperCasePluralizeMapper{
	wordMap: wordMap,
}
