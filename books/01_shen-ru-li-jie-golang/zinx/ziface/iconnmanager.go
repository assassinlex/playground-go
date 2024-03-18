package ziface

// IConnManager 连接管理抽象
type IConnManager interface {
	Add(conn IConnection)               // 添加
	Remove(conn IConnection)            // 移除
	Get(id uint32) (IConnection, error) // 获取
	Len() int                           // 当前连接数
	ClearConn()                         // 断开所有连接 & 清除
}
