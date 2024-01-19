/**
 * @Time: 2024/1/15 17:06
 * @Author: kniokwu@gmail.com
 * @File: errors.go
 * @Software: GoLand CCDSoftware
 * @Description: $
 */

package communication

// ErrorInvalidAddress 地址不正确
type ErrorInvalidAddress struct {
	Message string
}

func (err ErrorInvalidAddress) Error() string {
	return err.Message
}

// ErrorUnableToConnect 无法创建UDP连接
type ErrorUnableToConnect struct {
	Message string
}

func (err ErrorUnableToConnect) Error() string {
	return err.Message
}

// ErrorNoSuchInterface 没有该网卡
type ErrorNoSuchInterface struct {
	Message string
}

func (err ErrorNoSuchInterface) Error() string {
	return err.Message
}

// ErrorInvalidLocalPorts 端口被占用设置
type ErrorInvalidLocalPorts struct {
	Message string
}

func (err ErrorInvalidLocalPorts) Error() string {
	return err.Message
}

// ErrorInterfaceSettings 错误的网卡设置
type ErrorInterfaceSettings struct {
	Message string
}

func (err ErrorInterfaceSettings) Error() string {
	return err.Message
}

// ErrorManuallyUpdateARPCache 无法手动写入 ARP 缓存
type ErrorManuallyUpdateARPCache struct {
	Message string
}

func (err ErrorManuallyUpdateARPCache) Error() string {
	return err.Message
}
