package ziface

// IMsgHandle 消息管理抽象
type IMsgHandle interface {
	DoMsgHandler(request IRequest)       // 马上以非阻塞的方式处理消息
	AddRouter(id uint32, router IRouter) // 为消息添加具体逻辑
}
