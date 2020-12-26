// 字符串处理
package tool

import (
	"bytes"
	"strings"
	"unicode"
)

// 下划线转驼峰命名
func UnderscoreToCamelCase(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

// 驼峰转下划线命名
func CamelCaseToUnderscore(name string) string {
	buffer := new(bytes.Buffer)
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.WriteRune('_')
			}
			buffer.WriteRune(unicode.ToLower(r))
		} else {
			buffer.WriteRune(r)
		}
	}
	return buffer.String()
}
