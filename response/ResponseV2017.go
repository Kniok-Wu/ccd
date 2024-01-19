/**
 * @Time: 2024/1/18 16:01
 * @Author: kniokwu@gmail.com
 * @File: ResponseV2017.go
 * @Software: GoLand CCDSoftware
 * @Description: $
 */

package response

import "encoding/binary"

func ParseV2017(data []byte) *CCDDataV2017 {
	res := &CCDDataV2017{
		SeriesID:            0,
		Capacitance:         0,
		SwitchID:            0,
		DeviceID:            0,
		TemperatureAD:       0,
		FrameTick:           0,
		PixelLightIntensity: nil,
	}

	res.SeriesID = int(data[0])
	res.Capacitance = int(data[1] >> 7)
	res.SwitchID = int(data[1] & 0x7F)
	res.DeviceID = int(binary.BigEndian.Uint16(data[2:4]))
	res.TemperatureAD = int(binary.BigEndian.Uint16(data[4:6]))
	res.FrameTick = int(binary.BigEndian.Uint64(append(make([]byte, 2), data[6:12]...)))

	res.PixelLightIntensity = make([]int, 0)
	for i := 12; i < len(data); i += 2 {
		res.PixelLightIntensity = append(res.PixelLightIntensity, int(binary.BigEndian.Uint16(data[i:i+2])))
	}

	return res
}
