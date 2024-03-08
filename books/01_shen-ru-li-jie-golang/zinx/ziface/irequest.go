package ziface

// IRequest 请求接口
type IRequest interface {
	GetConnection() IConnection // 获取连接
	GetDate() []byte            // 获取数据
}
