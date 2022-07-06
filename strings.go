package go_utils

import "strings"

// TrimAllSpace 去除所有空格类字符
func TrimAllSpace(str *string) string {
	s := strings.Replace(*str, "\t", "", -1)
	s = strings.Replace(s, "\r\n", "", -1)
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, " ", "", -1)
	return s
}
