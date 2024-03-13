package znet

import (
	"playground/books/01_shen-ru-li-jie-golang/zinx/ziface"
)

// BaseRouter Route 基类
type BaseRouter struct{}

func (b *BaseRouter) PreHandle(_ ziface.IRequest) {}

func (b *BaseRouter) Handle(_ ziface.IRequest) {}

func (b *BaseRouter) PostHandle(_ ziface.IRequest) {}
