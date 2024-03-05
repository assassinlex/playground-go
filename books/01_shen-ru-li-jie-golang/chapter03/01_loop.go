package chapter03

func Loop(nums []int, step int) { // 按不同的步长测试相邻内存的局部性
	l := len(nums)
	for i := 0; i < step; i++ {
		for j := 0; j < l; j += step {
			nums[j] = 4
		}
	}
}
