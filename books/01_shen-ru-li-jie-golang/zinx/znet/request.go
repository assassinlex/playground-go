package znet

import "playground/books/01_shen-ru-li-jie-golang/zinx/ziface"

type Request struct {
	conn ziface.IConnection
	data []byte
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

func (r *Request) GetDate() []byte {
	return r.data
}
