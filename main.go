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
	"net"
	"time"
)

func main() {
	ifaces, _ := net.Interfaces()
	for _, iface := range ifaces {
		addr, _ := iface.Addrs()
		fmt.Println(iface.Name, addr)
	}

	conn, err := communication.InitConn("en8", "V20170721", "")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	err = conn.Test()
	time.Sleep(time.Second)
	err = conn.Test()
	time.Sleep(time.Second)
	err = conn.Test()
	time.Sleep(time.Second)
	err = conn.Test()
	time.Sleep(time.Second)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("测试指令发送成功。")

	err = conn.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
}
