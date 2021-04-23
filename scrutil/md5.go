package scrutil

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(s string) string {
	md5 := md5.New()
	_, err := md5.Write([]byte(s))
	if err != nil {
		return ""
	}
	return hex.EncodeToString(md5.Sum(nil))
}
