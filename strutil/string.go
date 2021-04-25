package strutil

import (
	"strings"
	"unicode/utf8"
)

func Join(array ...string) string {
	buf := new(strings.Builder)
	for _, s := range array {
		buf.WriteString(s)
	}
	return buf.String()
}

func Reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}
