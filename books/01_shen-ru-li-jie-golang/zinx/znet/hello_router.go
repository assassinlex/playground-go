package znet

import (
	"fmt"
	"playground/books/01_shen-ru-li-jie-golang/zinx/ziface"
)

// HelloRouter 自定义 ping router
type HelloRouter struct {
	BaseRouter
}

func (h *HelloRouter) PreHandle(_ ziface.IRequest) {
}

func (h *HelloRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Hello Router Handle")
	if err := request.GetConnection().SendMsg(request.GetMsgID(), []byte("hello zinx router")); err != nil {
		fmt.Println("call back ping error")
	}
}
