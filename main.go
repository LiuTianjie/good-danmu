/*
 * @Descripttion: good-danmu
 * @version: 1.0
 * @Author: Nickname4th
 * @Date: 2021-05-22 16:47:46
 * @LastEditors: Nickname4th
 * @LastEditTime: 2021-05-23 20:01:41
 */
package main

import in "good-danmu/src/init"

func main() {
	router := in.Routers()
	router.Run(":8000")
}
