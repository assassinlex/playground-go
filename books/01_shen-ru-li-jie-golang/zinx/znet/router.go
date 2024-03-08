package znet

import (
	"playground/books/01_shen-ru-li-jie-golang/zinx/ziface"
)

// BaseRouter Route 基类
type BaseRouter struct{}

func (b *BaseRouter) PreHandle(request ziface.IRequest) {}

func (b *BaseRouter) Handle(request ziface.IRequest) {}

func (b *BaseRouter) PostHandle(request ziface.IRequest) {}
