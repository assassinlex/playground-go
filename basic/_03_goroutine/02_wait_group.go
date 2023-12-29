package _03_goroutine

import (
	"fmt"
	"sync"
)

func demo11() {
	wg := sync.WaitGroup{}
	taskNums := 10
	wg.Add(taskNums)
	for i := 0; i < taskNums; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println("当前 i 值: ", i)
		}(i)
	}
	wg.Wait()
}
