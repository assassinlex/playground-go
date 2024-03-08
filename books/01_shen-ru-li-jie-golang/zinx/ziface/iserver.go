package ziface

// IServer 服务器接口
type IServer interface {
	Start()                   // 启动
	Stop()                    // 停止
	Serve()                   // 服务
	AddRouter(router IRouter) // 注册路由
}
