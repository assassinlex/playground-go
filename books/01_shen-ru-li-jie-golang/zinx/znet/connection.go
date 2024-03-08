package znet

import (
	"fmt"
	"net"
	"playground/books/01_shen-ru-li-jie-golang/zinx/ziface"
)

// Connection 连接实现
type Connection struct {
	Conn         *net.TCPConn   // 当前连接 socket
	ConnID       uint32         // session id
	Closed       bool           // 连接状态
	Router       ziface.IRouter // 路由
	ExitBuffChan chan bool      // 退出通知 channel
}

// NewConnection 连接构造器
func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	return &Connection{
		Conn:         conn,
		ConnID:       connID,
		Closed:       false,
		Router:       router,
		ExitBuffChan: make(chan bool, 1),
	}
}

// Start 连接启动 & 工作
func (c *Connection) Start() {
	go c.StartReader() // 连接读取请求数据 & 执行业务逻辑
	for {
		select {
		case <-c.ExitBuffChan: // 获取退出信号后直接返回
			return
		}
	}
}

func (c *Connection) Stop() {
	if c.Closed {
		return
	}
	c.Closed = true

	// todo: 关闭回调 显示调用

	_ = c.Conn.Close()     // 关闭 socket 连接
	c.ExitBuffChan <- true // 通知缓冲队列读取数据的业务, 该链接已经关闭
	close(c.ExitBuffChan)  // 关闭该连接的全部 channel
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// StartReader 处理 conn 数据读取的 goroutine
func (c *Connection) StartReader() {
	fmt.Println("Reader goroutine is running")
	defer fmt.Printf("%s conn reader exit.", c.RemoteAddr().String())
	defer c.Stop()
	for {
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Printf("recv buf error: %v", err)
			c.ExitBuffChan <- true
			continue
		}
		req := Request{
			conn: c,
			data: buf,
		}
		go func(request ziface.IRequest) { // 从 Routers 中找到注册绑定 Conn 的对应 HandleFunc
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
	}
}
