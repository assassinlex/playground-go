package core

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
			//
		}
	}
	return aioMgr
}
