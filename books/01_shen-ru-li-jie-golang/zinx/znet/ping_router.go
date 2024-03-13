package znet

import (
	"fmt"
	"playground/books/01_shen-ru-li-jie-golang/zinx/ziface"
)

// PingRouter 自定义 ping router
type PingRouter struct {
	BaseRouter
}

func (p *PingRouter) PreHandle(_ ziface.IRequest) {
	fmt.Println("Call router PreHandle")
}

func (p *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call router Handle")
	if err := request.GetConnection().SendMsg(request.GetMsgID(), []byte("ping...ping...ping")); err != nil {
		fmt.Println("call back ping error")
	}
}

func (p *PingRouter) PostHandle(_ ziface.IRequest) {
	fmt.Println("Call router PostHandle")
}
