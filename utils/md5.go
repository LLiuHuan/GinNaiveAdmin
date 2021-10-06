package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5V md5加密
//@author: [LLiuHuan](https://github.com/LLiuHuan)
// @function: MD5V
// @description: md5加密
// @param: str []byte
// @return: string
func MD5V(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}