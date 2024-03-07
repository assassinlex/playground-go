package chapter09

import (
	"bytes"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

func readMemStatus() {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	log.Printf(
		" ===> Alloc: %d(Bytes)\tHeapIdle: %d(Bytes)\t HeapReleased: %d(Bytes)\n",
		ms.Alloc,
		ms.HeapIdle,
		ms.HeapReleased)
}

func task() {
	container := make([]int, 8)
	log.Println(" ====> task begin.")
	for i := 0; i < 32*1000*1000; i++ {
		container = append(container, i) // 模拟 container 不断扩容
		if i == 16*1000*1000 {           // 任务进行到一半时, 读取内存信息
			readMemStatus()
		}
	}
	log.Println(" ====> task end.")
}

func pprof() {
	go func() { // 启动 web 服务, pprof 开始监听
		log.Println(http.ListenAndServe(":10000", nil))
	}()
	log.Println("====> [Start].")
	readMemStatus() // 读取程序初始内存情况
	task()
	readMemStatus() // 读取程序刚刚完成内存情况
	log.Println("====> [force GC]:")
	runtime.GC() // 强制 GC
	log.Println("====> [Done].")
	readMemStatus() // 读取 GC 后内存情况

	go func() {
		for {
			readMemStatus() // 每 10s 读取一次程序执行期间内存情况
			time.Sleep(10 * time.Second)
		}
	}()

	time.Sleep(3600 * time.Second) // 阻塞 main goroutine
}

func task2() {
	log.Println(" ====> task begin.")
	for i := 0; i < 1000; i++ {
		log.Println(genRandBytes())
	}
	log.Println(" ====> task end.")
}

func genRandBytes() *bytes.Buffer {
	var buf bytes.Buffer
	for i := 0; i < 20000; i++ {
		buf.Write([]byte{'0' + byte(rand.Intn(10))})
	}
	return &buf
}

func pprof2() {
	go func() {
		task2()
		time.Sleep(time.Second)
	}()
	_ = http.ListenAndServe(":10000", nil)
}
