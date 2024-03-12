package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

func TestDataPackServer(t *testing.T) {
	// 创建 socket tcp server
	listener, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		panic(err)
	}
	for {
		// 接口客户端连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("处理客户端连接失败: ", err)
			continue
		}
		// 创建 goroutine 处理连接业务
		go func(conn net.Conn) {
			pack := NewDataPack() // 创建拆包对象
			for {
				// 读取头部数据
				header := make([]byte, pack.GetHeaderLen())
				_, err = io.ReadFull(conn, header)
				if err != nil {
					fmt.Println("读取头数据失败")
					break
				}
				unpack, err := pack.Unpack(header)
				if err != nil {
					fmt.Println("数据 header 解包失败")
					return
				}
				if unpack.GetDataLen() > 0 {
					msg := unpack.(*Message)
					msg.Data = make([]byte, unpack.GetDataLen())
					_, err = io.ReadFull(conn, msg.Data)
					if err != nil {
						fmt.Println("数据 body 解包失败")
						return
					}
					fmt.Println("请求数据: ", msg)
				}
			}
		}(conn)
	}
}
