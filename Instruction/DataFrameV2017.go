/**
 * @Time: 2024/1/18 11:01
 * @Author: kniokwu@gmail.com
 * @File: DataFrameV2017.go
 * @Software: GoLand CCDSoftware
 * @Description: $
 */

package Instruction

import (
	"encoding/binary"
	"fmt"
	"net"
)

type InstructionV2017 struct {
	instruction []byte
}

func InitInstructionV2017(mac net.HardwareAddr) Instruction {
	res := &InstructionV2017{
		instruction: make([]byte, 36),
	}
	// 1. 头部
	header := uint32(0x12345678)
	binary.BigEndian.PutUint32(res.instruction[0:4], header)

	// 2. MAC 地址
	copy(res.instruction[4:10], mac)

	// 3. 一个常量
	cons := uint32(0x3501FEC0)
	binary.BigEndian.PutUint32(res.instruction[12:16], cons)

	// 4. 帧尾部
	tails := []uint32{0x80008000, 0xC0A80003, 0xC0A80002}
	for i, tail := range tails {
		binary.BigEndian.PutUint32(res.instruction[24+i*4:28+i*4], tail)
	}

	return res
}

func (ins *InstructionV2017) TestInstruction() []byte {
	// 放大倍率
	binary.BigEndian.PutUint16(ins.instruction[10:12], uint16(1))
	// 光开关号
	ins.instruction[16] = byte(0)
	// 积分电容值
	ins.instruction[17] = byte(0x0F)
	// 采样时钟
	binary.BigEndian.PutUint16(ins.instruction[18:20], uint16(20))
	// 分频比
	binary.BigEndian.PutUint16(ins.instruction[20:22], uint16(12))
	// 占空周期
	binary.BigEndian.PutUint16(ins.instruction[22:24], uint16(4918))

	fmt.Print("instruction: ")
	for i := 0; i < len(ins.instruction); i++ {
		fmt.Printf("%02x ", ins.instruction[i])
	}
	fmt.Println()

	return ins.instruction
}

func (ins *InstructionV2017) LastInstruction() []byte {
	return ins.instruction
}

// UploadFrame 数据返回帧
type UploadFrame struct {
}
