package helper

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
