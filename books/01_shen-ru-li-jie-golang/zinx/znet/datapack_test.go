package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
	"time"
)

func TestDataPackClient(t *testing.T) {
	// 创建客户端连接
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		panic(err)
	}
	for {
		pack := NewDataPack()
		// 消息 1 封包
		msg1 := &Message{
			Id:      0,
			DataLen: 5,
			Data:    []byte{'h', 'e', 'l', 'l', 'o'},
		}
		msg1Packed, _ := pack.Pack(msg1)
		// 消息 2 封包
		msg2 := &Message{
			Id:      1,
			DataLen: 7,
			Data:    []byte{'w', 'o', 'r', 'l', 'd', '!', '!'},
		}
		msg2Packed, _ := pack.Pack(msg2)
		// 组成撵包
		msgPacked := append(msg1Packed, msg2Packed...)
		// 发送数据包
		conn.Write(msgPacked)
		time.Sleep(3 * time.Second)
	}
	// 阻塞客户端
}

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
					fmt.Println("读取头数据失败: ", err)
					break
				}
				unpack, err := pack.Unpack(header)
				if err != nil {
					fmt.Println("数据 header 解包失败: ", err)
					return
				}
				if unpack.GetDataLen() > 0 {
					msg := unpack.(*Message)
					msg.Data = make([]byte, unpack.GetDataLen())
					_, err = io.ReadFull(conn, msg.Data)
					if err != nil {
						fmt.Println("数据 body 解包失败: ", err)
						return
					}
					fmt.Printf("请求数据: id %d\tlength %d\tdata %s\n", msg.Id, msg.DataLen, msg.Data)
				}
			}
		}(conn)
	}
}
