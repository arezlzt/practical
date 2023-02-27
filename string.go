package practical

import (
	"bytes"
	"strings"
	"unicode"
)

// buffer 内嵌bytes.Buffer，支持连写
type buffer struct {
	*bytes.Buffer
}

// CamelToCase 驼峰式写法转为下划线写法
func CamelToCase(name string) string {
	buff := &buffer{Buffer: new(bytes.Buffer)}
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				_, _ = buff.WriteString("_")
			}
			_, _ = buff.WriteRune(unicode.ToLower(r))
		} else {
			_, _ = buff.WriteRune(r)
		}
	}
	return buff.String()
}

// CaseToCamel 下划线写法转为驼峰写法
func CaseToCamel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}
