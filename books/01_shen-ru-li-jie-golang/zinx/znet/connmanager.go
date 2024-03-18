package znet

import (
	"errors"
	"fmt"
	"playground/books/01_shen-ru-li-jie-golang/zinx/ziface"
	"sync"
)

// ConnManager 连接管理实现
type ConnManager struct {
	connections map[uint32]ziface.IConnection // 连接集合
	connLock    sync.RWMutex                  // 连接读写锁
}

func (c *ConnManager) Add(conn ziface.IConnection) {
	c.connLock.Lock() // 写锁
	defer c.connLock.Unlock()
	c.connections[conn.GetConnID()] = conn
	fmt.Printf("connection add to conn manager succeed: length = %d\n", c.Len())
}

func (c *ConnManager) Remove(conn ziface.IConnection) {
	c.connLock.Lock() // 写锁
	defer c.connLock.Unlock()
	delete(c.connections, conn.GetConnID()) // 仅做移除操作, 未做断开连接 & 停止业务等操作
	fmt.Printf("connection %d removed from conn manager succeed: length = %d\n", conn.GetConnID(), c.Len())
}

func (c *ConnManager) Get(id uint32) (ziface.IConnection, error) {
	c.connLock.RLock() // 读锁
	defer c.connLock.RUnlock()
	conn, ok := c.connections[id]
	if !ok {
		return nil, errors.New("connection is not found")
	}
	return conn, nil

}

func (c *ConnManager) Len() int {
	return len(c.connections)
}

func (c *ConnManager) ClearConn() {
	c.connLock.Lock() // 写锁
	defer c.connLock.Unlock()
	for id, connection := range c.connections {
		connection.Stop() // 停止连接业务
		delete(c.connections, id)
	}
	fmt.Printf("all connections are cleared: length = %d\n", c.Len())
}

// NewConnManager 构造器
func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]ziface.IConnection),
	}
}
