package _11_runtime

import "runtime"

// 获取 cpu 个数
func getCpuNums() int {
	return runtime.NumCPU()
}
