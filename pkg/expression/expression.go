package expression

import (
	"github.com/eightpigs/i18n"
	"regexp"
	"strings"
)

// Expression parsing definition.
var parseHandles = map[string]func(string) string{
	"i18n": func(s string) string {
		return i18n.Get(s).(string)
	},
}

// Parse provides parsing of string expressions.
// If the expression parsing fails, it returns as it is.
func Parse(s string) string {
	reg := regexp.MustCompile(`(\$.*\$)+`)
	arr := reg.FindAllString(s, -1)
	if len(arr) > 0 {
		value := strings.Replace(reg.ReplaceAllString(s, ""), "_", "", -1)
		name := strings.Replace(arr[0], "$", "", -1)
		if f, ok := parseHandles[name]; ok {
			return f(value)
		}
	}
	return s
}
