package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"playground/books/01_shen-ru-li-jie-golang/zinx/ziface"
)

// Connection 连接实现
type Connection struct {
	Conn         *net.TCPConn      // 当前连接 socket
	ConnID       uint32            // session id
	Closed       bool              // 连接状态
	MsgHandler   ziface.IMsgHandle // 业务处理逻辑
	ExitBuffChan chan bool         // 退出通知 channel
	msgChan      chan []byte       // 读写 goroutine 之间的消息管道
}

// NewConnection 连接构造器
func NewConnection(conn *net.TCPConn, connID uint32, handler ziface.IMsgHandle) *Connection {
	return &Connection{
		Conn:         conn,
		ConnID:       connID,
		Closed:       false,
		MsgHandler:   handler,
		ExitBuffChan: make(chan bool, 1),
		msgChan:      make(chan []byte),
	}
}

// Start 连接启动 & 工作
func (c *Connection) Start() {
	go c.StartReader() // 读取客户端请求数据 & 执行业务逻辑
	go c.StartWriter() // 响应客户端
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

func (c *Connection) SendMsg(id uint32, data []byte) error {
	// 连接可用性检测
	if c.Closed == true {
		return errors.New("当前连接已关闭, 发送数据失败")
	}
	// 数据封包
	dp := NewDataPack()
	msg, err := dp.Pack(NewMsgPackage(id, data))
	if err != nil {
		fmt.Printf("msg packed error: id = %d\n", id)
		return errors.New("msg packed error")
	}
	// 数据发送
	c.msgChan <- msg
	return nil
}

// StartReader 处理 conn 数据读取的 goroutine
func (c *Connection) StartReader() {
	fmt.Println("Reader goroutine is running")
	defer fmt.Printf("%s conn reader exit.", c.RemoteAddr().String())
	defer c.Stop()
	for {
		// 创建封包解包对象
		dp := NewDataPack()
		// 读取客户端消息头
		headBuf := make([]byte, dp.GetHeaderLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headBuf); err != nil {
			fmt.Printf("read msg head error: %v", err)
			c.ExitBuffChan <- true
			continue
		}
		// 解包 head 数据
		msg, err := dp.Unpack(headBuf)
		if err != nil {
			fmt.Printf("unpacked error: %v", err)
			c.ExitBuffChan <- true
			continue
		}
		// 解包 body 数据
		var bodyBuf []byte
		if msg.GetDataLen() > 0 {
			bodyBuf = make([]byte, msg.GetDataLen())
			if _, err = io.ReadFull(c.GetTCPConnection(), bodyBuf); err != nil {
				fmt.Printf("read msg body error: %v", err)
				c.ExitBuffChan <- true
				continue
			}
		}
		msg.SetData(bodyBuf)
		req := Request{
			conn: c,
			data: msg,
		}
		//go func(request ziface.IRequest) { // 从 Routers 中找到注册绑定 Conn 的对应 HandleFunc
		//	c.Router.PreHandle(request)
		//	c.Router.Handle(request)
		//	c.Router.PostHandle(request)
		//}(&req)
		go c.MsgHandler.DoMsgHandler(&req) // 从绑定好的消息 & 对应方法中执行对应的 Handle 方法
	}
}

// StartWriter 写数据 goroutine
func (c *Connection) StartWriter() {
	fmt.Println("[Writer goroutine si running]")
	defer fmt.Printf("%s conn writer exit!\n", c.RemoteAddr().String())
	for {
		select {
		case data := <-c.msgChan: // 写数据
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Printf("send data error: %s,m conn writer exit!\n", err)
				return
			}
		case <-c.ExitBuffChan: // goroutine 退出
			return
		}
	}
}
