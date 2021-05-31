/*
@Time : 2021/5/31 14:22
@Author : nickname4th
@File : utils
@Software: GoLand
*/
package parse

import (
	"good-danmu/src/middleware"
)

func GetUserFromToken(token string) (Username string, err error) {
	j := middleware.NewJWT()
	claims, err := j.ParseJWT(token)
	if err != nil {
		return "", err
	}
	return claims.Username, nil
}
