package znet

type Message struct {
	Id      uint32 // 消息 ID
	DataLen uint32 // 消息长度
	Data    []byte // 消息内容
}

func (m *Message) GetDataLen() uint32 {
	return m.DataLen
}

func (m *Message) GetMsgID() uint32 {
	return m.Id
}

func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) SetDataLen(u uint32) {
	m.DataLen = u
}

func (m *Message) SetMsgID(u uint32) {
	m.Id = u
}

func (m *Message) SetData(bytes []byte) {
	m.Data = bytes
}
