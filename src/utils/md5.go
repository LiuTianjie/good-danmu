/*
 * @Descripttion: md5 function
 * @version: 1.0
 * @Author: Nickname4th
 * @Date: 2021-05-28 16:35:56
 * @LastEditors: Nickname4th
 * @LastEditTime: 2021-05-28 16:36:24
 */
package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5V(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}
