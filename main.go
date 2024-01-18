/**
 * @Time: 2024/1/17 17:41
 * @Author: kniokwu@gmail.com
 * @File: main.go
 * @Software: GoLand CCDSoftware
 * @Description: $
 */

package main

import (
	"CCDSoftware/communication"
	"fmt"
)

func main() {
	conn, err := communication.NewConn("en8", "V20170721")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	err = conn.Test()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = conn.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
}
