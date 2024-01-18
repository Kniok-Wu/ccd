/**
 * @Time: 2024/1/18 11:07
 * @Author: kniokwu@gmail.com
 * @File: InstructionV2017.go
 * @Software: GoLand CCDSoftware
 * @Description: $
 */

package Instruction

type Instruction interface {
	// TestInstruction 发送一个测试命令 监测是否存在消息的读取
	TestInstruction() []byte
	// LastInstruction 重发上一次的命令 解决出现的无法读取的问题
	LastInstruction() []byte
	// AdpgaInstruction
}
