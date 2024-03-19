package core

import (
	"fmt"
	"sync"
)

// Grid 格子
type Grid struct {
	GID       int          // 格子 id
	MinX      int          // 格子左边界坐标
	MaxX      int          // 格子右边界坐标
	MinY      int          // 格子上边界坐标
	MaxY      int          // 格子下边界坐标
	playerIDs map[int]bool // 当前格子内的玩家或者物体成员 id
	pIDLock   sync.RWMutex // playerIDs map 读写锁
}

func NewGrid(gid, minX, maxX, minY, maxY int) *Grid {
	return &Grid{
		GID:       gid,
		MinX:      minX,
		MaxX:      maxX,
		MinY:      minY,
		MaxY:      maxY,
		playerIDs: make(map[int]bool),
	}
}

// Add 添加玩家
func (g *Grid) Add(playerID int) {
	g.pIDLock.Lock()
	defer g.pIDLock.Unlock()
	g.playerIDs[playerID] = true
}

// Remove 移除玩家
func (g *Grid) Remove(playerID int) {
	g.pIDLock.Lock()
	defer g.pIDLock.Unlock()
	delete(g.playerIDs, playerID)
}

// GetPlayIDs 获取当前格子内所有玩家 id 集合
func (g *Grid) GetPlayIDs() []int {
	g.pIDLock.RLock()
	defer g.pIDLock.RUnlock()
	res := make([]int, 0, len(g.playerIDs))
	for id := range g.playerIDs {
		res = append(res, id)
	}
	return res
}

// 重写格式化输出方法
func (g *Grid) String() string {
	return fmt.Sprintf(
		"Grid id:%d minX:%d maxX:%d minY:%d maxY:%d playIDs:%v",
		g.GID, g.MinX, g.MaxX, g.MinY, g.MaxY, g.playerIDs,
	)
}
