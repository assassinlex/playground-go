package znet

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func ClientTest() {
	fmt.Println("client test start...")
	time.Sleep(3 * time.Second) // 3s 后发起测试
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("client connect error")
		return
	}
	for {
		_, err = conn.Write([]byte("hello zinx"))
		if err != nil {
			fmt.Printf("client request failed: %v\n", err)
			return
		}
		buf := make([]byte, 512)
		_, err = conn.Read(buf)
		if err != nil {
			fmt.Printf("client read response failed: %v\n", err)
			return
		}
		fmt.Printf("server response %s\n", string(buf))
		time.Sleep(1 * time.Second)
	}
}

func TestServer(t *testing.T) {
	s := NewServer("Zinx-server01")
	s.AddRouter(&PingRouter{})
	go ClientTest()
	s.Serve()
}
