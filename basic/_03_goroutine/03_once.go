package _03_goroutine

import (
	"fmt"
	"sync"
	"time"
)

func demo21() {
	once := sync.Once{}
	a := 0
	for i := 0; i < 10; i++ {
		go func(i int) {
			once.Do(func() {
				a++
				fmt.Printf("go routine %d  done\n", i)
			})
		}(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println(a)
}
