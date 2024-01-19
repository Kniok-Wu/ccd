/**
 * @Time: 2024/1/15 17:29
 * @Author: kniokwu@gmail.com
 * @File: dataframe_test.go

 * @Software: GoLand CCDSoftware
 * @Description: $
 */

package communication

import (
	"fmt"
	"testing"
)

func TestDataframe(t *testing.T) {
	conn, err := InitConn("en8")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = conn.SendTestInstruction()
	if err != nil {
		fmt.Println(err)
		return
	}

	data := make([]byte, 1024)
	for i := 0; i <= 10; i++ {
		fmt.Println("reading ")
		_, err = conn.udpConn.Read(data)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
