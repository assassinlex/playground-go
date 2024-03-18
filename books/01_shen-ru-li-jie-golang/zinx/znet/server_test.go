package znet

import (
	"fmt"
	"io"
	"math/rand"
	"net"
	"playground/books/01_shen-ru-li-jie-golang/zinx/ziface"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	fmt.Println("client test start...")
	time.Sleep(3 * time.Second) // 3s 后发起测试
	for i := 0; i < 20; i++ {   // 模拟 20 个客户端
		go func() {
			conn, err := net.Dial("tcp", "127.0.0.1:7777")
			if err != nil {
				fmt.Println("client connect error")
				return
			}
			for {
				dp := NewDataPack()
				// 打包消息数据 & 发送
				msg := NewMsgPackage(uint32(rand.Intn(2)), []byte("test msg"))
				req, _ := dp.Pack(msg)
				_, err = conn.Write(req)
				if err != nil {
					fmt.Printf("client request failed: %v\n", err)
					return
				}
				// 解包服务器响应
				head := make([]byte, dp.GetHeaderLen())
				if _, err = io.ReadFull(conn, head); err != nil {
					fmt.Printf("read head failed: %v\n", err)
					return
				}
				resp, err := dp.Unpack(head)
				if err != nil {
					fmt.Printf("unpacked response head failed: %v\n", err)
					return
				}
				if resp.GetDataLen() > 0 {
					msg = resp.(*Message)
					msg.Data = make([]byte, msg.GetDataLen())
					if _, err = io.ReadFull(conn, msg.Data); err != nil {
						fmt.Printf("unpacked response body failed: %v\n", err)
						return
					}
					fmt.Printf("响应数据: id %d\tlength %d\tdata %s\n", msg.Id, msg.DataLen, msg.Data)
				} else {
					fmt.Println("无请求数据")
				}
				time.Sleep(1 * time.Second)
			}
		}()
	}
	select {}
}

func TestClient1(t *testing.T) {
	fmt.Println("client test start...")
	time.Sleep(3 * time.Second) // 3s 后发起测试
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client connect error")
		return
	}
	for {
		dp := NewDataPack()
		// 打包消息数据 & 发送
		msg := NewMsgPackage(uint32(1), []byte("test msg"))
		req, _ := dp.Pack(msg)
		_, err = conn.Write(req)
		if err != nil {
			fmt.Printf("client request failed: %v\n", err)
			return
		}
		// 解包服务器响应
		head := make([]byte, dp.GetHeaderLen())
		if _, err = io.ReadFull(conn, head); err != nil {
			fmt.Printf("read head failed: %v\n", err)
			return
		}
		resp, err := dp.Unpack(head)
		if err != nil {
			fmt.Printf("unpacked response head failed: %v\n", err)
			return
		}
		if resp.GetDataLen() > 0 {
			msg = resp.(*Message)
			msg.Data = make([]byte, msg.GetDataLen())
			if _, err = io.ReadFull(conn, msg.Data); err != nil {
				fmt.Printf("unpacked response body failed: %v\n", err)
				return
			}
			fmt.Printf("响应数据: id %d\tlength %d\tdata %s\n", msg.Id, msg.DataLen, msg.Data)
		} else {
			fmt.Println("无请求数据")
		}
		time.Sleep(1 * time.Second)
	}
}

func TestServer(t *testing.T) {
	s := NewServer()
	s.SetOnConnStart(DoConnectionBegin)
	s.SetOnConnStop(DoConnectionLost)
	// s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloRouter{})
	s.Serve()
}

func DoConnectionBegin(conn ziface.IConnection) {
	fmt.Println("DoConnectionBegin is called ...")
	err := conn.SendMsg(2, []byte("DoConnectionBegin BEGIN ..."))
	if err != nil {
		fmt.Println(err)
	}
}

func DoConnectionLost(_ ziface.IConnection) {
	fmt.Println("DoConnectionLost is called ...")
}
