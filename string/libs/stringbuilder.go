package libs

import "strings"

func StringBuilder(str1 string, str2 string) string {
	var builder strings.Builder
	builder.WriteString(str1)
	builder.WriteString(str2)

	return builder.String()
}
