package ziface

// IMessage 请求消息接口
type IMessage interface {
	GetDataLen() uint32 // 获取消息数据大小
	GetMsgID() uint32   // 获取消息 ID
	GetData() []byte    // 获取消息内容

	SetDataLen(uint32) // 设置消息数据大小
	SetMsgID(uint32)   // 设置消息 ID
	SetData([]byte)    // 设置消息内容
}
