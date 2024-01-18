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

func (ins *InstructionV2017) Instruction() []byte {
	return ins.instruction
}

func (ins *InstructionV2017) TestInstruction() Instruction {
	// 放大倍率
	ins.AdpgaInstruction(1)
	// 光开关号
	ins.SwitchInstruction(0)
	// 积分电容值
	ins.CapacitanceInstruction()
	// 采样时钟
	ins.ExposurePeriodInstruction(4000)
	// 分频比
	ins.DividerRatioInstruction(6)
	// 占空周期
	ins.EmptyRateInstruction(4000)

	return ins
}

func (ins *InstructionV2017) AdpgaInstruction(adpga int) Instruction {
	binary.BigEndian.PutUint16(ins.instruction[10:12], uint16(adpga))
	return ins
}

func (ins *InstructionV2017) SwitchInstruction(s int) Instruction {
	ins.instruction[16] = byte(s)
	return ins
}

func (ins *InstructionV2017) CapacitanceInstruction() Instruction {
	if ins.instruction[17] == byte(0x0F) {
		ins.instruction[17] = byte(0xFF)
	} else {
		ins.instruction[17] = byte(0x0F)
	}

	return ins
}

func (ins *InstructionV2017) ExposurePeriodInstruction(period int) Instruction {
	binary.BigEndian.PutUint16(ins.instruction[20:22], uint16(period))
	return ins
}

func (ins *InstructionV2017) DividerRatioInstruction(ratio int) Instruction {
	binary.BigEndian.PutUint16(ins.instruction[20:22], uint16(ratio))
	return ins
}

func (ins *InstructionV2017) EmptyRateInstruction(rate int) Instruction {
	binary.BigEndian.PutUint16(ins.instruction[22:24], uint16(rate))
	return ins
}

func (ins *InstructionV2017) DisplayInstruction() {
	fmt.Print("instruction: ")
	for i := 0; i < len(ins.instruction); i++ {
		fmt.Printf("%02x ", ins.instruction[i])
	}
	fmt.Println()
}

// UploadFrame 数据返回帧
type UploadFrame struct {
}
