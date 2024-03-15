package znet

import (
	"fmt"
	"math/rand"
	"net"
	"playground/books/01_shen-ru-li-jie-golang/zinx/utils"
	"playground/books/01_shen-ru-li-jie-golang/zinx/ziface"
)

// Server 服务器实现
type Server struct {
	Name       string            // 名称
	IPVersion  string            // ipv4 or ipv6
	IP         string            // 地址
	Port       int               // 端口
	msgHandler ziface.IMsgHandle // 业务处理逻辑
}

// NewServer Server 构造器
func NewServer() ziface.IServer {
	utils.GlobalObject.Reload()
	return &Server{
		Name:       utils.GlobalObject.Name,
		IPVersion:  "tcp4",
		IP:         utils.GlobalObject.Host,
		Port:       utils.GlobalObject.TcpPort,
		msgHandler: NewMsgHandle(),
	}
}

func (s *Server) Start() {
	fmt.Printf("[START] Server listenning at IP: %s:%d\n", s.IP, s.Port)
	fmt.Printf("[Zinx] Server config %v\n", utils.GlobalObject)
	go func() {
		// 0. 启动工作池模式
		s.msgHandler.StartWorkerPool()
		// 1. 获取 TCP Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			panic(fmt.Sprintf("resolve tcp addr failed: %v", err))
		}
		// 2. 监听 TCP Addr
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			panic(fmt.Sprintf("listen tcp addr failed: %v", err))
		}
		fmt.Printf("start Zinx server %s:%d succees.\n", s.IP, s.Port)
		cid := s.ConnIDGenerator() // 生成 connection id
		//	3. 启动 server
		for {
			// 3.1 接受客户端链接
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Printf("Accept err: %v", err)
				continue
			}
			//	3.2 处理客户端请求, conn & handler 绑定
			dealConn := NewConnection(conn, cid, s.msgHandler)
			cid++

			// 3.3 调用当前连接处理业务
			go dealConn.Start()
		}
	}()
}

func (s *Server) Stop() {
	fmt.Printf("[STOP] Zinx server %s\n", s.Name)
	//	todo: 处理善后工作 & graceful quit
}

func (s *Server) Serve() {
	s.Start()
	//	todo: 其他逻辑
	// 阻塞 goroutine
	select {}
}

func (s *Server) AddRouter(id uint32, router ziface.IRouter) {
	s.msgHandler.AddRouter(id, router)
	fmt.Println("Add router succeed.")
}

// ConnIDGenerator connID 生成器
func (s *Server) ConnIDGenerator() uint32 {
	return uint32(rand.Intn(100000000))
}

// Callback2Client 回声服务
func Callback2Client(conn *net.TCPConn, data []byte, cnt int) error {
	fmt.Println("[Conn Handle] Callback to client...")
	if _, err := conn.Write(data[:cnt]); err != nil {
		return err
	}
	return nil
}
