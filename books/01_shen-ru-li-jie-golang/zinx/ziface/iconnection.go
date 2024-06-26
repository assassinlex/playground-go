package ziface

import "net"

// IConnection 连接接口
type IConnection interface {
	Start()                                  // 启动连接
	Stop()                                   // 停止连接
	GetTCPConnection() *net.TCPConn          // 获取当前连接的原始 socket
	GetConnID() uint32                       // 获取当前连接 id
	RemoteAddr() net.Addr                    // 获取客户端地址
	SendMsg(id uint32, data []byte) error    // 发送消息(无缓冲)
	SendBufMsg(id uint32, data []byte) error // 发送消息(有缓冲)
}

// HandleFunc 处理连接业务接口
type HandleFunc func(*net.TCPConn, []byte, int) error
