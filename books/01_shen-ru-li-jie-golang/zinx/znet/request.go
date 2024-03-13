package znet

import "playground/books/01_shen-ru-li-jie-golang/zinx/ziface"

type Request struct {
	conn ziface.IConnection
	data ziface.IMessage
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data.GetData()
}

func (r *Request) GetMsgID() uint32 {
	return r.data.GetMsgID()
}
