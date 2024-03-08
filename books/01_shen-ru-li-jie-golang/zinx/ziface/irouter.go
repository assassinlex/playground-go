package ziface

// IRouter 路由接口
type IRouter interface {
	PreHandle(request IRequest)  // 处理业务之前的钩子函数
	Handle(request IRequest)     // 处理业务的函数
	PostHandle(request IRequest) // 处理业务之后的钩子函数
}
