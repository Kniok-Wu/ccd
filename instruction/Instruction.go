/**
 * @Time: 2024/1/18 11:07
 * @Author: kniokwu@gmail.com
 * @File: InstructionV2017.go
 * @Software: GoLand CCDSoftware
 * @Description: $
 */

package instruction

type Instruction interface {
	// Instruction 返回该指令中包含的二进制数组
	Instruction() []byte
	// TestInstruction 发送一个测试命令 监测是否存在消息的读取
	TestInstruction() Instruction
	// AdpgaInstruction 放大倍率
	AdpgaInstruction(int) Instruction
	// SwitchInstruction 光开关号
	SwitchInstruction(int) Instruction
	// CapacitanceInstruction 积分电容值
	CapacitanceInstruction() Instruction
	// ExposurePeriodInstruction 曝光周期
	ExposurePeriodInstruction(int) Instruction
	// DividerRatioInstruction 分频比
	DividerRatioInstruction(int) Instruction
	// EmptyRateInstruction 占空周期
	EmptyRateInstruction(int) Instruction
	// DisplayInstruction 在命令行输出执行指令
	DisplayInstruction()
}
