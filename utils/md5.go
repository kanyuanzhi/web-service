package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5Str(s string) string {
	_md5 := md5.New()
	_md5.Write([]byte(s))
	return hex.EncodeToString(_md5.Sum(nil))
}
