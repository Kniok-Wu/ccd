/**
 * @Time: 2024/1/15 16:57
 * @Author: kniokwu@gmail.com
 * @File: api.go
 * @Software: GoLand CCDSoftware
 * @Description: $
 */

package communication

import (
	"CCDSoftware/Instruction"
	"fmt"
	"net"
	"time"
)

// UDPPort 上位机和CCD固定端口号
var UDPPort int = 32768

// RemoteIP CCD固定IP地址
var RemoteIP string = "192.168.0.2"

// Conn 管理与 CCD 的连接
type Conn struct {
	localAddr  *net.UDPAddr
	remoteAddr *net.UDPAddr
	ins        Instruction.Instruction
	udpConn    *net.UDPConn
}

// NewConn 新建一个UDP服务
func NewConn(eth string, version string) (*Conn, error) {
	conn := &Conn{
		localAddr:  nil,
		remoteAddr: nil,
		ins:        nil,
		udpConn:    nil,
	}

	// 1. 生成远程 UDP 地址
	remoteAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", RemoteIP, UDPPort))
	if err != nil {
		return nil, ErrorInvalidAddress{Message: "请检查远程IP地址是否正确。"}
	}

	// 2. 绑定网卡
	iface, err := net.InterfaceByName(eth)
	if err != nil {
		return nil, ErrorNoSuchInterface{Message: "请选择正确的网卡。"}
	}
	addrs, err := iface.Addrs()

	// 3. 获取网卡的 IPV4 地址
	var localIp *net.IPNet = nil
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && ipNet.IP.To4() != nil && !ipNet.IP.IsLoopback() { // !ipNet.IP.IsLoopback() 是为了避免回环 例如127.0.0.1
			localIp = ipNet
			break
		}
	}
	if localIp == nil {
		return nil, ErrorInterfaceSettings{Message: "请检查网卡设置是否正确。"}
	}

	// 4. 生成本地 UDP 地址
	localAddr, _ := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", localIp.IP, UDPPort))

	// 5. 建立 UDP 连接
	var udpConn *net.UDPConn

	udpConn, err = net.DialUDP("udp", localAddr, remoteAddr)

	if udpConn == nil {
		return nil, ErrorInvalidLocalPorts{Message: fmt.Sprintf("无法启动UDP服务，请尝试释放本地端口: %d。", UDPPort)}
	}

	conn.localAddr = localAddr
	conn.remoteAddr = remoteAddr
	switch version {
	case "V20170721":
		conn.ins = Instruction.InitInstructionV2017(iface.HardwareAddr)
	default:
		conn.ins = nil
	}
	conn.udpConn = udpConn

	// 6. 初始化指令
	return conn, nil
}

func (conn *Conn) send(data []byte) error {
	_, err := conn.udpConn.Write(data)
	return err
}

// Test 发送一个测试连接的 UDP 包
func (conn *Conn) Test() error {
	return conn.send(conn.ins.TestInstruction())
}

// Last 发送上一条命令
func (conn *Conn) Last() error {
	return conn.send(conn.ins.LastInstruction())
}

// Read 从连接中读取数据
func (conn *Conn) Read() error {
	var err error = nil
	data := make([]byte, 1024)
	for {
		err = conn.udpConn.SetReadDeadline(time.Now().Add(time.Second))
		if err != nil {
			return err
		}
		_, _, err = conn.udpConn.ReadFromUDP(data)
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			err = conn.Last()
			if err != nil {
				return err
			}
		} else {
			return err
		}

		fmt.Println(data)
	}
}

// Close 关闭 UDP 连接
func (conn *Conn) Close() {
	conn.udpConn.Close()
}
