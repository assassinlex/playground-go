package chapter01

import (
	"fmt"
	"os"
	"runtime/trace"
)

//

func Trace() {
	f, err := os.Create("trace.out") // 创建 trace 文件
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = trace.Start(f) // 启动 trace goroutine
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	fmt.Println("Hello World") // main 程序上
}
