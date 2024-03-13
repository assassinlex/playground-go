package ziface

// IRequest 请求接口
type IRequest interface {
	GetConnection() IConnection // 获取连接
	GetData() []byte            // 获取数据
	GetMsgID() uint32           // 获取数据 id
}
