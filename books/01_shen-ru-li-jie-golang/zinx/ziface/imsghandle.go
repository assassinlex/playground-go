package ziface

// IMsgHandle 消息管理抽象
type IMsgHandle interface {
	DoMsgHandler(request IRequest)       // 马上以非阻塞的方式处理消息
	AddRouter(id uint32, router IRouter) // 为消息添加具体逻辑
	StartWorkerPool()                    // 启动 worker 工作池
	SendMsgToTaskQueue(request IRequest) // 将消息交给 TaskQueue, 由 worker 处理
}
