package ziface

// IServer 服务器接口
type IServer interface {
	Start()                              // 启动
	Stop()                               // 停止
	Serve()                              // 服务
	AddRouter(id uint32, router IRouter) // 注册路由
	GetConnMgr() IConnManager            // 获取连接管理器
	SetOnConnStart(func(IConnection))    // hook 函数, 下同
	SetOnConnStop(func(IConnection))
	CallOnConnStart(conn IConnection)
	CallOnConnStop(conn IConnection)
}
