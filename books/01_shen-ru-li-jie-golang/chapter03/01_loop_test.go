package chapter03

import "testing"

var data = CreateSource(10000)

// 创建序列
func CreateSource(length int) []int {
	nums := make([]int, 0, length)
	for i := 0; i < length; i++ {
		nums = append(nums, i)
	}
	return nums
}

func BenchmarkLoopByStep1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Loop(data, 1)
	}
}

func BenchmarkLoopByStep2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Loop(data, 2)
	}
}

func BenchmarkLoopByStep3(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Loop(data, 3)
	}
}

func BenchmarkLoopByStep4(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Loop(data, 4)
	}
}

func BenchmarkLoopByStep5(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Loop(data, 5)
	}
}

func BenchmarkLoopByStep6(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Loop(data, 6)
	}
}

func BenchmarkLoopByStep12(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Loop(data, 12)
	}
}

func BenchmarkLoopByStep16(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Loop(data, 16)
	}
}
