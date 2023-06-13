package _01_slice

import "fmt"

/**
 * 测试目的: 不管变量名是否相同，slice 的地址都是其底层数组的内存地址
 */
func demo01() {
	s1 := []int{
		1, 2, 4, 8,
	}

	fmt.Printf("%p\n", &s1) // 0x1400000c048
	fmt.Println(s1)         // [1 2 4 8]

	s1 = append(s1, s1...)

	fmt.Printf("%p\n", &s1) //0x1400000c048
	fmt.Println(s1)         // [1 2 4 8 1 2 4 8]

	s2 := append(s1, s1...)
	fmt.Printf("%p\n", &s1) //0x1400000c048
	fmt.Println(s2)         // [1 2 4 8 1 2 4 8 1 2 4 8 1 2 4 8]

}
