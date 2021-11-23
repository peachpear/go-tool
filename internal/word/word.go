package word

import (
	"strings"
	"unicode"
)

// 转大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// 转小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

// 下划线转大写驼峰
func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", "", -1)
	s = strings.Title(s)

	return strings.Replace(s, " ", "", -1)
}

// 下线线转小写驼峰
func UnderscoreToUpperLowerCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)

	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}
