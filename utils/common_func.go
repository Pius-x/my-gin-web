package utils

import "strings"

// BuilderConcat 字符串拼接
func BuilderConcat(strArr ...string) string {

	var builder strings.Builder
	for i := 0; i < len(strArr); i++ {
		builder.WriteString(strArr[i])
	}

	return builder.String()
}
