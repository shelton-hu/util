package idutil

import (
	"strconv"
	"strings"

	uuid "github.com/satori/go.uuid"
)

var (
	_LetterChars = []string{
		"a", "b", "c", "d", "e", "f",
		"g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s",
		"t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F",
		"G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S",
		"T", "U", "V", "W", "X", "Y", "Z",
	}
	_Chars = []string{
		"a", "b", "c", "d", "e", "f",
		"g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s",
		"t", "u", "v", "w", "x", "y", "z", "0", "1", "2", "3", "4", "5",
		"6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H", "I",
		"J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V",
		"W", "X", "Y", "Z",
	}
)

func Gen8Uuid() string {
	uuidStr := strings.ReplaceAll(uuid.NewV4().String(), "-", "")
	res := make([]string, 8)
	for i := 0; i < 8; i++ {
		tmp := uuidStr[4*i : 4*(i+1)]
		tmpHex, _ := strconv.ParseInt(tmp, 16, 32)
		res[i] = _Chars[tmpHex%62]
	}
	return strings.Join(res, "")
}

func Gen8LetterUuid() string {
	uuidStr := strings.ReplaceAll(uuid.NewV4().String(), "-", "")
	res := make([]string, 8)
	for i := 0; i < 8; i++ {
		tmp := uuidStr[4*i : 4*(i+1)]
		tmpHex, _ := strconv.ParseInt(tmp, 16, 32)
		res[i] = _LetterChars[tmpHex%48]
	}
	return strings.Join(res, "")
}

func GenUuid() string {
	return uuid.NewV4().String()
}
