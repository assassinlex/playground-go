package znet

import (
	"fmt"
	"net"
	"playground/books/01_shen-ru-li-jie-golang/zinx/utils"
	"playground/books/01_shen-ru-li-jie-golang/zinx/ziface"
)

// Server 服务器实现
type Server struct {
	Name      string         // 名称
	IPVersion string         // ipv4 or ipv6
	IP        string         // 地址
	Port      int            // 端口
	Router    ziface.IRouter // 路由
}

// NewServer Server 构造器
func NewServer() ziface.IServer {
	utils.GlobalObject.Reload()
	return &Server{
		Name:      utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP:        utils.GlobalObject.Host,
		Port:      utils.GlobalObject.TcpPort,
		Router:    nil,
	}
}

func (s *Server) Start() {
	fmt.Printf("[START] Server listenning at IP: %s:%d\n", s.IP, s.Port)
	fmt.Printf("[Zinx] Server config %v\n", utils.GlobalObject)
	go func() {
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
		// todo: 生成 connection id
		cid := s.ConnIDGenerator()
		//	3. 启动 server
		for {
			// 3.1 接受客户端链接
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Printf("Accept err: %v", err)
				continue
			}
			//	3.2 todo 设置服务器最大连接数控制
			//	3.3 处理客户端请求, conn & handler 绑定
			dealConn := NewConnection(conn, cid, s.Router)
			cid++

			// 3.4 调用当前连接处理业务
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

func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
	fmt.Println("Add router succeed.")
}

// ConnIDGenerator connID 生成器
func (s *Server) ConnIDGenerator() uint32 {
	// todo:: 逻辑待实现
	return 0
}

// Callback2Client 回声服务
func Callback2Client(conn *net.TCPConn, data []byte, cnt int) error {
	fmt.Println("[Conn Handle] Callback to client...")
	if _, err := conn.Write(data[:cnt]); err != nil {
		return err
	}
	return nil
}
