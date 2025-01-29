package concat

import "strings"

func Concat(str []string) string {
	result := ""
	for _, v := range str {
		result += v
	}
	return result
}

func ConcatBuilder(str []string) string {
	var sb strings.Builder
	for _, s := range str {
		sb.WriteString(s)
	}
	return sb.String()
}

func ConcatCopy(str []string) string {
	totalLen := 0
	for _, s := range str {
		totalLen += len(s)
	}
	result := make([]byte, totalLen)
	position := 0
	for _, s := range str {
		copy(result[position:], s)
		position += len(s)
	}
	return string(result)
}
