package znet

import (
	"fmt"
	"playground/books/01_shen-ru-li-jie-golang/zinx/ziface"
)

// PingRouter 自定义 ping router
type PingRouter struct {
	BaseRouter
}

func (p *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call router PreHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping...\n"))
	if err != nil {
		fmt.Println("call back ping error")
	}
}

func (p *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call router Handle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping\n"))
	if err != nil {
		fmt.Println("call back ping error")
	}
}

func (p *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call router PostHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping...\n"))
	if err != nil {
		fmt.Println("call back ping error")
	}
}
