package ziface

// IDataPack 封包解包接口
type IDataPack interface {
	GetHeaderLen() uint32              // 获取包头长度
	Pack(msg IMessage) ([]byte, error) // 封包
	Unpack([]byte) (IMessage, error)   // 解包
}
