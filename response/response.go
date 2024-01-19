/**
 * @Time: 2024/1/18 16:01
 * @Author: kniokwu@gmail.com
 * @File: response.go
 * @Software: GoLand CCDSoftware
 * @Description: $
 */

package response

type CCDDataV2017 struct {
	SeriesID            int   // 帧序号
	Capacitance         int   // 电容值
	SwitchID            int   // 光开关号
	DeviceID            int   // 设备号
	TemperatureAD       int   // 温度AD值
	FrameTick           int   // 帧时刻
	PixelLightIntensity []int // 像素光强度
}
