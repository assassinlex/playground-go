package znet

import (
	"fmt"
	"playground/books/01_shen-ru-li-jie-golang/zinx/ziface"
	"strconv"
)

type MsgHandle struct {
	Apis map[uint32]ziface.IRouter
}

// NewMsgHandle 构造器
func NewMsgHandle() *MsgHandle {
	return &MsgHandle{Apis: make(map[uint32]ziface.IRouter)}
}

func (m *MsgHandle) DoMsgHandler(request ziface.IRequest) {
	handler, ok := m.Apis[request.GetMsgID()]
	if !ok {
		fmt.Printf("api msg-id: %d is nout found\n", request.GetMsgID())
	}
	//handler.PreHandle(request)
	handler.Handle(request)
	//handler.PostHandle(request)
}

// AddRouter 添加消息处理逻辑
func (m *MsgHandle) AddRouter(id uint32, router ziface.IRouter) {
	if _, ok := m.Apis[id]; ok {
		panic("repeated api: " + strconv.Itoa(int(id)))
	}
	m.Apis[id] = router
	fmt.Printf("api %d registed.\n", id)
}
