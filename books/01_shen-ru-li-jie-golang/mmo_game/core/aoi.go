package core

import "fmt"

// AOIManager AOI 管理模块 [区域]
type AOIManager struct {
	MinX  int           // 区域左边界坐标
	MaxX  int           // 区域右边界坐标
	CntX  int           // x 方向格子的数量
	MinY  int           // 区域上边界坐标
	MaxY  int           // 区域下边界坐标
	CntY  int           // y 方向格子的数量
	grids map[int]*Grid // 当前区域内的格子集合 map[格子 id]格子对象
}

func NewAOIManager(minX, maxX, cntX, minY, maxY, cntY int) *AOIManager {
	aioMgr := &AOIManager{
		MinX:  minX,
		MaxX:  maxX,
		CntX:  cntX,
		MinY:  minY,
		MaxY:  maxY,
		CntY:  cntY,
		grids: make(map[int]*Grid),
	}
	// 初始化 AOI: 计算所有格子 id, 编号等
	for y := 0; y < cntY; y++ {
		for x := 0; x < cntX; x++ {
			gid := y*cntX + x
			aioMgr.grids[gid] = NewGrid(
				gid,
				aioMgr.MinX+x*aioMgr.gridWidth(),
				aioMgr.MinX+(x+1)*aioMgr.gridWidth(),
				aioMgr.MinY+y*aioMgr.gridHeight(),
				aioMgr.MinY+(y+1)*aioMgr.gridHeight(),
			)
		}
	}
	return aioMgr
}

// 每个格子在 x 方向上的宽度
func (am *AOIManager) gridWidth() int {
	return (am.MaxX - am.MinX) / am.CntX
}

// 每个格子在 y 方向上的高度
func (am *AOIManager) gridHeight() int {
	return (am.MaxY - am.MinY) / am.CntY
}

// 重写格式化输出
func (am *AOIManager) String() string {
	gridStr := ""
	for _, grid := range am.grids {
		gridStr += fmt.Sprintf("\t%s", grid)
	}
	aioStr := fmt.Sprintf(`AIOManager:
    minxX: %d, maxX: %d, cntX: %d, minY: %d, maxY: %d, cntY:%d
Grids in AOI Manager:
%s
`, am.MinX, am.MaxX, am.CntX, am.MinY, am.MaxY, am.CntY, gridStr)
	return aioStr
}

// GetSurroundGridsByGid 根据给定格子的 id 返回该格子周围的相邻的格子 [Ps: 注意扣紧边界条件]
func (am *AOIManager) GetSurroundGridsByGid(gid int) (grids []*Grid) {
	if _, ok := am.grids[gid]; !ok { // 非法格子
		return
	}
	grids = append(grids, am.grids[gid]) // 当前格子加入结果集
	idx := gid % am.CntX                 // 计算格子 x 轴索引
	if idx > 0 {                         // 计算格子 x 轴左右两边是否有格子 & 格子加入结果集
		grids = append(grids, am.grids[gid-1])
	}
	if idx < am.CntX-1 {
		grids = append(grids, am.grids[gid+1])
	}
	gidX := make([]int, 0, len(grids)) // 获取 x 轴方向上所有格子的 id
	for _, grid := range grids {
		gidX = append(gidX, grid.GID)
	}
	for _, id := range gidX { // 挨个处理每个 id 对应格子上下是否有格子 & 格子加入结果集
		idy := id / am.CntX
		if idy > 0 {
			grids = append(grids, am.grids[id-am.CntX])
		}
		if idy < am.CntY-1 {
			grids = append(grids, am.grids[id+am.CntX])
		}
	}
	return
}

// GetGIDByPos 根据坐标获取 Grid.GID
func (am *AOIManager) GetGIDByPos(x, y float32) int {
	idx := (int(x) - am.MinX) / am.gridWidth()
	idy := (int(y) - am.MinY) / am.gridHeight()
	return idy*am.CntX + idx
}

// GetPIDsByPos 根据坐标获取九宫格内所有玩家 id 集合
func (am *AOIManager) GetPIDsByPos(x, y float32) (playerIDs []int) {
	gid := am.GetGIDByPos(x, y)            // 获取格子 id
	grids := am.GetSurroundGridsByGid(gid) // 获取格子九宫格
	for _, grid := range grids {           // 迭代每个格子获取玩家 id 集合
		playerIDs = append(playerIDs, grid.GetPlayIDs()...)
	}
	return
}

// GetPIDByGid 获取给定格子的玩家 id 集合
func (am *AOIManager) GetPIDByGid(gid int) []int {
	return am.grids[gid].GetPlayIDs()
}

// RemovePIDFromGrid 移除给定格子中的指定玩家 -- 通过格子 id
func (am *AOIManager) RemovePIDFromGrid(pid, gid int) {
	am.grids[gid].Remove(pid)
}

// AddPIDToGrid 添加玩家到指定的格子中 -- 根据格子 id
func (am *AOIManager) AddPIDToGrid(pid, gid int) {
	am.grids[gid].Add(pid)
}

// AddToGridByPos 添加玩家到指定的格子中 -- 根据横纵坐标
func (am *AOIManager) AddToGridByPos(pid int, x, y float32) {
	gid := am.GetGIDByPos(x, y) // 获取格子 id
	am.AddPIDToGrid(pid, gid)
}

// RemoveFromGridByPos 移除给定格子中的指定玩家 -- 通过横纵坐标
func (am *AOIManager) RemoveFromGridByPos(pid int, x, y float32) {
	gid := am.GetGIDByPos(x, y) // 获取格子 id
	am.RemovePIDFromGrid(pid, gid)
}
