package core

import (
	"fmt"
	"testing"
)

func TestNewAOIManager(t *testing.T) {
	fmt.Println(NewAOIManager(100, 300, 4, 200, 450, 5))
}

func TestAOIManager_GetSurroundGridsByGid(t *testing.T) {
	aoiMgr := NewAOIManager(0, 250, 5, 0, 250, 5)
	for gid, _ := range aoiMgr.grids {
		grids := aoiMgr.GetSurroundGridsByGid(gid)
		fmt.Printf("gid: %d, grids len = %d\n", gid, len(grids))
		GIDs := make([]int, 0, len(grids))
		for _, grid := range grids {
			GIDs = append(GIDs, grid.GID)
		}
		fmt.Printf("\tsurrounding grid ids are %v\n", GIDs)
	}
}
