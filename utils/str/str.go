package str

import "strings"

func GetRepeated(str string, join string, count int) string {
	slice := []string{}
	for i := 0; i < count; i++ {
		slice = append(slice, str)
	}
	return strings.Join(slice, join)
}
