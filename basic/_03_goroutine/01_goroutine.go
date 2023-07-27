package _03_goroutine

import (
	"fmt"
)

func demo01() {
	go func(v interface{}) {
		fmt.Println(v)
	}("hello")
}

func demo02() {
	//
}
