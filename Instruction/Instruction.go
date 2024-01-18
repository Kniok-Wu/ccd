/**
 * @Time: 2024/1/18 11:07
 * @Author: kniokwu@gmail.com
 * @File: InstructionV2017.go
 * @Software: GoLand CCDSoftware
 * @Description: $
 */

package Instruction

type Instruction interface {
	TestInstruction() []byte
	LastInstruction() []byte
}
